package api

import (
	"encoding/json"
	"net/http"

	"github.com/beaquant/echo-vue/auth"
	"github.com/beaquant/echo-vue/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"time"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserSignup -
func (api *API) UserSignup(c echo.Context) error {

	//decoder := json.NewDecoder(req.Body)
	//jsondata := UserJSON{}
	//err := decoder.Decode(&jsondata)
	//
	//if err != nil || jsondata.Username == "" || jsondata.Password == "" {
	//	http.Error(w, "Missing username or password", http.StatusBadRequest)
	//	return
	//}
	//
	//if api.users.HasUser(jsondata.Username) {
	//	http.Error(w, "username already exists", http.StatusBadRequest)
	//	return
	//}
	//
	//user := api.users.AddUser(jsondata.Username, jsondata.Password)
	//
	//jsontoken := auth.GetJSONToken(user)
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(jsontoken))
}

// UserLogin -
func (api *API) UserLogin(c echo.Context) error {

	//decoder := json.NewDecoder(req.Body)
	//jsondata := UserJSON{}
	//err := decoder.Decode(&jsondata)
	//
	//if err != nil || jsondata.Username == "" || jsondata.Password == "" {
	//	http.Error(w, "Missing username or password", http.StatusBadRequest)
	//	return
	//}
	//
	//user := api.users.FindUser(jsondata.Username)
	//if user.Username == "" {
	//	http.Error(w, "username not found", http.StatusBadRequest)
	//	return
	//}
	//
	//if !api.users.CheckPassword(user.Password, jsondata.Password) {
	//	http.Error(w, "bad password", http.StatusBadRequest)
	//	return
	//}
	//
	//jsontoken := auth.GetJSONToken(user)
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(jsontoken))

}
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
