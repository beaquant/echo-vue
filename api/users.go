package api

import (
	"net/http"

	"errors"
	"fmt"
	"github.com/beaquant/echo-vue/auth"
	"github.com/beaquant/echo-vue/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// UserSignup -
func (api *API) UserSignup(c echo.Context) error {
	fmt.Println("UserSignup")

	u := new(UserJSON)

	if err := c.Bind(u); err != nil {
		fmt.Println("c.Bind err :", err)
		return err
	}
	if err := c.Validate(u); err != nil {
		fmt.Println("c.Validate err :", err)
		return err
	}
	fmt.Println("u :", u)

	if u.Password == "" || u.Username == "" {
		return c.JSON(http.StatusBadRequest, "Missing username or password")
	}

	if api.users.HasUser(u.Username) {
		return c.JSON(http.StatusBadRequest, "username already exists")
	}

	user := api.users.AddUser(u.Username, u.Password)

	jsontoken := auth.GetJSONToken(user)

	return c.JSON(http.StatusOK, map[string]string{"id_token": jsontoken})
}

// UserLogin -
func (api *API) UserLogin(c echo.Context) error {
	u := new(UserJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	if u.Password == "" || u.Username == "" {
		return c.JSON(http.StatusBadRequest, "Missing username or password")
	}
	user := api.users.FindUser(u.Username)
	if user.Username == "" {
		return c.JSON(http.StatusBadRequest, errors.New("username not found"))
	}

	if !api.users.CheckPassword(user.Password, u.Password) {
		return c.JSON(http.StatusBadRequest, errors.New("bad password"))
	}

	jsontoken := auth.GetJSONToken(user)
	return c.JSON(http.StatusOK, map[string]string{"id_token": jsontoken})
}

func (api *API) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func (api *API) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(c echo.Context) *models.User {
	fmt.Println("GetUserFromContext")

	userclaims := auth.GetUserClaimsFromContext(c)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(c echo.Context) error {
	c.Request()
	fmt.Println("UserInfo")
	user := api.GetUserFromContext(c)
	fmt.Println("user", user)

	return c.String(http.StatusOK, "Welcome "+user.Username+"!")
}
