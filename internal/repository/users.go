package repository

import (
	"context"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByusername(ctx context.Context, username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetByusername(ctx context.Context, username string) (*entity.User, error) {
	result := new(entity.User)
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}