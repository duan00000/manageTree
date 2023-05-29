/**
  @Author ZXQ-Administrator
  @Date 2021-12-05
  @node:签订协议
**/
package elePowDR

import (
	"elenewenergy/controllers/admin"
	"elenewenergy/models"
	"math"
	"strconv"
)

type SignContractController struct {
	admin.BaseController
}
type signContract struct {
	SignContract  int   `json:"sign_contract"` //签订协议 //0 默认 1通过
}

func (c *SignContractController) Get() {
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
		where += " AND c_name like \"%" + keyword + "%\""
	}
	where += " AND review = 1"
	//查询数据
	orderDemandList := []models.OrderDemand{}
	//select * from tableName limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&orderDemandList)

	//查询goods表里面的数量
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("order_demand").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(orderDemandList) == 0 && count != 0 {
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/signContract?page=" + strconv.Itoa(prvPageBack))
	} else {
		//	订单为空
	}

	c.Data["orderDemandList"] = orderDemandList

	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword

	//预加载 将Pid=0的数据挂载到GoodsCateItem
	c.TplName = "admin/elePowDR/signContract/index.html"
}

func (c *SignContractController) Sign() {
	//1、获取商品数据
	orderId, errId := c.GetInt("orderId")
	if errId != nil {
		c.Error("非法请求", "/signContract")
		return
	}
	partUser := models.OrderPart{}
	models.DB.Where("id=?", 1).Find(&partUser)
	signContractExit := signContract{}
	signContractExit.SignContract = 0
	orderDemand := models.OrderDemand{}
	errInfo := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error

	orderBasic := models.OrderDemandBasic{}
	errInfo1 := models.DB.Where("order_id=?", orderId).Find(&orderBasic).Error

	orderCap := models.OrderDemandCapability{}
	errInfo2 := models.DB.Where("order_id=?", orderId).Find(&orderCap).Error

	orderDevice := models.OrderDemandDevice{}
	errInfo3 := models.DB.Where("order_id=?", orderId).Find(&orderDevice).Error

	dsmUserDevice := []models.DsmUserDevice{}
	errInfo5 := models.DB.Where("t_uid=? AND status=?", orderDemand.Uid, 0).Find(&dsmUserDevice).Error

	orderContract := []models.OrderContract{}
	errInfo4 := models.DB.Where("order_id=?", orderId).Find(&orderContract).Error

	if len(orderContract) < 1{
		signContractExit.SignContract = 0
		c.Data["signContractExit"] = signContractExit

	}else {
		c.Data["signContractExit"] = orderContract[0]

	}
	if errInfo != nil || errInfo1 != nil || errInfo2 != nil || errInfo3 != nil || errInfo4 != nil|| errInfo5 != nil {
		c.Error("查询订单数据失败", "/signContract")
		return
	}

	c.Data["orderDemand"] = orderDemand
	c.Data["orderBasic"] = orderBasic
	c.Data["orderCap"] = orderCap
	c.Data["orderDevice"] = orderDevice
	c.Data["partUser"] = partUser
	c.Data["dsmUserDevice"] = dsmUserDevice

	//上一页地址
	c.Data["prevPageBack"] = c.Ctx.Request.Referer()
	c.TplName = "admin/elePowDR/signContract/sign.html"
}

func (c *SignContractController) DoSign() {
	orderId, errId := c.GetInt("order_id")

	orderDemand := models.OrderDemand{}
	errInfo := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error
	if errId != nil || errInfo != nil {
		c.Error("查询OrderDemand失败", "/signContract")
		return
	}
	lawValue := c.GetString("law_id")
	addressValue := c.GetString("address_id")
	emailValue := c.GetString("email_id")

	linkValue := c.GetString("link_id")
	phoneValue := c.GetString("phone_id")
	faxValue := c.GetString("fax_id")
	timeContract := int(models.GetUnix())

	signContract := models.OrderContract{
		CPhone:       phoneValue,
		CName:        orderDemand.CName,
		CAddress:     addressValue,
		CLaw:         lawValue,
		Fax:          faxValue,
		Email:        emailValue,
		LinkUser:     linkValue,
		OrderId:      orderDemand.OrderId,
		SignContract: 1,
		AddTime:      timeContract,
	}
	errContract := models.DB.Create(&signContract).Error
	if errContract != nil {
		c.Error("Contract创建失败", "/signContract")
		return
	}

	respM := models.RespondManager{}
	errRespM := models.DB.Where("order_id=?", orderId).Find(&respM).Error
	if errRespM != nil {
		c.Error("查询RespondManager失败，请联系管理员", "/signContract")
		return
	}
	respM.OrderStatus = 4
	respM.SignContract = 1
	errRespM2 := models.DB.Save(&respM).Error
	if errRespM2 != nil {
		c.Error("RespondManager合同保存失败，请联系管理员", "/signContract")
		return
	}

	orderDemand.OrderStatus = 4
	orderDemand.SignContract = 1
	orderDemand.FouTime = timeContract
	errDemand := models.DB.Save(&orderDemand).Error
	if errDemand != nil {
		c.Error("SignContract合同保存失败，请联系管理员", "/signContract")
		return
	}
	c.Success("修改数据成功", "/signContract")
}

