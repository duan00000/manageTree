/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: 登录 验证码：https://pkg.go.dev/github.com/astaxie/beego/utils/captcha
**/
package admin

import (
	"elenewenergy/models"
	"strings"
)


type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	//获取user2表的数据验证 数据库是否链接成功
	//user2 := []models.User2{}
	//models.DB.Find(&user2)
	//beego.Info(user2)
	c.TplName = "admin/login/login.html"
}
func (c *LoginController) DoLogin() {
	//1、验证用户输入的验证码是否正确
	var flag = models.Cpt.VerifyReq(c.Ctx.Request)
	if flag {
		//2、获取表单传过来的用户名和密码
		username := strings.Trim(c.GetString("username"), "")
		password := models.Md5(strings.Trim(c.GetString("password"), ""))
		//3、去数据库匹配
		managerUn :=[]models.Manager{}
		 manager :=[]models.Manager{}
		//被屏蔽的用户登录
		models.DB.Where("username=? AND password=? AND status!=1",username,password).Find(&managerUn)
		//正常登录的用户
		models.DB.Where("username=? AND password=? AND status=1", username, password).Find(&manager)
		if len(manager) > 0 {
			//登录成功 1、保存用户信息 2、跳转到后台管理系统
			c.SetSession("userInfo", manager[0])
			c.Success("登录成功", "/")
		} else if len(managerUn) > 0 {
			//显示信息
			c.Error("该用户已被屏蔽，请联系管理员", "/login")
		}else {
				c.Error("该用户信息错误", "/login")
		}
		//c.Ctx.WriteString("正确")
	} else {
		c.Error("验证码错误", "/login")
	}

}
func (c *LoginController)LoginOut()  {
	c.DelSession("userInfo")
	c.Success("退出登录成功","/login")
}
