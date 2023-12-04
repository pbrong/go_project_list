package policies

var UserPolicy UserPolicyIF

func Init() {
	UserPolicy = new(userPolicy)
}
