/**
  @Author ZXQ-Administrator
  @Date 2021-09-04
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"strconv"
)

type GoodsCateController struct {
	BaseController
}

func (c *GoodsCateController) Get() {
	goodsCate := []models.GoodsCate{}
	//预加载 将Pid=0的数据挂载到GoodsCateItem
	models.DB.Preload("GoodsCateItem").Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/index.html"

}
func (c *GoodsCateController) Add() {
	//加载Pid=0的数据
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/add.html"

}
func (c *GoodsCateController) DoAdd() {
	pid, errPid := c.GetInt("pid")
	title := c.GetString("title")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, errSort := c.GetInt("sort")
	status, errStatus := c.GetInt("status")

	if errPid != nil || errStatus != nil {
		c.Error("传入参数错误", "/goodsCate/add")
		return
	}
	if errSort != nil {
		c.Error("排序值必须为整数", "/goodsCate/add")
		return
	}

	goodsCateImgSrc, errImg := c.UploadImg("cate_img", title)
	if errImg != nil {
		beego.Error(errImg)
	}
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		Link:        link,
		Sort:        sort,
		SubTitle:    subTitle,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     goodsCateImgSrc,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		c.Error("增加失败", "/goodsCate/add")
		return
	}
	c.Success("增加商品分类成功", "/goodsCate")
}

func (c *GoodsCateController) Edit() {
	//获取当前数据
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	c.Data["goodsCate"] = goodsCate
	// 顶级分类
	topGoodsCate := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&topGoodsCate)
	c.Data["goodsCateList"] = topGoodsCate

	c.TplName = "admin/goodsCate/edit.html"
}
func (c *GoodsCateController) DoEdit() {
	id, errId := c.GetInt("id")
	title := c.GetString("title")
	pid, errPid := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, errSort := c.GetInt("sort")
	status, errStatus := c.GetInt("status")

	if errId != nil || errPid != nil || errStatus != nil {
		c.Error("传入参数类型不正确", "/goodsCate")
		return
	}
	if errSort != nil {
		c.Error("排序值必须是整数", "/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	goodsCateImgSrc, _ := c.UploadImg("cate_img", title)

	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if goodsCateImgSrc != "" {
		goodsCate.CateImg = goodsCateImgSrc
	}
	err := models.DB.Save(&goodsCate).Error
	if err != nil {
		c.Error("修改失败", "/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/goodsCate")

}
func (c *GoodsCateController) Delete() {
	id,err:=c.GetInt("id")
	if err !=nil{
		c.Error("传入参数类型不正确", "/goodsCate")
		return
	}
	goodsCate:=models.GoodsCate{Id:id}
	models.DB.Find(&goodsCate)
	//顶级分类
	if goodsCate.Pid == 0{
		goodsCatePid := []models.GoodsCate{}
		models.DB.Where("pid=?", goodsCate.Id).Find(&goodsCatePid)
		if len(goodsCatePid) > 0 {
			c.Error("当前分类下面还子分类，无法删除", "/goodsCate")
			return
		}
	}
	goodsCateSub := models.GoodsCate{Id: id}
	models.DB.Delete(&goodsCateSub)
	c.Success("删除成功", "/goodsCate")
}
