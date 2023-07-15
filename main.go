package main

import (
	"Weather/models"
	"Weather/routers"
	"Weather/utils"
	"log"
)

func main() {
	//初始化配置文件
	utils.Config()
	//初始化MySQL
	err := models.InitializeMYSQL()
	if err != nil {
		log.Fatalf("MySQL start err :" + err.Error())
	}
	log.Println("mysql start OK!")
	//启动服务器并且开启端口号
	r := routers.Init()
	if err := r.Run(); err != nil {
		log.Fatalln("端口号启动失败！")
	}
}
