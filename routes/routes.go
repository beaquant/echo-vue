package routes

import (
	"github.com/beaquant/echo-vue/api"
	"github.com/beaquant/echo-vue/auth"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// NewRoutes builds the routes for the api
func NewRoutes(a *api.API, e *echo.Echo) {

	// client static files
	//e.File("/", "index.html")
	//e.Static("/static", "static")

	// api

	e.Validator = &api.CustomValidator{Validator: validator.New()}

	// users
	e.POST("/api/user/signup", a.UserSignup)
	e.POST("/api/user/login", a.UserLogin)

	r := e.Group("/api/user/info")
	r.Use(middleware.JWT(auth.GetSigningKey()))
	r.GET("", a.UserInfo)

	//e.GET("/api/user/info", a.UserInfo)
	//e.GET("/api/user/random", a.SecretQuote)
	rr := e.Group("/api/quote/protected/random")
	rr.Use(middleware.JWT(auth.GetSigningKey()))
	rr.GET("", a.SecretQuote)

	// quotes
	e.GET("/api/quote/random", a.Quote)

}
