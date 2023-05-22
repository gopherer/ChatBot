package middlewares

import (
	"chatbot/config"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cli := config.Client{}
		err := ctx.BindJSON(&cli)
		if err != nil {
			return
		}
		if cli.UserName != "root" || cli.Password != "root" {
			// 用户未登录
			ctx.Redirect(302, "/")
			fmt.Println("登录失败")
			ctx.Abort()
			return
		}
		session := sessions.Default(ctx)
		session.Set("user", cli.UserName)
		session.Save()
	}
}
