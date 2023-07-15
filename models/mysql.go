package models

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Event struct {
	gorm.Model
	Email   string //收件人邮箱
	Address string //收件人地址
}

// 城市编码
type Encode struct {
	Name     string
	Adcode   string
	Aitycode string
}

type Statement struct {
	Id      int
	Chinese string
	English string
}

var DB *gorm.DB

// 定义一个方法程序开始时，
// 初始化数据并且连接数据库拿到数据存放在redis中
func InitializeMYSQL() (err error) {
	dns := viper.GetString("mysql.dns")
	//创建orm连接池
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		//设置日志输出条件
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}
	//拿到底层的MySQL连接池DB
	sqldb, err := DB.DB()
	if err != nil {
		return err
	}
	//测试是否连接上数据库
	err = sqldb.Ping()
	if err != nil {
		return err
	}
	//连接上了设置连接池的设置(默认的设置就可以应用于平常的需求)
	//设置连接池的数量时间等
	sqldb.SetMaxOpenConns(10)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxIdleTime(time.Hour)
	sqldb.SetConnMaxLifetime(time.Hour)

	//设置好后顺便创建一个数据表
	//DB.AutoMigrate(&Event{})
	return
}

// 将接受到的数据存入数据库中
func CreateUser(email, addrs string) {
	user := Event{
		Email:   email,
		Address: addrs,
	}
	err := DB.Create(&user).Error
	if err != nil {
		log.Println("用户" + email + "存入错误！")
	}
}

// 将获取的城市查询对应的编码
func FincCode(addrs string) *Encode {
	usercode := &Encode{}
	err := DB.Model(&Encode{}).Where("name=?", addrs).First(usercode).Error
	if err != nil {
		log.Println(addrs + "城市编码查询错误！err:" + err.Error())
	}
	return usercode
}

// 根据数字编号查询拿出对应的值
func FindNumber(id int) *Statement {
	meassage := &Statement{}
	err := DB.Model(&Statement{}).Where("id=?", id).First(meassage).Error
	if err != nil {
		log.Println("每日一句获取失败！" + err.Error())
	}
	return meassage
}
