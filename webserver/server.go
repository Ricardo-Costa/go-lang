package webserver

import (
	"github.com/Ricardo-Costa/go-lang/seeders"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {

}

func NewWebServer () *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo .Context) error {
	return c.JSON(http.StatusOK, seeders.Products)
}