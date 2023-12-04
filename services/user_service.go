package services

import "test/entity"

type UserServiceIF interface {
	UserLogin(u *entity.User) error
}

type userService struct{}

// UserLogin implements UserServiceIF
func (*userService) UserLogin(u *entity.User) error {
	panic("unimplemented")
}
