package questiontypes

import "strconv"

type ValueType string

const (
	TypeInt   ValueType = "int"
	TypeFloat ValueType = "float"
)

type Answer struct {
	Value     string  `json:"value"`
	Tolerance float64 `json:"tolerance"`
}

type NumericalQuestion struct {
	Question     string    `json:"question"`
	ImageURL     *string   `json:"image_url,omitempty"`
	Answers      []Answer  `json:"answers"`
	Type         ValueType `json:"type"`
	RequireOrder bool      `json:"require_order,omitempty" default:"false"`
}

func (q NumericalQuestion) GetQuestion() string {
	return q.Question
}

func (q NumericalQuestion) GetImageURL() *string {
	return q.ImageURL
}

func (q NumericalQuestion) CheckAnswer(answers []string) bool {
	if len(answers) != len(q.Answers) {
		return false
	}

	used := make([]bool, len(q.Answers))
	for i, answer := range answers {
		if q.RequireOrder {
			// In order mode, only check against the current expected answer
			if !used[i] && q.checkSingleAnswer(answer, q.Answers[i]) {
				used[i] = true
				continue
			}
			return false
		} else {
			// Try to match with any unused answer
			found := false
			for j, expected := range q.Answers {
				if !used[j] && q.checkSingleAnswer(answer, expected) {
					used[j] = true
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func (q NumericalQuestion) checkSingleAnswer(given string, expected Answer) bool {
	switch q.Type {
	case TypeInt:
		givenInt, err := strconv.ParseInt(given, 10, 64)
		if err != nil {
			return false
		}
		expectedInt, err := strconv.ParseInt(expected.Value, 10, 64)
		if err != nil {
			return false
		}
		return givenInt == expectedInt

	case TypeFloat:
		givenFloat, err := strconv.ParseFloat(given, 64)
		if err != nil {
			return false
		}
		expectedFloat, err := strconv.ParseFloat(expected.Value, 64)
		if err != nil {
			return false
		}
		diff := givenFloat - expectedFloat
		if diff < 0 {
			diff = -diff
		}
		return diff <= expected.Tolerance

	default:
		return false
	}
}
