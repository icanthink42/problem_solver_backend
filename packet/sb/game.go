package sb

import "problem_solver/packet/types"

// NextQuestionPacket is sent by the host to start the next question
type NextQuestionPacket struct {
	types.SBBasePacket
}

// EndQuestionPacket is sent by the host to end the current question
type EndQuestionPacket struct {
	types.SBBasePacket
}

// NewNextQuestionPacket creates a new next question packet
func NewNextQuestionPacket() *NextQuestionPacket {
	return &NextQuestionPacket{
		SBBasePacket: types.SBBasePacket{Type: types.SBPacketTypeNextQuestion},
	}
}

// NewEndQuestionPacket creates a new end question packet
func NewEndQuestionPacket() *EndQuestionPacket {
	return &EndQuestionPacket{
		SBBasePacket: types.SBBasePacket{Type: types.SBPacketTypeEndQuestion},
	}
}
