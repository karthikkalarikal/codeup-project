package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type Problem struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	Difficulty     string             `bson:"difficulty" json:"difficulty" validate:"oneof=easy medium hard"`
	TestCases      []TestCase         `bson:"test_cases" json:"test_cases"`
	TimeLimit      int                `bson:"time_limit" json:"time_limit"`
	MemoryLimit    int                `bson:"memory_limit" json:"memory_limit"`
	FirstHalfCode  []byte             `bson:"first_half" json:"first_half"`
	SecondHalfCode []byte             `bson:"second_half" json:"second_half"`
	Tags           []string           `bson:"tags" json:"tags"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	Prime          bool               `bson:"prime" json:"prime"`
}
