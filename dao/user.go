package dao

import "test/entity"

type UserDaoIF interface {
	GetUser(userID int64) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userDao struct{}

// CreateUser implements UserDaoIF
func (*userDao) CreateUser(user *entity.User) error {
	panic("unimplemented")
}

// GetUser implements UserDaoIF
func (*userDao) GetUser(userID int64) (*entity.User, error) {
	panic("unimplemented")
}
