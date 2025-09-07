package questiontypes

import "strconv"

type MultipleChoiceQuestion struct {
	Question    string   `json:"question"`
	ImageURL    *string  `json:"image_url,omitempty"`
	Options     []string `json:"options"`
	AnswerIndex int      `json:"answer_index"`
}

func (q MultipleChoiceQuestion) GetQuestion() string {
	return q.Question
}

func (q MultipleChoiceQuestion) GetImageURL() *string {
	return q.ImageURL
}

func (q MultipleChoiceQuestion) CheckAnswer(answers []string) bool {
	if len(answers) != 1 {
		return false
	}
	index, err := strconv.Atoi(answers[0])
	if err != nil {
		return false
	}
	return index == q.AnswerIndex
}
