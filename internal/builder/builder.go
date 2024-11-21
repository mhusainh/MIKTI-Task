package builder

import (
	"github.com/mhusainh/MIKTI-Task/internal/repository"
	"github.com/mhusainh/MIKTI-Task/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []route.Route {
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	return nil
}

func BuildPrivateRoutes(db *gorm.DB) []route.Route {
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	return nil
}
