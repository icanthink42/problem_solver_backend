package missionloader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type QuestionFile struct {
	Type     string                 `toml:"type"`
	Question map[string]interface{} `toml:"question"`
}

// LoadQuestionsFromFolder loads all questions from TOML files in the specified folder
func LoadQuestionsFromFolder(folderPath string) ([]Question, error) {
	var questions []Question

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if not a .toml file
		if !info.IsDir() && filepath.Ext(path) == ".toml" {
			question, err := loadQuestionFile(path)
			if err != nil {
				return fmt.Errorf("error loading %s: %v", path, err)
			}
			questions = append(questions, question)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return questions, nil
}

// getFloat64 handles both int64 and float64 values and converts them to float64
func getFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case float64:
		return val, nil
	case int64:
		return float64(val), nil
	default:
		return 0, fmt.Errorf("value is neither float64 nor int64: %v", v)
	}
}

func loadQuestionFile(filePath string) (Question, error) {
	var qFile QuestionFile

	// Read and decode the TOML file
	if _, err := toml.DecodeFile(filePath, &qFile); err != nil {
		return nil, err
	}

	// Convert the generic question data based on type
	switch qFile.Type {
	case "multiple_choice":
		var q MultipleChoiceQuestion
		q.Question = qFile.Question["question"].(string)
		q.Options = make([]string, len(qFile.Question["options"].([]interface{})))
		for i, opt := range qFile.Question["options"].([]interface{}) {
			q.Options[i] = opt.(string)
		}
		q.AnswerIndex = int(qFile.Question["answer_index"].(int64))
		if img, ok := qFile.Question["image_url"]; ok {
			imgStr := img.(string)
			q.ImageURL = &imgStr
		}
		return q, nil

	case "numerical":
		var q NumericalQuestion
		q.Question = qFile.Question["question"].(string)
		q.Type = ValueType(qFile.Question["type"].(string))

		answers := qFile.Question["answers"].([]interface{})
		q.Answers = make([]Answer, len(answers))
		for i, ans := range answers {
			ansMap := ans.(map[string]interface{})
			tolerance, err := getFloat64(ansMap["tolerance"])
			if err != nil {
				return nil, fmt.Errorf("invalid tolerance in answer %d: %v", i, err)
			}
			q.Answers[i] = Answer{
				Value:     ansMap["value"].(string),
				Tolerance: tolerance,
			}
		}
		if img, ok := qFile.Question["image_url"]; ok {
			imgStr := img.(string)
			q.ImageURL = &imgStr
		}
		return q, nil

	default:
		return nil, fmt.Errorf("unknown question type: %s", qFile.Type)
	}
}
