package types

// CBPacketType identifies the type of packet being sent from server to client
type CBPacketType string

const (
	CBPacketTypeLoginResponse CBPacketType = "login_response"
	CBPacketTypeLoginFailure  CBPacketType = "login_failure"
	CBPacketTypeQuestion      CBPacketType = "question"
	CBPacketTypeQuestionGrade CBPacketType = "question_grade"
	CBPacketTypeAnswerConfirm CBPacketType = "answer_confirm"
)

// CBBasePacket is the interface that all client-bound packets must implement
type CBBasePacket interface {
	Type() CBPacketType
	ToJSON() ([]byte, error)
}
