/**
  @Author ZXQ-Administrator
  @Date 2021-09-15
  @node:
**/
package admin

import (
	"elenewenergy/models"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Get() {
	setting := models.Setting{}
	models.DB.First(&setting) //查找数据库第一条记录
	c.Data["setting"] = setting
	c.TplName = "admin/setting/index.html"

}

func (c *SettingController) DoEdit() {
	//1.获取数据库里面的数据
	setting := models.Setting{}
	models.DB.Find(&setting)
	//2.修改数据
	err := c.ParseForm(&setting)
	if err != nil {
		c.Error("form解析错误", "/setting/doEdit")
		return
	}
	//3.上传图片吗site_logo no_picture
	siteLogo, errSiteLogo := c.UploadImg("site_logo", "site_logo")
	if len(siteLogo) > 0 && errSiteLogo == nil {
		setting.SiteLogo = siteLogo
	}
	noPicture, err2 := c.UploadImg("no_picture", "no_picture")
	if len(noPicture) > 0 && err2 == nil {
		setting.NoPicture = noPicture
	}
	//4.执行保存数据
	errInfo := models.DB.Where("id=1").Save(&setting).Error
	if errInfo != nil {
		c.Error("修改数据失败", "/setting")
		return
	}
	c.Success("修改数据成功", "/setting")
}
