package main

// type JSONPayload struct {
// 	Title       string          `bson:"title" json:"title"`
// 	Description string          `bson:"description" json:"description"`
// 	Difficulty  string          `bson:"difficulty" json:"difficulty"`
// 	TestCases   []data.TestCase `bson:"test_cases" json:"test_cases"`
// 	TimeLimit   int             `bson:"time_limit" json:"time_limit"`
// 	MemoryLimit int             `bson:"memory_limit" json:"memory_limit"`
// 	Tags        []string        `bson:"tags" json:"tags"`
// }

// func (app *Config) WriteTest(e echo.Context) error {
// 	// read json into var

// 	var requestPayload JSONPayload
// 	_ = app.readJSON(e, &requestPayload)

// 	// insert data

// 	event := data.Problem{
// 		Title:       requestPayload.Title,
// 		Description: requestPayload.Description,
// 		Difficulty:  requestPayload.Difficulty,
// 		TestCases:   requestPayload.TestCases,
// 		TimeLimit:   requestPayload.TimeLimit,
// 		MemoryLimit: requestPayload.MemoryLimit,
// 		Tags:        requestPayload.Tags,
// 	}

// 	err := app.Models.Insert(event)
// 	if err != nil {
// 		app.ErrorJson(e, err)
// 		return err
// 	}

// 	resp := jsonResponse{
// 		Error:   false,
// 		Message: "inserted",
// 	}

// 	app.writeJSON(e, http.StatusAccepted, resp)
// 	return nil
// }
