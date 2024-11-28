package builder

import (
	"github.com/mhusainh/MIKTI-Task/internal/http/handler"
	"github.com/mhusainh/MIKTI-Task/internal/http/router"
	"github.com/mhusainh/MIKTI-Task/internal/repository"
	"github.com/mhusainh/MIKTI-Task/internal/service"
	"github.com/mhusainh/MIKTI-Task/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []route.Route {
	// repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	// end

	// service
	_ = service.NewUserService(userRepository)
	movieService := service.NewMovieService(movieRepository)
	//end

	// handler
	movieHandler := handler.NewMovieHandler(movieService)
	// end

	return router.PublicRoutes(movieHandler)
}

func BuildPrivateRoutes(db *gorm.DB) []route.Route {
	_ = repository.NewUserRepository(db)
	_ = repository.NewMovieRepository(db)
	return nil
}
