package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type Problem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Difficulty  string             `bson:"difficulty" json:"difficulty"`
	TestCases   []TestCase         `bson:"test_cases" json:"test_cases"`
	TimeLimit   int                `bson:"time_limit" json:"time_limit"`
	MemoryLimit int                `bson:"memory_limit" json:"memory_limit"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
