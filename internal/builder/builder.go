package builder

import (
	"github.com/mhusainh/MIKTI-Task/config"
	"github.com/mhusainh/MIKTI-Task/internal/http/handler"
	"github.com/mhusainh/MIKTI-Task/internal/http/router"
	"github.com/mhusainh/MIKTI-Task/internal/repository"
	"github.com/mhusainh/MIKTI-Task/internal/service"
	"github.com/mhusainh/MIKTI-Task/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	// repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	// end

	// service
	userService := service.NewUserService(userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	movieService := service.NewMovieService(movieRepository)
	//end

	// handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	// end

	return router.PublicRoutes(movieHandler, userHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)

	userService := service.NewUserService(userRepository)
	movieService := service.NewMovieService(movieRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)

	userHandler := handler.NewUserHandler(tokenService, userService)
	movieHandler := handler.NewMovieHandler(movieService)
	return router.PrivateRoutes(movieHandler, userHandler)
}
