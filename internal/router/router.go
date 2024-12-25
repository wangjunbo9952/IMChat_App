package router

import (
	api "IMChat_App/api/websocket"
	"IMChat_App/internal/auth"
	"IMChat_App/internal/router/middleware"
	"IMChat_App/internal/websocket"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.Static("/css", "./web/static/css")
	r.Static("/js", "./web/static/js")
	r.LoadHTMLFiles("web/templates/login.html", "web/templates/chat.html")

	publicGroup := r.Group("")
	{
		// 页面路由
		publicGroup.GET("/login", auth.LoadLoginPage)
		publicGroup.GET("/chat/:account", auth.LoadChatPage)

		// 认证相关API
		publicGroup.POST("/user/register", auth.Register)
		publicGroup.POST("/user/login", auth.Login)
		publicGroup.GET("/ws", websocket.RunSocket)
	}

	// 需要JWT认证的路由组
	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware())        // JWT认证中间件
	authorized.Use(middleware.AutoRefreshMiddleware()) // JWT刷新中间件
	{

		// 用户相关API
		userGroup := authorized.Group("/user")
		{
			// 好友相关
			userGroup.POST("/adduser", auth.AddUser)
			userGroup.GET("/searchuser", auth.SearchUser)
			userGroup.GET("/getfriends", auth.GetFriendsList)

			// 消息相关
			userGroup.POST("/history", api.GetHistory)
		}
	}

	r.Run(":9090")
}
