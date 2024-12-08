package router

import (
	"IMChat_App/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() {
	r := gin.Default()
	r.Use(Cors())

	r.LoadHTMLFiles("web/page/chat.html")

	r.GET("/chat/:account", handler.LoadChatPage)
	r.POST("/user/register", handler.Register)
	r.POST("/user/login", handler.Login)
	r.POST("/user/adduser", handler.AddUser)
	r.GET("/user/searchuser", handler.SearchUser)

	r.Run(":9090")
}

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		c.Next()
	}
}
