/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: 轮播图
**/
package admin

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"strconv"
)

type FocusController struct {
	BaseController
}

func (c *FocusController) Get() {
	focus := []models.Focus{}
	models.DB.Find(&focus)
	c.Data["focusList"] = focus

	c.TplName = "admin/focus/index.html"
}
func (c *FocusController) Add() {

	c.TplName = "admin/focus/add.html"
}
func (c *FocusController) DoAdd() {
	//获取数据
	focusType, errType := c.GetInt("focus_type")
	focusSort, errSort := c.GetInt("sort")
	focusStatus, errStatus := c.GetInt("status")
	title := c.GetString("title")
	link := c.GetString("link")

	if errType != nil || errSort != nil || errStatus != nil {
		c.Error("数据请求不合法", "/focus")
	}
	// 执行图片上传 保存名称为 title_时间戳
	focusImgSrc, errImg := c.UploadImg("focus_img", title)
	if errImg != nil {
		beego.Error(errImg)
	}
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      focusSort,
		Status:    focusStatus,
		AddTime:   int(models.GetUnix()),
	}
	models.DB.Create(&focus)
	c.Success("增加轮播图成功", "/focus")
}
func (c *FocusController) Edit() {
	//获取信息
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/manager")
		return
	}
	//获取id对应的轮播图信息
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.Data["focus"] = focus
	c.TplName = "admin/focus/edit.html"
}
func (c *FocusController) DoEdit() {
	id,errId:=c.GetInt("id")
	focusType, errType := c.GetInt("focus_type")
	focusSort, errSort := c.GetInt("sort")
	focusStatus, errStatus := c.GetInt("status")
	title := c.GetString("title")
	link := c.GetString("link")
	if errType != nil || errStatus != nil || errId != nil {
		c.Error("数据请求不合法", "/focus")
	}
	if errSort !=nil{
		c.Error("排序表单里面输入的数据不合法","/focus/edit?id="+strconv.Itoa(id))
	}
	focus:=models.Focus{Id:id}
	models.DB.Find(&focus)

	focus.Title=title
	focus.FocusType=focusType
	focus.Link=link
	focus.Status=focusStatus
	focus.Sort=focusSort
	// 执行图片上传 保存名称为 title_时间戳
	focusImgSrc, errImg := c.UploadImg("focus_img", title)
	if errImg != nil {
		beego.Error(errImg)
	}
	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	//执行修改
	error:=models.DB.Save(&focus).Error
	if error!=nil{
		c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改数据成功", "/focus")


}
func (c *FocusController) Delete() {
	id,errId:=c.GetInt("id")
	if errId !=nil{
		c.Error("传入参数错误","/focus")
		return
	}
	focus:=models.Focus{Id:id}
	models.DB.Delete(&focus)
	c.Success("删除轮播图成功", "/focus")

}
