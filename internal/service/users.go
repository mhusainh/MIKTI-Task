package service

import (
	"context"
	"errors"

	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"github.com/mhusainh/MIKTI-Task/internal/repository"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if user.Password != password {		
		return nil, errors.New("username atau password salah")
	}
	return user, nil
}
