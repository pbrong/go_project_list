package entity

type UserType string

const (
	UserTypeAdmin  UserType = "admin"
	UserTypeNormal UserType = "normal"
)

type User struct {
	UserType UserType
	UserID   int64
	UserName string
}

func (u *User) IsAdmin() bool {
	return u.UserType == UserTypeAdmin
}
