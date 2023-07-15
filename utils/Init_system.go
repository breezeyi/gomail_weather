package utils

import (
	"log"

	"github.com/spf13/viper"
)

func Config() {
	viper.SetConfigFile("./config/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误！" + err.Error())
	}
	log.Println("配置文件加载成功！")
}
