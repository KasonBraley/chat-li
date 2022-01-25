package main

import (
	"github.com/KasonBraley/chat-li/socket"
	"github.com/gin-gonic/gin"
)

func SetupRouter(hub *socket.Hub) *gin.Engine {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/ws", func(c *gin.Context) {
		socket.ServeWs(hub, c.Writer, c.Request)
	})

	r.POST("/signin")
	r.POST("/signup")

	chatRooms := r.Group("/chat/room")
	{
		chatRooms.GET("/", GetAllRooms)
		chatRooms.GET("/:id", GetOneRoom)
		chatRooms.POST("/", CreateRoom)
		chatRooms.DELETE("/:id", DeleteRoom)
	}

	users := r.Group("/users")
	{
		users.GET("/")       //getAllUsers
		users.GET("/:id")    //getOneUser
		users.POST("/")      //createUser
		users.DELETE("/:id") //deleteUser
	}

	return r
}

func rootHandler(c *gin.Context) {
	c.String(200, "These are not the bugs you are looking for")
}
