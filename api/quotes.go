package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// Quote -
func (api *API) Quote(c echo.Context) error {
	quote := api.quotes.RandomQuote()
	return c.String(http.StatusOK, "Welcome "+quote.Text+"!")
}

// SecretQuote -
func (api *API) SecretQuote(c echo.Context) error {
	quote := api.quotes.RandomQuote()
	return c.String(http.StatusOK, "Welcome "+quote.Text+"!")
}
