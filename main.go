package main

import (
	"os"

	"github.com/beaquant/echo-vue/api"
	"github.com/beaquant/echo-vue/config"
	"github.com/beaquant/echo-vue/models"
	"github.com/beaquant/echo-vue/routes"
	"github.com/labstack/echo"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return port
}

func main() {
	e := echo.New()
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes.NewRoutes(api, e)
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	config.Setup(e)

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}
