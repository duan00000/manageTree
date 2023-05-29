/**
  @Author ZXQ-Administrator
  @Date 2021-09-25
  @node: api接口登录页面  V2版本
**/
package v2

import "github.com/astaxie/beego"

type V2Controller struct {
	beego.Controller
}

func (c *V2Controller) Get() {
	c.Ctx.WriteString("ap1 v2")
}
