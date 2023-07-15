package models

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Gomail(usercode *Encode, email string) {

	//fmt.Println(email)
	for {
		data := GetTheWeather(usercode.Adcode)
		seed := rand.NewSource(time.Now().UnixNano())

		// 初始化随机数生成器
		r := rand.New(seed)

		// 生成 1-14 之间的随机整数序列
		randomNums := r.Intn(14) + 1
		//查询出对应数字编号的值
		msg := FindNumber(randomNums)
		fmt.Println(msg)
		now := time.Now()
		//计算下一次的时间
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 8, 0, 0, 0, now.Location())
		duration := next.Sub(now)
		time.Sleep(duration)

		// 假设这里是从 API 中获取到的天气信息
		city := data.WeatherResponse.Lives[0].City
		weather := data.WeatherResponse.Lives[0].Weather
		temperature := data.WeatherResponse.Lives[0].Temperature
		windPower := data.WeatherResponse.Lives[0].WindPower
		humidity := data.WeatherResponse.Lives[0].Humidity
		windDirection := data.WeatherResponse.Lives[0].WindDirection

		mailTo := []string{
			email,
		}
		//邮件主题
		subject := city + "天气预报"
		// 邮件正文
		body := "<html><body><p>" +
			"城市：" + city + "<br>" +
			"天气：" + weather + "<br>" +
			"温度：" + temperature + "℃<br>" +
			"风力：" + windPower + "级<br>" +
			"湿度：" + humidity + "%<br>" +
			"风向：" + windDirection + "<br>" +
			"今日小意见:" + data.Prompt + "<br>" +
			"每日一句：" + "<br>" +
			msg.English + "<br>" +
			msg.Chinese + "<br>" +
			"</p></body></html>"

		err := SendMail(mailTo, subject, body)
		if err != nil {
			log.Println(err)
			fmt.Println("send fail")
		}
		fmt.Println("发送成功！")
	}

}

func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": viper.GetString("smtp.user"),
		"pass": viper.GetString("smtp.pass"),
		"host": viper.GetString("smtp.host"),
		"port": viper.GetString("smtp.port"),
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "晚风漪天气")) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                    //发送给多个用户
	m.SetHeader("Subject", subject)                                 //设置邮件主题
	m.SetBody("text/html", body)                                    //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}

// 第一次提示
func OneGomail(email string) {
	// now := time.Now()
	// //计算下一次的时间
	// next := time.Date(now.Year(), now.Month(), now.Day(), 0, 5, 0, 0, now.Location())
	// duration := next.Sub(now)
	// time.Sleep(duration)

	mailTo := []string{
		email,
	}
	//邮件主题
	subject := "订阅天气成功"
	// 邮件正文
	body := "<html><body><p>" +
		"您好，您已成功订阅天气提醒服务。每天早上 8 点，我们将会向您发送当天的天气情况。谢谢您的使用！" +
		"</p></body></html>"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")

	}
	fmt.Println("第一次订阅提示返送成功！")

}

// 订阅的时候错过早上八点时
func ErrGomail(email string) {
	// now := time.Now()
	// //计算下一次的时间
	// next := time.Date(now.Year(), now.Month(), now.Day(), 0, 5, 0, 0, now.Location())
	// duration := next.Sub(now)
	// time.Sleep(duration)

	mailTo := []string{
		email,
	}
	//邮件主题
	subject := "订阅提示"
	// 邮件正文
	body := "<html><body><p>" +
		"您好，我们发现您错过了今天早上八点的订阅消息。我们将会在第二天继续给您推送信息。如果您需要帮助或有任何疑问，请随时联系我们。" +
		"</p></body></html>"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")

	}
	fmt.Println("第一次订阅时间错过提示返送成功！")

}

// 当天订阅时小于八点
func TimeGomail(usercode *Encode, email string) {
	data := GetTheWeather(usercode.Adcode)
	//fmt.Println(email)
	seed := rand.NewSource(time.Now().UnixNano())

	// 初始化随机数生成器
	r := rand.New(seed)

	// 生成 1-14 之间的随机整数序列
	randomNums := r.Intn(14) + 1
	//查询出对应数字编号的值
	msg := FindNumber(randomNums)
	//fmt.Println(msg)
	now := time.Now()
	//计算下一次的时间
	next := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
	duration := next.Sub(now)
	time.Sleep(duration)

	// 假设这里是从 API 中获取到的天气信息
	city := data.WeatherResponse.Lives[0].City
	weather := data.WeatherResponse.Lives[0].Weather
	temperature := data.WeatherResponse.Lives[0].Temperature
	windPower := data.WeatherResponse.Lives[0].WindPower
	humidity := data.WeatherResponse.Lives[0].Humidity
	windDirection := data.WeatherResponse.Lives[0].WindDirection

	mailTo := []string{
		email,
	}
	//邮件主题
	subject := city + "天气预报"
	// 邮件正文
	body := "<html><body><p>" +
		"城市：" + city + "<br>" +
		"天气：" + weather + "<br>" +
		"温度：" + temperature + "℃<br>" +
		"风力：" + windPower + "级<br>" +
		"湿度：" + humidity + "%<br>" +
		"风向：" + windDirection + "<br>" +
		"今日小意见:" + data.Prompt + "<br>" +
		"每日一句：" + "<br>" +
		msg.English + "<br>" +
		msg.Chinese + "<br>" +
		"</p></body></html>"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")

	}
	fmt.Println("当天订阅八点信息发送成功！")

}
