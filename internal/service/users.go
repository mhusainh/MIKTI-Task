package service

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mhusainh/MIKTI-Task/config"
	"github.com/mhusainh/MIKTI-Task/internal/entity"
	"github.com/mhusainh/MIKTI-Task/internal/http/dto"
	"github.com/mhusainh/MIKTI-Task/internal/repository"
	"github.com/mhusainh/MIKTI-Task/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.JWTCustomClaims, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error
	Create(ctx context.Context, req dto.CreateUserRequest) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, req dto.UpdateUserRequest) error
	Delete(ctx context.Context, user *entity.User) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	RequestResetPassword(ctx context.Context, username string) error
}

type userService struct {
	cfg            *config.Config
	userRepository repository.UserRepository
}

func NewUserService(
	cfg *config.Config,
	userRepository repository.UserRepository,
) UserService {
	return &userService{cfg, userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string) (*entity.JWTCustomClaims, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username atau password salah")
	}

	if user.IsVerified == 0 {
		return nil, errors.New("Silahkan verifikasi email terlebih dahulu")
	}

	expiredTime := time.Now().Add(time.Minute * 10)

	claims := &entity.JWTCustomClaims{
		Username: user.Username,
		FullName: user.FullName,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "movie-app",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	return claims, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	user := new(entity.User)
	user.Username = req.Username
	user.FullName = req.FullName
	user.Role = "User"
	user.ResetPasswordToken = utils.RandomString(16)
	user.VerifyEmailToken = utils.RandomString(16)
	user.IsVerified = 0

	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	templatePath := "./templates/email/verify-email.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var ReplacerEmail = struct {
		Token string
	}{
		Token: user.VerifyEmailToken,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.SMTPConfig.Username)
	m.SetHeader("To", user.Username)
	m.SetHeader("Subject", "Verifikasi Email !")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		s.cfg.SMTPConfig.Host,
		s.cfg.SMTPConfig.Port,
		s.cfg.SMTPConfig.Username,
		s.cfg.SMTPConfig.Password,
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}

func (s *userService) Create(ctx context.Context, req dto.CreateUserRequest) error {
	user := new(entity.User)
	user.Username = req.Username
	user.FullName = req.FullName
	user.Role = req.Role

	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.userRepository.Create(ctx, user)
}

func (s *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) error {
	user, err := s.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	return s.userRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, user *entity.User) error {
	return s.userRepository.Delete(ctx, user)
}

func (s *userService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	user, err := s.userRepository.GetByResetPasswordToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token reset password salah")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.Update(ctx, user)
}

func (s *userService) RequestResetPassword(ctx context.Context, username string) error {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return errors.New("username tersebut tidak ditemukan")
	}

	templatePath := "./templates/email/reset-password.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var ReplacerEmail = struct {
		Token string
	}{
		Token: user.ResetPasswordToken,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.SMTPConfig.Username)
	m.SetHeader("To", user.Username)
	m.SetHeader("Subject", "Reset Password Request !")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		s.cfg.SMTPConfig.Host,
		s.cfg.SMTPConfig.Port,
		s.cfg.SMTPConfig.Username,
		s.cfg.SMTPConfig.Password,
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error {
	user, err := s.userRepository.GetByVerifyEmailToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token verifikasi email salah")
	}
	user.IsVerified = 1
	return s.userRepository.Update(ctx, user)
}