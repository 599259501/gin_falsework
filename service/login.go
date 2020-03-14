package services

type RequestUser struct {
}

type BaseLogin interface {
	Auth(user RequestUser) (hasLogin bool)
}

type Login struct {
}

func NewLoginService() *Login {
	return &Login{
	}
}
func (service *Login)Auth(user RequestUser) (hasLogin bool){
	return hasLogin
}