package dao

var UserDao UserDaoIF

func InitDAO() {
	UserDao = new(userDao)
}
