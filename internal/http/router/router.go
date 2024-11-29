package router

import (
	"net/http"

	"github.com/mhusainh/MIKTI-Task/internal/http/handler"
	"github.com/mhusainh/MIKTI-Task/pkg/route"
)

func PublicRoutes(movieHandler handler.MovieHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/movies",
			Handler: movieHandler.GetMovies,
		},
		{
			Method:  http.MethodGet,
			Path:    "/movies/:id",
			Handler: movieHandler.GetMovie,
		},
	}
}

func PrivateRoutes() []route.Route {
	return []route.Route{}
}
