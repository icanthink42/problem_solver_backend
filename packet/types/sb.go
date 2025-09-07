package types

// SBPacketType identifies the type of packet being sent from client to server
type SBPacketType string

const (
	SBPacketTypeLogin        SBPacketType = "login"
	SBPacketTypeStartLobby   SBPacketType = "start_lobby"
	SBPacketTypeNextQuestion SBPacketType = "next_question"
	SBPacketTypeAnswer       SBPacketType = "answer"
	SBPacketTypeEndQuestion  SBPacketType = "end_question"
)

// SBBasePacket contains fields common to all server-bound packets
type SBBasePacket struct {
	Type SBPacketType `json:"type"`
}
