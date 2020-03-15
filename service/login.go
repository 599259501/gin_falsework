package service

import "shop_mall/core"

type RequestUser struct {
	name     string
	password string
}

type Login struct {
	core.AppUnit
}

func NewLoginService(app *core.App) *Login {
	login := &Login{}
	login.BindApp(app)

	return login
}

func (service *Login) Auth(user RequestUser) (hasLogin bool) {
	var userInfo interface{}
	userInfo = service.GetApp().GetSession().Get()
	if userInfo != nil {
		return false
	}

	return true
}
