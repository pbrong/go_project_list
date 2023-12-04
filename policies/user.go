package policies

type UserPolicyIF interface {
	CanLogin(userID int64) bool
	CanRegister(userID int64) bool
}

type userPolicy struct{}

// CanLogin implements UserPolicyIF
func (*userPolicy) CanLogin(userID int64) bool {
	panic("unimplemented")
}

// CanRegister implements UserPolicyIF
func (*userPolicy) CanRegister(userID int64) bool {
	panic("unimplemented")
}
