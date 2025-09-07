package packet

import (
	"encoding/json"
	"fmt"

	"problem_solver/packet/cb"
	"problem_solver/packet/sb"
	"problem_solver/packet/types"
)

// HandlePacket processes an incoming server-bound packet and returns a client-bound packet response
func HandlePacket(data []byte) (*types.CBBasePacket, error) {
	var base types.SBBasePacket
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Type {
	case types.SBPacketTypeLogin:
		var packet sb.LoginPacket
		if err := json.Unmarshal(data, &packet); err != nil {
			return nil, err
		}
		response := cb.NewLoginResponse(types.StateWaiting, false)
		return &response, nil

	case types.SBPacketTypeStartLobby:
		var packet sb.StartLobbyPacket
		if err := json.Unmarshal(data, &packet); err != nil {
			return nil, err
		}
		response := cb.NewLoginResponse(types.StateWaiting, true)
		return &response, nil

	default:
		return nil, fmt.Errorf("unknown packet type: %s", base.Type)
	}
}
