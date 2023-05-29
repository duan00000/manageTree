/**
  @Author ZXQ-Administrator
  @Date 2021-09-08
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"strconv"
	"strings"
	"sync"
)

//实例化该结构体
var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

func (c *GoodsController) Get() {
	//当前分页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每页显示数量
	pageSize := 10

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	goodsList := []models.Goods{}
	//select * from tableName limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	//查询goods表里面的数量
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("goods").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(goodsList) == 0  && count != 0{
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPageBack))
	}else{
		//	商品为空
	}
	c.Data["goodsList"] = goodsList
	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword
	c.TplName = "admin/goods/index.html"
}

func (c *GoodsController) Add() {
	//获取商品分类
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=?", 0).Preload("GoodsCateItem").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	beego.Info(goodsCate)
	//获取商品颜色
	goodsColor := []models.GoodsColor{}
	models.DB.Find(&goodsColor)
	c.Data["goodsColor"] = goodsColor
	//获取商品类型信息
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType

	c.TplName = "admin/goods/add.html"
}

func (c *GoodsController) DoAdd() {
	//1. 获取表单提交过来的数据
	title := c.GetString("title")
	subTitle := c.GetString("sub_title")
	goodsSn := c.GetString("goods_sn")
	cateId, _ := c.GetInt("cate_id")
	goodsNumber, _ := c.GetInt("goods_number")
	marketPrice, _ := c.GetFloat("market_price")
	price, _ := c.GetFloat("price")
	relationGoods := c.GetString("relation_goods")
	goodsAttr := c.GetString("goods_attr")
	goodsVersion := c.GetString("goods_version")
	goodsGift := c.GetString("goods_gift")
	goodsFitting := c.GetString("goods_fitting")
	goodsColor := c.GetStrings("goods_color") //获取多个
	goodsKeywords := c.GetString("goods_keywords")
	goodsDesc := c.GetString("goods_desc")
	goodsContent := c.GetString("goods_content")
	isDelete, _ := c.GetInt("is_delete")
	isHot, _ := c.GetInt("is_hot")
	isBest, _ := c.GetInt("is_best")
	isNew, _ := c.GetInt("is_new")
	goodsTypeId, _ := c.GetInt("goods_type_id")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	addTime := int(models.GetUnix())
	if title == "" {
		c.Error("标题不能为空", "/goods/add")
		return
	}

	//2. 获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColor, ",")

	//3.上传图片 生成缩略图
	goodsImg, errGoodsImg := c.UploadImg("goods_img", title)
	if errGoodsImg == nil && len(goodsImg) > 0 {
		//生成缩略图
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			wg.Add(1)
			go func() {
				models.ResizeImage(goodsImg)
				wg.Done()
			}()
		}
	}
	//4. 增加商品数据
	goods := models.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	errCreate := models.DB.Create(&goods).Error
	if errCreate != nil {
		c.Error("增加失败", "/goods/add")
	}
	//5. 增加图库 信息
	//执行标记加1
	wg.Add(1)
	//协程 匿名函数
	go func() {
		goodsImageList := c.GetStrings("goods_image_list")
		for _, value := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			//将上面获取的goods信息 添加到图库
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = value
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		//执行标记减1
		wg.Done()
	}()

	//6. 增加商品属性
	wg.Add(1)
	go func() {
		attrIdList := c.GetStrings("attr_id_list")
		attrValueList := c.GetStrings("attr_value_list")

		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			models.DB.Find(&goodsTypeAttributeObj)

			goodsAttrObj := models.GoodsAttr{}
			goodsAttrObj.GoodsId = goods.Id
			goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
			goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
			goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
			goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
			goodsAttrObj.AttributeValue = attrValueList[i]
			goodsAttrObj.Status = 1
			goodsAttrObj.Sort = 10
			goodsAttrObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()
	//等待协程执行完毕
	wg.Wait()
	c.Success("增加数据成功", "/goods")

}
func (c *GoodsController) Edit() {
	//1、获取商品数据
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("非法请求", "/goods")
	}
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	c.Data["goods"] = goods

	//2、获取商品分类
	//获取商品分类
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=?", 0).Preload("GoodsCateItem").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate

	//3、获取所有颜色 以及选中的颜色
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}
	//获取商品颜色
	goodsColor := []models.GoodsColor{}
	models.DB.Find(&goodsColor)
	//已选择的颜色选中
	for i := 0; i < len(goodsColor); i++ {
		if _, ok := goodsColorMap[strconv.Itoa(goodsColor[i].Id)]; ok {
			goodsColor[i].Checked = true
		}
	}
	c.Data["goodsColor"] = goodsColor
	//4、商品的图库信息
	goodsImage := []models.GoodsImage{}
	models.DB.Where("goods_id=?", goods.Id).Find(&goodsImage)
	c.Data["goodsImage"] = goodsImage
	//5、获取商品类型
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	//6、获取规格信息
	goodsAttr := []models.GoodsAttr{}
	models.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)

	var goodsAttrStr string
	for _, v := range goodsAttr {
		if v.AttributeType == 1 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {

			// 获取 attr_value  获取可选值列表
			oneGoodsTypeAttribute := models.GoodsTypeAttribute{Id: v.AttributeId}
			models.DB.Find(&oneGoodsTypeAttribute)
			attrValueSlice := strings.Split(oneGoodsTypeAttribute.AttrValue, "\n")
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			for j := 0; j < len(attrValueSlice); j++ {
				if attrValueSlice[j] == v.AttributeValue {
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[j], attrValueSlice[j])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[j], attrValueSlice[j])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)
		}
	}
	c.Data["goodsAttrStr"] = goodsAttrStr
	//上一页地址
	c.Data["prevPageBack"] = c.Ctx.Request.Referer()
	c.TplName = "admin/goods/edit.html"
}

func (c *GoodsController) DoEdit() {
	//1、获取要修改的商品数据
	page, _ := c.GetInt("page")
	fmt.Printf("page:-----", page)
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("非法参数", "/goods")
	}
	title := c.GetString("title")
	subTitle := c.GetString("sub_title")
	goodsSn := c.GetString("goods_sn")
	cateId, _ := c.GetInt("cate_id")
	goodsNumber, _ := c.GetInt("goods_number")
	marketPrice, _ := c.GetFloat("market_price")
	price, _ := c.GetFloat("price")
	relationGoods := c.GetString("relation_goods")
	goodsAttr := c.GetString("goods_attr")
	goodsVersion := c.GetString("goods_version")
	goodsGift := c.GetString("goods_gift")
	goodsFitting := c.GetString("goods_fitting")
	goodsColor := c.GetStrings("goods_color")
	goodsKeywords := c.GetString("goods_keywords")
	goodsDesc := c.GetString("goods_desc")
	goodsContent := c.GetString("goods_content")
	isDelete, _ := c.GetInt("is_delete")
	isHot, _ := c.GetInt("is_hot")
	isBest, _ := c.GetInt("is_best")
	isNew, _ := c.GetInt("is_new")
	goodsTypeId, _ := c.GetInt("goods_type_id")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	//提交后获取修改页面上一页的地址
	prevPageBack := c.GetString("prevPageBack")
	if title == "" {
		c.Error("标题不能为空,返回上一页", prevPageBack)
		return
	}

	//2、获取颜色信息  把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColor, ",")
	//将获取的数据赋值
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//3、上传图片 生成缩略图
	goodsImg, errImg := c.UploadImg("goods_img", title)
	if errImg == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		//	处理图片 生成缩略图
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			wg.Add(1)
			go func() {
				models.ResizeImage(goodsImg)
				wg.Done()
			}()
		}

	}
	//4、执行修改商品
	errInfo := models.DB.Save(&goods).Error
	if errInfo != nil {
		c.Error("修改数据失败,返回上一页", prevPageBack)
		return
	}
	//5、修改图库数据（增加）
	wg.Add(1)
	go func() {
		goodsImageList := c.GetStrings("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()
	//6、修改商品类型属性数据 	1、删除当前商品的id对应的类型属性   2、执行增加
	//删除当前商品id对应的类型属性
	goodsAttrObj := models.GoodsAttr{}
	models.DB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	//执行增加
	wg.Add(1)
	go func() {
		attrIdList := c.GetStrings("attr_id_list")
		attrValueList := c.GetStrings("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			models.DB.Find(&goodsTypeAttributeObj)

			goodsAttrObj := models.GoodsAttr{}
			goodsAttrObj.GoodsId = goods.Id
			goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
			goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
			goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
			goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
			goodsAttrObj.AttributeValue = attrValueList[i]
			goodsAttrObj.Status = 1
			goodsAttrObj.Sort = 10
			goodsAttrObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()

	wg.Wait()
	c.Success("修改数据成功", prevPageBack)
}

func (c *GoodsController) Delete() {
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/goods")
		return
	}
	goods := models.Goods{Id: id}
	errInfo := models.DB.Delete(&goods).Error
	if errInfo != nil {
		//	删除属性
		goodsAttr := models.GoodsAttr{GoodsId: id}
		models.DB.Delete(&goodsAttr)
		//	删除图库信息
		goodsImage := models.GoodsImage{GoodsId: id}
		models.DB.Delete(&goodsImage)
	}
	// os.Remove() 删除硬盘图

	c.Success("删除商品数据成功", c.Ctx.Request.Referer())
}

// 给富文本编辑器上传图片
func (c *GoodsController) DoUpload() {
	savePath, err := c.UploadImg("file", "RichText_")
	//https://froala.com/wysiwyg-editor/docs/options?#imageUploadURL
	// 返回json数据 {}{link: 'path/to/image.jpg'}.
	if err != nil {
		c.Error("上传失败", "")
		c.Data["json"] = map[string]interface{}{
			"link": "",
		}
		c.ServeJSON()
	} else {
		//生成缩略图
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			models.ResizeImage(savePath)
		}
		//返回json数据 {link: 'path/to/image.jpg'}
		//oss存储地址
		if ossStatus{
			c.Data["json"] = map[string]interface{}{
				"link": beego.AppConfig.String("ossDomain") + "/" + savePath,
			}
		}else {
			c.Data["json"] = map[string]interface{}{
				"link": "/" + savePath,
			}
		}
		c.ServeJSON()
	}
}

//给图库上传图片
func (c *GoodsController) DoUploadPhoto() {
	savePath, err := c.UploadImg("file","Gallery_")
	if err != nil {
		beego.Error("失败")
		c.Data["json"] = map[string]interface{}{
			"link": "",
		}
		c.ServeJSON()
	} else {
		//生成缩略图
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			models.ResizeImage(savePath)
		}
		c.Data["json"] = map[string]interface{}{
			"link": savePath,
		}
		c.ServeJSON()
	}
}
//获取商品类型属性
func (c *GoodsController) GetGoodsTypeAttribute() {
	cateId, errCateId := c.GetInt("cate_id")
	GoodsTypeAttribute := []models.GoodsTypeAttribute{}
	errInfo := models.DB.Where("cate_id=?", cateId).Find(&GoodsTypeAttribute).Error
	if errCateId != nil || errInfo != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "",
			"success": false,
		}
		c.ServeJSON()

	} else {
		c.Data["json"] = map[string]interface{}{
			"result":  GoodsTypeAttribute,
			"success": true,
		}
		c.ServeJSON()
	}
}

//图片对应颜色信息
func (c *GoodsController) ChangeGoodsImageColor() {
	colorId, errColorId := c.GetInt("id")
	goodsImageId, errGoodsImageId := c.GetInt("goods_image_id")
	goodsImage := models.GoodsImage{Id: goodsImageId}
	models.DB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	errInfo := models.DB.Save(&goodsImage).Error
	if errColorId != nil || errGoodsImageId != nil || errInfo != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "更新失败",
			"success": false,
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{
			"result":  "更新成功",
			"success": true,
		}
		c.ServeJSON()
	}
}

//删除图库
func (c *GoodsController) RemoveGoodsImage() {
	goodsImageId, errGoodsImageId := c.GetInt("goods_image_id")
	goodsImage := models.GoodsImage{Id: goodsImageId}
	errInfo := models.DB.Delete(&goodsImage).Error
	//   /static/upload/20200622/1592799750042592100.jpg

	if errGoodsImageId != nil || errInfo != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "删除失败",
			"success": false,
		}
		c.ServeJSON()
	} else {
		//删除图片
		// os.Remove()
		c.Data["json"] = map[string]interface{}{
			"result":  "已删除",
			"success": true,
		}
		c.ServeJSON()
	}

}
