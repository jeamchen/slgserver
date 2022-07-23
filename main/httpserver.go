package main

import (
	"log"
	"slgserver/config"
	"slgserver/db"
	"slgserver/server/httpserver/controller"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func main() {

	db.TestDB()

	e := echo.New()
	e.Use(mw.Recover())

	g := e.Group("")
	new(controller.AccountController).RegisterRoutes(g)
	e.Server.Addr = getHttpAddr()
	log.Fatal(e.StartServer(e.Server))
}

func getHttpAddr() string {
	host := config.File.MustValue("httpserver", "host", "127.0.0.1")
	port := config.File.MustValue("httpserver", "port", "8088")
	return host + ":" + port
}
