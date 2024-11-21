package service

import (
	"context"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"github.com/mhusainh/MIKTI-Task/internal/repository"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Create(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, movie *entity.Movie) error
	Delete(ctx context.Context, movie *entity.Movie) error
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movieRepository}
}


func (s *movieService) GetAll(ctx context.Context) ([]entity.Movie, error){
	return s.movieRepository.GetAll(ctx)
}

func (s *movieService) GetByID(ctx context.Context, id int64) (*entity.Movie, error){
	return s.movieRepository.GetByID(ctx, id)
}

func (s *movieService) Create(ctx context.Context, movie *entity.Movie) error{
	return s.movieRepository.Create(ctx, movie)
}

func (s *movieService) Update(ctx context.Context, movie *entity.Movie) error{
	return s.movieRepository.Update(ctx, movie)
}

func (s *movieService) Delete(ctx context.Context, movie *entity.Movie) error{
	return s.movieRepository.Delete(ctx, movie)
}