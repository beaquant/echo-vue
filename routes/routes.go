package routes

import (
	"net/http"

	"github.com/beaquant/echo-vue/api"
	"github.com/beaquant/echo-vue/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API, e *echo.Echo) {
	//func (controller *TasksController) Setup() {
	//	controller.Router.Add("POST", "/tasks", controller.createTask)
	//	controller.Router.Add("GET", "/tasks", controller.listTasks)
	//	controller.Router.Add("GET", "/tasks/:task_id", controller.getTask)
	//	controller.Router.Add("PUT", "/tasks/:task_id", controller.updateTask)
	//	controller.Router.Add("DELETE", "/tasks/:task_id", controller.deleteTask)
	//}

	//router.Add("POST", "/users", saveUser)
	//router.GET("/users/:id", getUser)
	//router.PUT("/users/:id", updateUser)
	//router.DELETE("/users/:id", deleteUser)
	//
	//mux := mux.NewRouter()
	//
	//// client static files
	e.File("/", "index.html")
	e.Static("/static", "static")
	// api
	//e.POST("/users", saveUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	// users
	e.POST("/api/users/signup", api.UserSignup)
	e.POST("/api/users/login", api.UserLogin)
	//e.POST("/api/users/info", api.UserSignup)
	//u.Handle("/info", negroni.New(
	//	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	//	negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	//))

	// quotes
	e.GET("/api/quote/random", api.Quote)
	//e.GET("/api/protected/random", api.UserLogin)
	//
	//q.Handle("/protected/random", negroni.New(
	//	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	//	negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	//))

}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
