package main

import (
	"disgord/server"

	"github.com/labstack/echo/v4"
)

const (
	ip   = "127.0.0.1"
	port = 5002
)

func main() {
	srv := server.NewServer(ip, port)
	go srv.Serve()

	e := echo.New()

	e.GET("/join", srv.Join)
	e.Logger.Fatal(e.Start(":8080"))
}
