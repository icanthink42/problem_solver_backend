package sb

import "problem_solver/packet/types"

// AnswerPacket is sent by the client with their answer to a question
type AnswerPacket struct {
	types.SBBasePacket
	Answers []string `json:"answers"` // Multiple answers for numerical questions with multiple parts
}

// NewAnswerPacket creates a new answer packet
func NewAnswerPacket(answers []string) *AnswerPacket {
	return &AnswerPacket{
		SBBasePacket: types.SBBasePacket{Type: types.SBPacketTypeAnswer},
		Answers:      answers,
	}
}
