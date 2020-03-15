package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop_mall/core/service"
	"shop_mall/core/support/helper"
)

var (
	config map[string]map[string]interface{}
)

func init() {
	config = make(map[string]map[string]interface{})
	// 加载session配置信息
	config["session"], _ = loadSessionCnf()
}

func loadSessionCnf() (map[string]interface{}, error) {
	var (
		cnf = make(map[string]interface{}, 3)
		err error
	)

	if cnf, err = helper.Config("./session.yaml"); err != nil {
		fmt.Println("加载session配置文件失败", err)
		panic("加载session配置文件失败")
		return cnf, err
	}

	return cnf, nil
}

type App struct {
	session *service.Session
	context *gin.Context
}

func NewApp(context *gin.Context) *App {
	app := &App{
		session: nil,
		context: context,
	}
	app.startApp()
	return app
}

func (app *App) startApp() {
	// StartApp 这里主要是处理sessionId的问题
	app.RefreshSession()
}

func (app *App) RefreshSession() {
	app.GetSession().SyncSessionIdToCookie()
}

func (app *App) GetSession() *service.Session {
	if app.session != nil {
		return app.session
	}

	// 读取配置文件，看看session下的store驱动是啥
	driverName, _ := config["session"]["driver"].(string)
	sessionName, _ := config["session"]["session_name"].(string)
	liftTime, _ := config["session"]["lifttime"].(int)
	if liftTime == 0 {
		liftTime = 3600
	}
	app.session = service.NewSession(app.context, sessionName, liftTime)
	app.session.CreateDriver(driverName)
	return app.session
}

func (app *App) GetContext() *gin.Context {
	return app.context
}
