package main

import (
	"exampleapp.com/m/v2/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 			param.ClientIP,
	// 			param.TimeStamp.Format(time.RFC1123),
	// 			param.Method,
	// 			param.Path,
	// 			param.Request.Proto,
	// 			param.StatusCode,
	// 			param.Latency,
	// 			param.Request.UserAgent(),
	// 			param.ErrorMessage,
	// 	)
	// }))
	// r.Use(gin.Recovery())
	// r.Use(middleware.CORS())


	r.Use(cors.Default())
	pong := new(controllers.PongController)

	r.GET("/ping", pong.Reply)
	r.GET("/ws", wsConnect)
	r.GET("/socket.io", wsConnect)
	return r	
}


func wsConnect(c *gin.Context)   {
	roomId := "room1"
	serveWs(c.Writer, c.Request, roomId)
}

