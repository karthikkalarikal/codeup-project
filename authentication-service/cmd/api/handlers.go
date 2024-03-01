package main

// type RequestPayload struct {
// 	Action string      `json:"action"`
// 	Auth   AuthPayload `json:"auth,omitempty"`
// }

// type AuthPayload struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
// type RequestPayloadSignUp struct {
// 	Action     string    `json:"action"`
// 	SignUpData data.User `json:"auth,omitempty"`
// }

// func (app *Config) Authenticate(e echo.Context) error {
// 	fmt.Println("here---")
// 	// {
// 	// 	fmt.Println("logging")
// 	// 	fmt.Println("Request Headers:", e.Request().Header)
// 	// 	body, err := io.ReadAll(e.Request().Body)
// 	// 	fmt.Println("Raw Request Body: ", body, " error ", err)

// 	// }

// 	var payload RequestPayload
// 	// e.Request().Body
// 	err := e.Bind(&payload)
// 	// err := app.readJSON(e, &payload)
// 	// fmt.Println("err in auth --:", err)
// 	if err != nil {
// 		app.ErrorJson(e, err, http.StatusBadRequest)
// 		return err
// 	}
// 	fmt.Println("payload", payload.Action)

// 	// validate the user against the database

// 	user, err := app.Models.User.GetByEmail(payload.Auth.Email)
// 	fmt.Println("user:", user)
// 	if err != nil {
// 		app.ErrorJson(e, err, http.StatusBadRequest)
// 		return err
// 	}
// 	app.writeJSON(e, http.StatusAccepted, user)
// 	return nil
// }

// func (app *Config) SignUp(e echo.Context) error {
// 	fmt.Println("here---")
// 	// {
// 	// 	fmt.Println("logging")
// 	// 	fmt.Println("Request Headers:", e.Request().Header)
// 	// 	body, err := io.ReadAll(e.Request().Body)
// 	// 	fmt.Println("Raw Request Body: ", body, " error ", err)

// 	// }

// 	var payload RequestPayloadSignUp
// 	// e.Request().Body
// 	err := e.Bind(&payload)
// 	// err := app.readJSON(e, &payload)
// 	// fmt.Println("err in auth --:", err)
// 	if err != nil {
// 		app.ErrorJson(e, err, http.StatusBadRequest)
// 		return err
// 	}
// 	fmt.Println("payload", payload.Action)

// 	// validate the user against the database

// 	user, err := app.Models.User.GetByEmail(payload.SignUpData.Email.String)
// 	fmt.Println("user:", user)
// 	if err == nil {
// 		app.ErrorJson(e, err, http.StatusBadRequest)
// 		return err
// 	}
// 	data, err := app.Models.User.UserSignUp(payload.SignUpData)
// 	if err != nil {
// 		app.ErrorJson(e, err, http.StatusBadRequest)
// 		return err
// 	}

// 	app.writeJSON(e, http.StatusAccepted, data)
// 	return nil
// }
