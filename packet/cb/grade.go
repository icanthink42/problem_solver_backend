package cb

import (
	"encoding/json"
	"problem_solver/packet/types"
)

// QuestionGradePacket is sent to clients with their question results
type QuestionGradePacket struct {
	PacketType types.CBPacketType `json:"type"`
	IsCorrect  bool               `json:"is_correct"`
	State      types.State        `json:"state"`
}

func (p QuestionGradePacket) Type() types.CBPacketType {
	return p.PacketType
}

func (p QuestionGradePacket) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// NewQuestionGradePacket creates a new question grade packet
func NewQuestionGradePacket(isCorrect bool, state types.State) types.CBBasePacket {
	return QuestionGradePacket{
		PacketType: types.CBPacketTypeQuestionGrade,
		IsCorrect:  isCorrect,
		State:      state,
	}
}
