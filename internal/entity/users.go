package entity

type User struct {
	ID       int64
	Username string
	Password string
	FullName string
}

func (User) TableName() string {
	return "users"
}