package entity

type User struct { 
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "users"
}