package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"shop_mall/core"
	"shop_mall/router"
)

func main()  {
	LoadEnv()
	Run()
}


func Run(){
	router.Render(core.GRouter)
	core.GRouter.Run()
}

func LoadEnv() {
	if err:=godotenv.Load();err!=nil{
		fmt.Println(err)
		panic("加载配置文件失败")
	}

	fmt.Println("开始加载配置文件")
}