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

type GoodsTypeAttrController struct {
	BaseController
}

func (c *GoodsTypeAttrController) Get() {

	cateId, errCateId := c.GetInt("cate_id")
	if errCateId != nil {
		c.Error("非法请求", "/goodsType")
	}
	//获取当前的类型
	goodsType := models.GoodsType{Id: cateId}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType

	//查询当前类型下面的商品类型属性
	goodsTypeAttr := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttr)
	c.Data["goodsTypeAttrList"] = goodsTypeAttr

	c.TplName = "admin/goodsTypeAttribute/index.html"
}

func (c *GoodsTypeAttrController) Add() {

	cateId, errCateId := c.GetInt("cate_id")
	if errCateId != nil {
		c.Error("非法请求", "/goodsType")
	}

	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.Data["cateId"] = cateId
	c.TplName = "admin/goodsTypeAttribute/add.html"
}

func (c *GoodsTypeAttrController) DoAdd() {

	title := c.GetString("title")
	cateId, errCateId := c.GetInt("cate_id")
	attrType, errAttrType := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, errSort := c.GetInt("sort")

	if errCateId != nil || errAttrType != nil {
		c.Error("非法请求", "/goodsType")
		return
	}
	//去除空格 并判断是否为空
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}

	if errSort != nil {
		c.Error("排序值不对", "/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(models.GetUnix()),
	}
	models.DB.Create(&goodsTypeAttr)
	c.Success("增加成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))

}

func (c *GoodsTypeAttrController) Edit() {
	// c.Ctx.WriteString("修改")
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("非法请求", "/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttr)
	c.Data["goodsTypeAttr"] = goodsTypeAttr

	//获取所有类型
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType

	c.TplName = "admin/goodsTypeAttribute/edit.html"
}

func (c *GoodsTypeAttrController) DoEdit() {

	id, errId := c.GetInt("id")
	title := c.GetString("title")
	cateId, errCateId := c.GetInt("cate_id")
	attrType, errAttrType := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, errSort := c.GetInt("sort")

	if errId != nil || errCateId != nil || errAttrType != nil {
		c.Error("非法请求", "/goodsTypeAttribute")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	if errSort != nil {
		c.Error("排序值不对", "/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort

	err := models.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		c.Error("修改数据失败", "/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("需改数据成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))

}

func (c *GoodsTypeAttrController) Delete() {

	id, errId := c.GetInt("id")
	cateId, _ := c.GetInt("cate_id")
	if errId != nil {
		c.Error("传入参数错误", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))
		return
	}
	//if errCateId != nil {
	//	c.Error("传入参数错误", "/goodsType")
	//	return
	//}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Delete(&goodsTypeAttr)
	c.Success("删除数据成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
