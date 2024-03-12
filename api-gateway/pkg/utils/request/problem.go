package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllProbles struct {
}

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type InsertProblem struct {
	Title          string     `bson:"title" json:"title"`
	Description    string     `bson:"description" json:"description"`
	Difficulty     string     `bson:"difficulty" json:"difficulty"`
	TestCases      []TestCase `bson:"test_cases" json:"test_cases"`
	TimeLimit      int        `bson:"time_limit" json:"time_limit"`
	MemoryLimit    int        `bson:"memory_limit" json:"memory_limit"`
	FirstHalfCode  []byte     `bson:"first_half" json:"first_half"`
	SecondHalfCode []byte     `bson:"second_half" json:"second_half"`
	Tags           []string   `bson:"tags" json:"tags"`
	CreatedAt      time.Time  `bson:"created_at" json:"created_at"`
}

type GetOneProblemById struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
}

type SubmitCodeIdRequest struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code []byte             `json:"code"`
}

type FirstHalfCode struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstHalfCode []byte             `bson:"first_half" json:"first_half"`
	// SecondHalfCode []byte             `bson:"second_half" json:"second_half"`
}
type SecondHalfCode struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// FirstHalfCode  []byte             `bson:"first_half" json:"first_half"`
	SecondHalfCode []byte `bson:"second_half" json:"second_half"`
}
