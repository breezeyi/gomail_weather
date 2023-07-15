package controllers

import (
	"Weather/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 测试
func Test(this *gin.Context) {
	this.String(200, "测试")
}

// 获取天气
func GetTheWeather(this *gin.Context) {
	city := this.PostForm("city")
	// 高德天气API接口地址
	apiUrl := "https://restapi.amap.com/v3/weather/weatherInfo"

	// 构造请求参数
	params := url.Values{}
	params.Set("key", viper.GetString("weather.key"))
	params.Set("city", city)

	// 发送HTTP GET请求
	resp, err := http.Get(apiUrl + "?" + params.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 解析JSON响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var weatherResp WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		log.Fatal(err)
	}

	// 根据天气情况输出相应的提示
	switch {
	case strings.Contains(weatherResp.Lives[0].Weather, "雨"):
		fmt.Println("今天可能会下雨，请注意带好雨具")
	case strings.Contains(weatherResp.Lives[0].Weather, "晴") && weatherResp.Lives[0].TemperatureFlt >= strconv.Itoa(35):
		fmt.Println("今天天气晴朗，气温较高，请注意防晒和补水")
	case strings.Contains(weatherResp.Lives[0].Weather, "晴") && weatherResp.Lives[0].TemperatureFlt <= strconv.Itoa(5):
		fmt.Println("今天天气晴朗，气温较低，请注意保暖")
	default:
		fmt.Println("今天天气不错，可以出门逛逛")
	}

	// 输出天气信息
	fmt.Printf("城市：%s\n", weatherResp.Lives[0].City)
	fmt.Printf("天气：%s\n", weatherResp.Lives[0].Weather)
	fmt.Printf("温度：%s℃\n", weatherResp.Lives[0].Temperature)
	fmt.Printf("风力：%s级\n", weatherResp.Lives[0].WindPower)
	fmt.Printf("湿度：%s%%\n", weatherResp.Lives[0].Humidity)
	fmt.Printf("风向：%s\n", weatherResp.Lives[0].WindDirection)
	fmt.Printf("发布时间：%s\n", weatherResp.Lives[0].ReportTime)

	this.String(200, string(body))
	fmt.Println("++++++++=====>")
}

type WeatherResponse struct {
	Status string        `json:"status"`   // 响应状态
	Count  string        `json:"count"`    // 响应数量
	Info   string        `json:"info"`     // 响应信息
	Code   string        `json:"infocode"` // 响应代码
	Lives  []WeatherInfo `json:"lives"`    // 天气信息列表
}

type WeatherInfo struct {
	Province       string `json:"province"`          // 省份
	City           string `json:"city"`              // 城市
	Adcode         string `json:"adcode"`            // 区域编码
	Weather        string `json:"weather"`           // 天气情况
	Temperature    string `json:"temperature"`       // 温度
	WindDirection  string `json:"winddirection"`     // 风向
	WindPower      string `json:"windpower"`         // 风力
	Humidity       string `json:"humidity"`          // 湿度
	ReportTime     string `json:"reporttime"`        // 发布时间
	TemperatureFlt string `json:"temperature_float"` // 温度（浮点数）
	HumidityFlt    string `json:"humidity_float"`    // 湿度（浮点数）
}

func ReceiveData(this *gin.Context) {
	email := this.PostForm("email")
	city := this.PostForm("city")
	district := this.PostForm("district")
	// fmt.Println(email)
	// fmt.Println(city)
	// fmt.Println(district)

	//确认最终获取天气的地址
	var addrs string
	if district != "choose" {
		addrs = district
	} else {
		addrs = city
	}
	//this.String(200, "提交成功！")
	go models.OneGomail(email)
	go models.CreateUser(email, addrs)
	//查询对应的城市code编码
	usercode := models.FincCode(addrs)
	// fmt.Println(usercode)
	// 获取当前时间
	now := time.Now()

	// 设置早上八点的时间
	eightAM := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())

	// 判断当前时间是否大于早上八点
	if now.After(eightAM) {
		fmt.Println("当前时间大于早上八点")
		go models.ErrGomail(email)
		go models.Gomail(usercode, email)
	} else {
		fmt.Println("当前时间小于等于早上八点")
		go models.TimeGomail(usercode, email)
		go models.Gomail(usercode, email)
	}
	//go models.Gomail(usercode, email)
	this.String(200, "提交成功！")
}
