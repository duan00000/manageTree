/**
  @Author ZXQ-Administrator
  @Date 2021-09-24
  @node: 判断前端用户是否登录
**/
package middleware

import (
	"elenewenergy/models"
	"github.com/astaxie/beego/context"
)

func FrontAuth(ctx *context.Context) {
	//判断前端用户有没有登录
	user := models.User{}
	models.Cookie.Get(ctx, "userinfo", &user)
	if len(user.Phone) != 11 {
		ctx.Redirect(302, "/pass/login")
	}
}
