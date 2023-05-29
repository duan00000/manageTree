/**
  @Author ZXQ-Administrator
  @Date 2021-09-05
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"strconv"
	"strings"
)

type GoodsTypeController struct {
	BaseController
}

func (c *GoodsTypeController) Get() {

	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.TplName = "admin/goodsType/index.html"
}

func (c *GoodsTypeController) Add() {
	c.TplName = "admin/goodsType/add.html"
}

func (c *GoodsTypeController) DoAdd() {
	title := strings.Trim(c.GetString("title"), " ")
	description := strings.Trim(c.GetString("description"), " ")
	status, errStatus := c.GetInt("status")
	if errStatus != nil {
		c.Error("传入参数不正确", "/goodsType/add")
		return
	}
	if title == "" {
		c.Error("标题不能为空", "/goodsType/add")
		return
	}
	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(models.GetUnix())
	err := models.DB.Create(&goodsType).Error
	if err != nil {
		c.Error("增加商品类型失败", "/goodsType/add")
	} else {
		c.Success("增加商品类型成功", "/goodsType")
	}
}

func (c *GoodsTypeController) Edit() {
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}

	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	c.TplName = "admin/goodsType/edit.html"
}

func (c *GoodsTypeController) DoEdit() {

	id, errId := c.GetInt("id")

	title := strings.Trim(c.GetString("title"), " ")
	description := strings.Trim(c.GetString("description"), " ")
	status, errStatus := c.GetInt("status")
	if errId != nil || errStatus != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		return
	}
	//修改
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err := models.DB.Save(&goodsType).Error
	if err != nil {
		c.Error("修改数据失败", "/goodsType/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/goodsType")
	}

}

func (c *GoodsTypeController) Delete() {
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Delete(&goodsType)
	c.Success("删除数据成功", "/goodsType")

}
