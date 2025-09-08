package missionloader

import questiontypes "problem_solver/mission_loader/question_types"

// Re-export the Question interface and types
type Question = questiontypes.Question
type MultipleChoiceQuestion = questiontypes.MultipleChoiceQuestion
type NumericalQuestion = questiontypes.NumericalQuestion
type PointSelectorQuestion = questiontypes.PointSelectorQuestion
type Answer = questiontypes.Answer
type ValueType = questiontypes.ValueType

// Re-export constants
const (
	TypeInt   = questiontypes.TypeInt
	TypeFloat = questiontypes.TypeFloat
)
