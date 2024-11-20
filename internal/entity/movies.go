package entity

type Movie struct {
	ID          int64
	Title       string
	Year        int64
	Director    string
	Description string
}

func (Movie) TableName() string {
	return "movies"
}