package router

import (
	"IMChat_App/internal/handler"
	"IMChat_App/internal/middleware"
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
		publicGroup.GET("/login", handler.LoadLoginPage)
		publicGroup.GET("/chat/:account", handler.LoadChatPage)

		// 认证相关API
		publicGroup.POST("/user/register", handler.Register)
		publicGroup.POST("/user/login", handler.Login)
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
			userGroup.POST("/adduser", handler.AddUser)
			userGroup.GET("/searchuser", handler.SearchUser)
			userGroup.GET("/getfriends", handler.GetFriendsList)

			// 消息相关
			userGroup.POST("/history", handler.GetMsg)
		}
	}

	r.Run(":9090")
}
