package models

// User .. user data for authentication
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

// TableName .. todo table name
func (u *User) TableName() string {
	return "users"
}
