package dto

type GetMovieByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Year        int64  `json:"year" validate:"required"`
	Director    string `json:"director" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateMovieRequest struct {
	ID          int64  `param:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Year        int64  `json:"year" validate:"required"`
	Director    string `json:"director" validate:"required"`
	Description string `json:"description" validate:"required"`
}


type DeleteMovieRequest struct {
	ID int64 `param:"id" validate:"required"`
}