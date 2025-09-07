package sb

import "problem_solver/packet/types"

// LoginPacket is sent by the client to join a lobby
type LoginPacket struct {
	types.SBBasePacket
	Name      string `json:"name"`
	LobbyCode string `json:"lobby_code"`
}

// StartLobbyPacket is sent by the client to create and host a new lobby
type StartLobbyPacket struct {
	types.SBBasePacket
	Name      string `json:"name"`
	LobbyCode string `json:"lobby_code"`
}

// NewLoginPacket creates a new login packet
func NewLoginPacket(name, lobbyCode string) *LoginPacket {
	return &LoginPacket{
		SBBasePacket: types.SBBasePacket{Type: types.SBPacketTypeLogin},
		Name:         name,
		LobbyCode:    lobbyCode,
	}
}

// NewStartLobbyPacket creates a new start lobby packet
func NewStartLobbyPacket(name, lobbyCode string) *StartLobbyPacket {
	return &StartLobbyPacket{
		SBBasePacket: types.SBBasePacket{Type: types.SBPacketTypeStartLobby},
		Name:         name,
		LobbyCode:    lobbyCode,
	}
}
