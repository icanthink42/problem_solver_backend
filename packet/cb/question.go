package cb

import (
	"encoding/json"
	missionloader "problem_solver/mission_loader"
	"problem_solver/packet/types"
)

// QuestionPacket is sent to all clients when a new question starts
type QuestionPacket struct {
	PacketType   types.CBPacketType     `json:"type"`
	Question     missionloader.Question `json:"question"`
	QuestionType string                 `json:"question_type"` // "multiple_choice", "numerical", or "point_selector"
}

// AnswerConfirmPacket is sent to all clients when an answer is confirmed
type AnswerConfirmPacket struct {
	PacketType types.CBPacketType `json:"type"`
	Answers    []string           `json:"answers"`
}

func (p AnswerConfirmPacket) Type() types.CBPacketType {
	return p.PacketType
}

func (p AnswerConfirmPacket) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// NewAnswerConfirmPacket creates a new answer confirm packet
func NewAnswerConfirmPacket(answers []string) types.CBBasePacket {
	return AnswerConfirmPacket{
		PacketType: types.CBPacketTypeAnswerConfirm,
		Answers:    answers,
	}
}

func (p QuestionPacket) Type() types.CBPacketType {
	return p.PacketType
}

func (p QuestionPacket) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// NewQuestionPacket creates a new question packet from a Question interface
func NewQuestionPacket(q missionloader.Question) types.CBBasePacket {
	var questionType string
	switch q.(type) {
	case missionloader.MultipleChoiceQuestion:
		questionType = "multiple_choice"
	case missionloader.NumericalQuestion:
		questionType = "numerical"
	case missionloader.PointSelectorQuestion:
		questionType = "point_selector"
	default:
		// This should never happen, but better to have a default
		questionType = "unknown"
	}

	return QuestionPacket{
		PacketType:   types.CBPacketTypeQuestion,
		Question:     q,
		QuestionType: questionType,
	}
}
