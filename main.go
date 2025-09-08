package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"problem_solver/lobby"
	missionloader "problem_solver/mission_loader"
	"problem_solver/packet/cb"
	"problem_solver/packet/sb"
	"problem_solver/packet/types"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

func handleStartLobby(data []byte, conn *websocket.Conn, lobbies map[string]*lobby.Lobby) (*lobby.Player, error) {
	var startPacket sb.StartLobbyPacket
	if err := json.Unmarshal(data, &startPacket); err != nil {
		return nil, err
	}

	// Load questions from the questions folder
	questions, err := missionloader.LoadQuestionsFromFolder("questions")
	if err != nil {
		return nil, err
	}

	// Create new lobby with questions
	l := lobby.NewLobby(startPacket.LobbyCode)
	for _, q := range questions {
		l.AddQuestion(q)
	}
	lobbies[startPacket.LobbyCode] = l

	// Create host player
	player := lobby.NewPlayer(startPacket.Name, conn, l, true)
	l.AddPlayer(player)

	// Send success response
	response := cb.NewLoginResponse(types.StateWaiting, true)
	if err := player.SendPacket(response); err != nil {
		return nil, err
	}

	return player, nil
}

func handleLogin(data []byte, conn *websocket.Conn, lobbies map[string]*lobby.Lobby) (*lobby.Player, error) {
	var loginPacket sb.LoginPacket
	if err := json.Unmarshal(data, &loginPacket); err != nil {
		return nil, err
	}

	// Check if lobby exists
	l, exists := lobbies[loginPacket.LobbyCode]
	if !exists {
		response := cb.NewLoginFailure("Lobby does not exist")
		data, err := response.ToJSON()
		if err != nil {
			return nil, err
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return nil, err
		}
		return nil, nil
	}

	// Create player and add to lobby
	player := lobby.NewPlayer(loginPacket.Name, conn, l, false)
	l.AddPlayer(player)

	// Send success response
	response := cb.NewLoginResponse(l.State, false)
	if err := player.SendPacket(response); err != nil {
		return nil, err
	}

	return player, nil
}

func handleWebSocket(c *gin.Context, lobbies map[string]*lobby.Lobby) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	// Wait for login packet
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading login message: %v", err)
		return
	}

	// Check packet type
	var base types.SBBasePacket
	if err := json.Unmarshal(message, &base); err != nil {
		log.Printf("Error parsing packet: %v", err)
		return
	}

	var player *lobby.Player
	switch base.Type {
	case types.SBPacketTypeStartLobby:
		player, err = handleStartLobby(message, conn, lobbies)
	case types.SBPacketTypeLogin:
		player, err = handleLogin(message, conn, lobbies)
	default:
		log.Printf("First packet must be login or start_lobby")
		return
	}

	if err != nil {
		log.Printf("Error handling packet: %v", err)
		return
	}

	if player == nil {
		return // Login failed
	}

	// Main packet handling loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		// Let the player handle all other packets
		response, err := player.HandlePacket(message)
		if err != nil {
			log.Printf("Error handling packet: %v", err)
			continue
		}

		// Only send response if one was returned
		if response != nil {
			responseData, err := (*response).ToJSON()
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}
			if err := conn.WriteMessage(messageType, responseData); err != nil {
				log.Printf("Error writing response: %v", err)
				break
			}
		}
	}
}

func main() {
	// Ensure questions folder path is absolute
	questionsPath, err := filepath.Abs("questions")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loading questions from: %s", questionsPath)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lobbies := make(map[string]*lobby.Lobby)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "WebSocket server is running",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c, lobbies)
	})

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
