package server

import "github.com/labstack/echo/v4"

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	return &Server{e}
}