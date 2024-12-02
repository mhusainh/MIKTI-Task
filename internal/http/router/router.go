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
		{
			Method:  http.MethodPost,
			Path:    "/movies",
			Handler: movieHandler.CreateMovie,
		},
		{
			Method:  http.MethodPut,
			Path:    "/movies/:id",
			Handler: movieHandler.UpdateMovie,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/movies/:id",
			Handler: movieHandler.DeleteMovie,
		},
	}
}

func PrivateRoutes() []route.Route {
	return []route.Route{}
}
