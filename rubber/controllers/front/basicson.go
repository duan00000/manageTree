/**
  @Author ZXQ-Administrator
  @Date 2021-09-15
  @node:
**/
package front

import (
	"elenewenergy/models"
)

type BasicsonController struct {
	BaseController
}

func (c *BasicsonController) Get() {
	need := models.Need{}
	models.DB.First(&need) //查找数据库第一条记录
	c.Data["need"] = need
	c.TplName = "front/user/basic_adddevice.html"

}

func (c *BasicsonController) DoEdit() {
	//1.获取数据库里面的数据
	need := models.Need{}
	models.DB.Find(&need)
	//2.修改数据
	err := c.ParseForm(&need)
	if err != nil {
		c.Error("form解析错误", "/basic")
		return
	}

	//4.执行保存数据
	errInfo := models.DB.Where("id=1").Save(&need).Error
	if errInfo != nil {
		c.Error("修改数据失败", "/need")
		return
	}
	c.Success("修改数据成功", "/need")
}


