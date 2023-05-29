/**
  @Author ZXQ-Administrator
  @Date 2021-12-05
  @node:响应管理  orderStatus 0:未开始； 1：已申请；2：初审通过；3：复审通过；4：签订完成；5：流程完成。
**/
package elePowDR

import (
	"elenewenergy/blockchain"
	"elenewenergy/controllers/admin"
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"math"
	"strconv"
)

type ResponseManageController struct {
	admin.BaseController
}

type verifySelect struct {
	Id    int
	Value string
}
var (
	channelId   = beego.AppConfig.String("channel_id_un")
	chainCodeId = beego.AppConfig.String("chaincode_id_auth")
	chainCodeId1 = beego.AppConfig.String("chaincode_id_auth1")
	chainCodeId2 = beego.AppConfig.String("chaincode_id_auth2")
	userId      = beego.AppConfig.String("user_id")
	org_process    = beego.AppConfig.String("org_process")
	org_company    = beego.AppConfig.String("org_company")
	org_hospital    = beego.AppConfig.String("org_hospital")
)
func (c *ResponseManageController) CreateOrderItem() {



	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/responseManagement")
		beego.Info("Initialize.......", err)
		return
	}
	id := "AF-028"
	name := "西部橡胶林区26"
	date := "2019-12-12"
	quality := "优"
	healthy := "健康"
	soil := "湿润"
	abscissa := "100"
	ordinate := "60"
	open := "可割胶"
	temperature := "18"
	humidity := "0.2"
	phenology := "抽芽期"
	speed := "3"
	method := "深割"
	yellow := "0.3"
	way := " "
	//function string, args [][]byte
	//'{"Args":["query","a"]}'
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(id))
	bodyBytes = append(bodyBytes, []byte(name))
	bodyBytes = append(bodyBytes, []byte(date))
	bodyBytes = append(bodyBytes, []byte(quality))
	bodyBytes = append(bodyBytes, []byte(healthy))
	bodyBytes = append(bodyBytes, []byte(soil))
	bodyBytes = append(bodyBytes, []byte(abscissa))
	bodyBytes = append(bodyBytes, []byte(ordinate))
	bodyBytes = append(bodyBytes, []byte(open))
	bodyBytes = append(bodyBytes, []byte(temperature))
	bodyBytes = append(bodyBytes, []byte(humidity))
	bodyBytes = append(bodyBytes, []byte(phenology))
	bodyBytes = append(bodyBytes, []byte(speed))


	bodyBytes = append(bodyBytes, []byte(method))
	bodyBytes = append(bodyBytes, []byte(yellow))
	bodyBytes = append(bodyBytes, []byte(way))

	//调用智能合约

	data, errC := ccs.ChainCodeUpdate("setvalue", bodyBytes, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("errC", "/responseManagement")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	_ = data
	c.Success("数据上链成功", "/treeInformation")
}



func (c *ResponseManageController) QueryOrderItem() {

	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/responseManagement")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte("AF-001")}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/responseManagement")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *ResponseManageController) QueryOrderItem1() {

	ccs, err := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/responseManagement")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte("BF-001")}, "peer0.orgcompany.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/responseManagement")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *ResponseManageController) QueryOrderItem2() {

	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/responseManagement")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte("CF-001")}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/responseManagement")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *ResponseManageController) Get() {
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
	//查询数据
	respondManager := []models.RespondManager{}
	//select * from tableName limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&respondManager)

	//查询goods表里面的数量
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("respond_manager").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(respondManager) == 0 && count != 0 {
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/responseManagement?page=" + strconv.Itoa(prvPageBack))
	} else {
		//	订单为空
	}

	c.Data["respondManager"] = respondManager

	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword

	//预加载 将Pid=0的数据挂载到GoodsCateItem
	c.TplName = "admin/elePowDR/respManagement/index.html"

}

func (c *ResponseManageController) Verify() {
	//1、获取商品数据
	orderId, errId := c.GetInt("orderId")
	if errId != nil {
		c.Error("非法请求", "/responseManagement")
		return
	}
	// 审核 选择
	verifySelect := []*verifySelect{
		&verifySelect{0, "请选择"},
		&verifySelect{1, "审核通过"},
		&verifySelect{2, "审核拒绝"},
	}
	orderDemand := models.OrderDemand{}
	errInfo := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error

	orderBasic := models.OrderDemandBasic{}
	errInfo1 := models.DB.Where("order_id=?", orderId).Find(&orderBasic).Error

	orderCap := models.OrderDemandCapability{}
	errInfo2 := models.DB.Where("order_id=?", orderId).Find(&orderCap).Error

	orderDevice := models.OrderDemandDevice{}
	errInfo3 := models.DB.Where("order_id=?", orderId).Find(&orderDevice).Error

	respM := models.RespondManager{}
	errInfo4 := models.DB.Where("order_id=?", orderId).Find(&respM).Error

	dsmUserDevice := []models.DsmUserDevice{}
	errInfo5 := models.DB.Where("t_uid=? AND status=?", orderDemand.Uid, 0).Find(&dsmUserDevice).Error

	if errInfo != nil || errInfo1 != nil || errInfo2 != nil || errInfo3 != nil || errInfo4 != nil || errInfo5 != nil {
		c.Error("查询订单数据失败", "/responseManagement")
		return
	}
	c.Data["verifySelect"] = verifySelect

	c.Data["orderDemand"] = orderDemand
	c.Data["orderBasic"] = orderBasic
	c.Data["orderCap"] = orderCap
	c.Data["orderDevice"] = orderDevice
	c.Data["respondManager"] = respM
	c.Data["dsmUserDevice"] = dsmUserDevice

	//上一页地址
	c.Data["prevPageBack"] = c.Ctx.Request.Referer()
	c.TplName = "admin/elePowDR/respManagement/verify.html"
}

func (c *ResponseManageController) DoVerify() {
	orderId, errId := c.GetInt("order_id")
	if errId != nil {
		c.Error("非法请求", "/responseManagement")
		return
	}
	//上一页地址
	firstTrial, _ := c.GetInt("first_trial_id")
	reviewTrial, _ := c.GetInt("review_trial_id")
	if firstTrial == 2 && reviewTrial == 1 {
		reviewTrial = 0
	}
	respM := models.RespondManager{}
	errRespM := models.DB.Where("order_id=?", orderId).Find(&respM).Error
	if errRespM != nil {
		c.Error("查找审核检验数据失败，请联系管理员", "/responseManagement")
		return
	}
	respM.FirstTrial = firstTrial
	respM.Review = reviewTrial

	// 进程判断
	if reviewTrial == 1 && firstTrial == 1 {
		respM.OrderStatus = 3
	}
	if firstTrial == 1 && reviewTrial != 1 {
		respM.OrderStatus = 2
	}
	if firstTrial != 1 {
		respM.OrderStatus = 1
		respM.Review = 0
	}
	errInfo := models.DB.Save(&respM).Error
	if errInfo != nil {
		c.Error("保存审核检验失败，请联系管理员", "/responseManagement")
		return
	}

	orderDemand := models.OrderDemand{}
	errDemand := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error
	if errDemand != nil {
		c.Error("查找审核检验数据失败，请联系管理员", "/responseManagement")
		return
	}
	orderDemand.FirstTrial = firstTrial
	orderDemand.Review = reviewTrial
	// 进程判断 时间添加
	// 初审 请选择
	if firstTrial == 0 {
		orderDemand.FirTime = 0
		orderDemand.ThiTime = 0
		orderDemand.Review = 0
		orderDemand.OrderStatus = 1
	}
	// 初审 拒绝
	if firstTrial == 2 {
		orderDemand.FirTime = int(models.GetUnix())
		orderDemand.ThiTime = 0
		orderDemand.Review = 0
		orderDemand.OrderStatus = 1
	}
	// 初审通过 复审通过
	if reviewTrial == 1 && firstTrial == 1 {
		if orderDemand.SecTime == 0 {
			orderDemand.SecTime = int(models.GetUnix())
			orderDemand.ThiTime = int(models.GetUnix())
		} else {
			orderDemand.ThiTime = int(models.GetUnix())
		}
		orderDemand.OrderStatus = 3
	}
	// 初审通过 复审请选择
	if firstTrial == 1 && reviewTrial == 0 {
		if orderDemand.SecTime == 0 {
			orderDemand.SecTime = int(models.GetUnix())
			orderDemand.ThiTime = 0
		} else {
			orderDemand.ThiTime = 0
		}
		orderDemand.OrderStatus = 2
	}
	// 初审通过 复审拒绝
	if firstTrial == 1 && reviewTrial == 2 {
		if orderDemand.SecTime == 0 {
			orderDemand.SecTime = int(models.GetUnix())
			orderDemand.ThiTime = int(models.GetUnix())
		} else {
			orderDemand.ThiTime = int(models.GetUnix())
		}
		orderDemand.OrderStatus = 2
	}
	errInfo2 := models.DB.Save(&orderDemand).Error

	if errInfo2 != nil {
		c.Error("保存审核检验失败，请联系管理员", "/responseManagement")
		return
	}
	c.Success("修改数据成功", "/responseManagement")
}
func (c *ResponseManageController) DoConfirm() {
	orderId, errId := c.GetInt("order_id")

	orderDemand := models.OrderDemand{}
	errInfo := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error
	if errId != nil || errInfo != nil {
		c.Error("查询OrderDemand失败", "/responseManagement")
		return
	}
	respM := models.RespondManager{}
	errRespM := models.DB.Where("order_id=?", orderId).Find(&respM).Error
	if errRespM != nil {
		c.Error("查询RespondManager失败，请联系管理员", "/responseManagement")
		return
	}
	respM.OrderStatus = 5
	errRespM2 := models.DB.Save(&respM).Error
	if errRespM2 != nil {
		c.Error("RespondManager保存失败，请联系管理员", "/responseManagement")
		return
	}
	timeFinish := int(models.GetUnix())
	// orderDemand表
	orderDemand.OrderStatus = 5
	orderDemand.FinishTime = timeFinish
	orderDemand.ProcessTime = timeFinish-orderDemand.FirTime
	errDemand2 := models.DB.Save(&orderDemand).Error
	if errDemand2 != nil {
		c.Error("OrderDemand保存失败，请联系管理员", "/responseManagement")
		return
	}
	c.Success("确认流程成功", "/responseManagement")
}

func (c *ResponseManageController) Delete() {
	orderId, errId := c.GetInt("orderId")
	if errId != nil {
		c.Error("传入参数错误", "/responseManagement")
		return
	}
	orderItem := models.OrderDemand{}
	models.DB.Where("order_Id=?", orderId).Find(&orderItem)

	if orderItem.FirstTrial == 1 {
		c.Error("市级初审通过后不可删除订单，请联系管理员", "/responseManagement")
		return
	}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderItem)

	orderBasic := models.OrderDemandBasic{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderBasic)

	orderCap := models.OrderDemandCapability{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderCap)

	orderDevice := models.OrderDemandDevice{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderDevice)

	respM := models.RespondManager{}
	models.DB.Where("order_Id = ?", orderId).Delete(&respM)

	c.Success("删除响应订单成功", "/responseManagement")
}
