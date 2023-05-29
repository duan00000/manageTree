/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: 前端登录页面
**/
package index

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName="index/login.html"
}