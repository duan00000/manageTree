/**
  @Author ZXQ-Administrator
  @Date 2021-12-30
  @node:
**/
package v1

import (
	"elenewenergy/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type VueTestController struct {
	beego.Controller
}

//执行登录 Post

// 1、copyrequestbody=true   1、c.Ctx.Input.RequestBody
func (c *VueTestController) DoLogin() {
	//var flag = models.Cpt.VerifyReq(c.Ctx.Request)

	user := models.User{}
	data := c.Ctx.Input.RequestBody
	beego.Info(string(data))
	json.Unmarshal(data, &user)
	c.Data["json"] = map[string]interface{}{
		"phone":    user.Phone,
		"password": user.Password,
	}
	c.ServeJSON()
}

//修改 Put
type Article struct {
	Title string
}

func (c *VueTestController) DoEdit() {
	article := Article{}
	data := c.Ctx.Input.RequestBody
	beego.Info(string(data))
	_ = json.Unmarshal(data, &article)

	c.Data["json"] = map[string]interface{}{
		"title":   article.Title,
		"success": true,
	}
	c.ServeJSON()
}

//删除数据
func (c *VueTestController) DeleteNav() {
	id, _ := c.GetInt("id")

	c.Data["json"] = map[string]interface{}{
		"id":      id,
		"message": "删除数据成功",
		"success": true,
	}
	c.ServeJSON()
}
// 解决cookie 跨域
//设置Session
func (c *VueTestController) SetUser() {
	c.SetSession("username", "张三")
	c.Data["json"] = map[string]interface{}{
		"message": "设置Session",
		"success": true,
	}
	c.ServeJSON()
}

//获取Session
func (c *VueTestController) GetUser() {
	username, ok := c.GetSession("username").(string)

	if ok {
		c.Data["json"] = map[string]interface{}{
			"username": username,
			"success":  true,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"message": "获取session失败",
			"success": false,
		}
	}
	c.ServeJSON()
}


