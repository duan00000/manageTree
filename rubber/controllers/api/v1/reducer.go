/**
  @Author ZXQ-Administrator
  @Date 2021-10-27
  @node: api接口登录页面  V1版本 使用中
**/
package v1

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
)

type ReducerController struct {
	beego.Controller
}


//导航的api接口
func (c *ReducerController) Reducer() {
	reducer := []models.Reducer{}
	models.DB.Find(&reducer)
	c.Data["json"] = map[string]interface{}{
		"result": reducer,
		"success": true,
	}
	c.ServeJSON()
}