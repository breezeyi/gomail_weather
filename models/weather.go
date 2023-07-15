package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// 获取天气
func GetTheWeather(city string) Data {
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
	var prompt string
	switch {
	case strings.Contains(weatherResp.Lives[0].Weather, "雨"):
		prompt = "今天可能会下雨，请注意带好雨具"
		//fmt.Println("今天可能会下雨，请注意带好雨具")
	case strings.Contains(weatherResp.Lives[0].Weather, "晴") && weatherResp.Lives[0].TemperatureFlt >= strconv.Itoa(35):
		prompt = "今天天气晴朗，气温较高，请注意防晒和补水"

		//fmt.Println("今天天气晴朗，气温较高，请注意防晒和补水")
	case strings.Contains(weatherResp.Lives[0].Weather, "晴") && weatherResp.Lives[0].TemperatureFlt <= strconv.Itoa(5):
		prompt = "今天天气晴朗，气温较低，请注意保暖"

		//fmt.Println("今天天气晴朗，气温较低，请注意保暖")
	default:
		prompt = "今天天气不错，可以出门逛逛"

		//fmt.Println("今天天气不错，可以出门逛逛")
	}

	// 输出天气信息
	fmt.Printf("城市：%s\n", weatherResp.Lives[0].City)
	fmt.Printf("天气：%s\n", weatherResp.Lives[0].Weather)
	fmt.Printf("温度：%s℃\n", weatherResp.Lives[0].Temperature)
	fmt.Printf("风力：%s级\n", weatherResp.Lives[0].WindPower)
	fmt.Printf("湿度：%s%%\n", weatherResp.Lives[0].Humidity)
	fmt.Printf("风向：%s\n", weatherResp.Lives[0].WindDirection)
	fmt.Printf("发布时间：%s\n", weatherResp.Lives[0].ReportTime)
	fmt.Println(prompt)
	//将获取的信息传递过去
	datauser := Data{
		WeatherResponse: weatherResp,
		Prompt:          prompt,
	}

	return datauser
}

type Data struct {
	WeatherResponse
	Prompt string
}

type WeatherResponse struct {
	Status string `json:"status"`   // 响应状态
	Count  string `json:"count"`    // 响应数量
	Info   string `json:"info"`     // 响应信息
	Code   string `json:"infocode"` // 响应代码

	Lives []WeatherInfo `json:"lives"` // 天气信息列表
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
