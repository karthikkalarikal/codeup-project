package request

type AllProbles struct {
}

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type InsertProblem struct {
	Title       string     `bson:"title" json:"title"`
	Description string     `bson:"description" json:"description"`
	Difficulty  string     `bson:"difficulty" json:"difficulty"`
	TestCases   []TestCase `bson:"test_cases" json:"test_cases"`
	TimeLimit   int        `bson:"time_limit" json:"time_limit"`
	MemoryLimit int        `bson:"memory_limit" json:"memory_limit"`
	Tags        []string   `bson:"tags" json:"tags"`
}
