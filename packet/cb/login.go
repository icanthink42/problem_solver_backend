package cb

import (
	"encoding/json"
	"problem_solver/packet/types"
)

// LoginResponse is sent by the server in response to a successful login
type LoginResponse struct {
	PacketType types.CBPacketType `json:"type"`
	State      types.State        `json:"state"`
	IsHost     bool               `json:"is_host"`
}

func (p LoginResponse) Type() types.CBPacketType {
	return p.PacketType
}

func (p LoginResponse) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// LoginFailure is sent when login fails (e.g., lobby doesn't exist)
type LoginFailure struct {
	PacketType types.CBPacketType `json:"type"`
	Message    string             `json:"message"`
}

func (p LoginFailure) Type() types.CBPacketType {
	return p.PacketType
}

func (p LoginFailure) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// NewLoginResponse creates a new login response packet
func NewLoginResponse(state types.State, isHost bool) types.CBBasePacket {
	return LoginResponse{
		PacketType: types.CBPacketTypeLoginResponse,
		State:      state,
		IsHost:     isHost,
	}
}

// NewLoginFailure creates a new login failure packet
func NewLoginFailure(message string) types.CBBasePacket {
	return LoginFailure{
		PacketType: types.CBPacketTypeLoginFailure,
		Message:    message,
	}
}
