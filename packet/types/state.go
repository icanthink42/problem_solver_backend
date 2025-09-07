package types

// State represents the current state of the game
type State string

const (
	StateWaiting        State = "waiting"
	StateQuestion       State = "question"
	StateQuestionReview State = "question_review"
	StateFinished       State = "finished"
)
