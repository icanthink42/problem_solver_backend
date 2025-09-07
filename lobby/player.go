package lobby

import (
	"encoding/json"
	"fmt"
	"problem_solver/packet/cb"
	"problem_solver/packet/sb"
	"problem_solver/packet/types"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name                   string
	IsHost                 bool
	conn                   *websocket.Conn
	lobby                  *Lobby
	CurrentQuestionAnswers []string
}

func NewPlayer(name string, conn *websocket.Conn, lobby *Lobby, isHost bool) *Player {
	return &Player{
		Name:   name,
		IsHost: isHost,
		conn:   conn,
		lobby:  lobby,
	}
}

func (p *Player) HandlePacket(data []byte) (*types.CBBasePacket, error) {
	var base types.SBBasePacket
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Type {
	case types.SBPacketTypeNextQuestion:
		if !p.IsHost {
			return nil, fmt.Errorf("only host can request next question")
		}

		// Get next question
		question := p.lobby.NextQuestion()
		if question == nil {
			// No more questions, game is finished
			p.lobby.State = types.StateFinished
			response := cb.NewLoginResponse(types.StateFinished, p.IsHost)
			p.lobby.BroadcastToAll(response)
			return nil, nil
		}

		// Create question packet
		questionPacket := cb.NewQuestionPacket(question)

		// Update everyone's state
		p.lobby.State = types.StateQuestion
		stateUpdate := cb.NewLoginResponse(types.StateQuestion, p.IsHost)
		p.lobby.BroadcastToAll(stateUpdate)

		// Send question to everyone
		p.lobby.BroadcastToAll(questionPacket)
		return nil, nil
	case types.SBPacketTypeAnswer:
		var packet sb.AnswerPacket
		if err := json.Unmarshal(data, &packet); err != nil {
			return nil, err
		}
		p.CurrentQuestionAnswers = packet.Answers
		response := cb.NewAnswerConfirmPacket(p.CurrentQuestionAnswers)
		p.SendPacket(response)
		return nil, nil
	case types.SBPacketTypeEndQuestion:
		if !p.IsHost {
			return nil, fmt.Errorf("only host can end question")
		}
		p.lobby.State = types.StateQuestionReview
		for _, player := range p.lobby.Players {
			question := player.lobby.Questions[player.lobby.CurrentQuestion]
			isCorrect := question.CheckAnswer(player.CurrentQuestionAnswers)
			player.SendPacket(cb.NewQuestionGradePacket(isCorrect, types.StateQuestionReview))
			player.CurrentQuestionAnswers = []string{}
		}
		return nil, nil

	default:
		return nil, fmt.Errorf("unknown packet type: %s", base.Type)
	}
}

func (p *Player) SendPacket(packet types.CBBasePacket) error {
	data, err := packet.ToJSON()
	if err != nil {
		return err
	}
	return p.conn.WriteMessage(websocket.TextMessage, data)
}
