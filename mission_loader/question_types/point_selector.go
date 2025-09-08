package questiontypes

import (
	"fmt"
	"math"
	"strconv"
)

type PointSelectorQuestion struct {
	Question      string  `json:"question"`
	ImageURL      string  `json:"image_url"`    // Required image for the vector diagram
	XComponent    string  `json:"x_comp"`       // Label for x component
	YComponent    string  `json:"y_comp"`       // Label for y component
	CorrectX      float64 `json:"correct_x"`    // Correct x coordinate
	CorrectY      float64 `json:"correct_y"`    // Correct y coordinate
	CorrectRadius float64 `json:"radius"`       // Acceptable radius from correct point
	PointerType   string  `json:"pointer_type"` // Optional: "vector" for vector arrow, otherwise dot
}

func (q PointSelectorQuestion) GetQuestion() string {
	return q.Question
}

func (q PointSelectorQuestion) GetImageURL() *string {
	return &q.ImageURL
}

// Validate ensures the question has all required fields
func (q PointSelectorQuestion) Validate() error {
	if q.ImageURL == "" {
		return fmt.Errorf("point selector questions must have an image")
	}
	return nil
}

func (q PointSelectorQuestion) CheckAnswer(answers []string) bool {
	if len(answers) != 2 {
		return false
	}

	// Parse x and y coordinates from answers
	x, err := strconv.ParseFloat(answers[0], 64)
	if err != nil {
		return false
	}

	y, err := strconv.ParseFloat(answers[1], 64)
	if err != nil {
		return false
	}

	// Calculate distance from correct point
	dx := x - q.CorrectX
	dy := y - q.CorrectY
	distance := math.Sqrt(dx*dx + dy*dy)

	// Check if point is within acceptable radius
	return distance <= q.CorrectRadius
}
