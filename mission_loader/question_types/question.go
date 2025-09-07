package questiontypes

// Question is the interface that all question types must implement
type Question interface {
	GetQuestion() string
	GetImageURL() *string
	CheckAnswer(answers []string) bool
}
