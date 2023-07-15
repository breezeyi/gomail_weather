package routers

import (
	"Weather/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Use(Cors()) // 添加 Cors 中间件,允许跨域问题
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./templates/*")
	//选择界面
	r.GET("/add", func(this *gin.Context) {
		this.HTML(200, "index.html", "")
	})

	//测试地址
	r.GET("/test", controllers.Test)

	//获取提交的用户信息
	r.POST("/submit", controllers.ReceiveData)
	//测试获取天气
	r.POST("/weather", controllers.GetTheWeather)
	return r
}

// 中间件设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}