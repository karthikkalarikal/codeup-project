package api

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/render"
	"github.com/karthikkalarikal/api-gateway/pkg/api/routes"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	port   string
	engine *echo.Echo
}

//go:generate cp -r ../../templates ./templates-dir
//go:embed templates-dir/*.gohtml
var content embed.FS




func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler, userHandler handler.UserHandler, adminHandler handler.AdminHandler) *Server {
	fmt.Println("here in server")
	e := echo.New()

	templates := template.Must(template.ParseFS(content, "templates-dir/*.gohtml"))

	// Set up TemplateRenderer
	renderer := &render.TemplateRenderer{Templates: templates}
	e.Renderer = renderer

	// to do - use custom logger to make the echo's logging more appealing , zap package.
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// functions := template.FuncMap{}
	// templates := template.Must(template.New("").Funcs(functions).ParseFS(content, "templates-dir/*.gohtml"))
	// e.Renderer = &render.TemplateRenderer{Templates: templates}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodConnect},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})) // to allow front end to connect
	// e.LoadHTMLGlob()

	e.GET("/payment", userHandler.Payment)
	e.POST("/payment-intent", userHandler.GetPaymentIntent)
	e.POST("/payment-succeeded", userHandler.PaymentSuccess)

	routes.SetupUserRoutes(e.Group("/user"), authHandler, userHandler)
	routes.SetupAdminRoutes(e.Group("/admin"), adminHandler)

	return &Server{
		engine: e,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	// c.engine.Run(c.port)
	fmt.Println("port", c.port)
	if err := c.engine.Start(fmt.Sprintf(":%s", c.port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// func (app *Server) RenderTemplate(w http.ResponseWriter, r *http.Request, page string, td *TemplateData) error {
// 	templateToRender := fmt.Sprintf("templates-dir/%s.page.gohtml", page)

// 	template, existsInCache := app.engine.Renderer.(*render.TemplateRenderer).Templates.Lookup(templateToRender)

// 	if existsInCache {
// 		err := template.Execute(w, td)
// 		if err != nil {
// 			// app.errorLog.Println(err)
// 			return err
// 		}
// 	} else {
// 		err := fmt.Errorf("template %s not found", templateToRender)
// 		// app.errorLog.Println(err)
// 		return err
// 	}
// 	return nil
// }
