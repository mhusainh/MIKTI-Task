package repository

import (
	"context"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	result := new(entity.User)
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Updates(&user).Error
}

func (u *userRepository) Delete(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Delete(&user).Error
}