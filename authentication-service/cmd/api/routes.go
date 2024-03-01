package main

// func (app *Config) routes() http.Handler {

// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins:     []string{"https://*", "http://*"},
// 		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodConnect},
// 		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
// 		ExposeHeaders:    []string{"Link"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	}))

// 	e.POST("/authenticate", app.Authenticate)
// 	e.POST("/signup", app.SignUp)
// 	return e
// }
