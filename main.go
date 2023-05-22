package main

import (
	"chatbot/middlewares"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("ChatBot"))
	router.Use(sessions.Sessions("mysession", store))

	// 这将使得在 /assets 路径下可以访问到 public 文件夹下的所有内容。
	router.Static("/assets", "./public")
	router.LoadHTMLGlob("public/html/*")

	router.GET("/", Index)
	router.POST("/chat", middlewares.CheckUser(), chat)

	router.Run(":80")
}

func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func chat(ctx *gin.Context) {
	fmt.Println("hello")

	session := sessions.Default(ctx)
	fmt.Println(session.Get("user"))

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func completion(ctx *gin.Context) {
	fmt.Println("hello")
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
