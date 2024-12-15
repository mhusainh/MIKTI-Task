package dto

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type GetUserByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type UpdateUserRequest struct {
	ID       int64  `param:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type DeleteUserRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type ResetPasswordRequest struct {
	Token    string `param:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestResetPassword struct {
	Username string `json:"username" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `param:"token" validate:"required"`
}