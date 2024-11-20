package repository

import (
	"context"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll(ctx context.Context) ([]*entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Create(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, id int64, data map[string]interface{}) error
	Delete(ctx context.Context, id int64) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (m *movieRepository) GetAll(ctx context.Context) ([]*entity.Movie, error) {
	var result []*entity.Movie
	if err := m.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (m *movieRepository) GetByID(ctx context.Context, id int64) (*entity.Movie, error) {
	result := new(entity.Movie)
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (m *movieRepository) Create(ctx context.Context, movie *entity.Movie) error {
	if err := m.db.WithContext(ctx).Create(movie).Error; err != nil {
		return err
	}
	return nil
}

func (m *movieRepository) Update(ctx context.Context, id int64, data map[string]interface{}) error {
	if err := m.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (m *movieRepository) Delete(ctx context.Context, id int64) error {
	if err := m.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Movie{}).Error; err != nil{
		return err
	}
	return nil
}
