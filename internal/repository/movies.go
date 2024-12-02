package repository

import (
	"context"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll(ctx context.Context) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Create(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, movie *entity.Movie) error
	Delete(ctx context.Context, movie *entity.Movie) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) GetAll(ctx context.Context) ([]entity.Movie, error) {
	result := make([]entity.Movie, 0)
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *movieRepository) GetByID(ctx context.Context, id int64) (*entity.Movie, error) {
	result := new(entity.Movie)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *movieRepository) Create(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Create(&movie).Error
}

func (r *movieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Updates(&movie).Error
}

func (r *movieRepository) Delete(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Delete(&movie).Error
}
