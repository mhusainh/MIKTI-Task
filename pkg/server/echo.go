package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhusainh/MIKTI-Task/pkg/route"
)

type Server struct {
	*echo.Echo
}

func NewServer(publicRoutes, privateRoutes []route.Route) *Server {
	e := echo.New()
	v1 := e.Group("/api/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		for _, route := range privateRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return &Server{e}
}
