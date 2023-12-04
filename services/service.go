package services

var UserService UserServiceIF

func Init() {
	UserService = new(userService)
}
