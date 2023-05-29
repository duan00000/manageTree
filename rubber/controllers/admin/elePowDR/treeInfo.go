package elePowDR

import (
	"elenewenergy/blockchain"
	"elenewenergy/controllers/admin"
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"math"
	"strconv"
	"strings"
)


type TreeController struct {
	admin.BaseController
}
//查询西部林区数据
func (c *TreeController) QueryTree1() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte(treeId)}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询中部林区数据
func (c *TreeController) QueryTree2() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte(treeId)}, "peer0.orgcompany.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询东部林区数据
func (c *TreeController) QueryTree3() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("query", [][]byte{[]byte(treeId)}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)

}


//查询东部林区数据
func (c *TreeController) Trace() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("trace", [][]byte{[]byte(treeId)}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}


	//c.Ctx.ResponseWriter.WriteHeader(200)
	//c.Ctx.ResponseWriter.Write(data)
	s:=string(data)
	c.Data["treeId"] = treeId
	c.Data["data"] = s
	c.TplName = "admin/elePowDR/treeInformation/test.html"

	//c.Success(s,"/treeInformation")
}

func (c *TreeController) Get() {
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
		where += " AND tree_name like \"%" + keyword + "%\""
	}
	//查询数据
	treeInfoList := []models.TreeInfo{}
	//select * from tableName limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&treeInfoList)

	//查询goods表里面的数量
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("tree_info").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(treeInfoList) == 0 && count != 0 {
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/tree?page=" + strconv.Itoa(prvPageBack))
	} else {
		//	订单为空
	}

	c.Data["treeInfoList"] = treeInfoList

	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword

	//预加载 将Pid=0的数据挂载到GoodsCateItem
	c.TplName = "admin/elePowDR/treeInformation/index.html"

}
func (c *TreeController) Edit() {
	//获取管理员信息
	id,err:=c.GetInt("id")
	if err!=nil{
		c.Error("非法请求", "/treeInformation")
		return
	}

	//获取manager信息
	treeinfo:=models.TreeInfo{Id:id}

	models.DB.Find(&treeinfo)
	c.Data["treeinfo"]=treeinfo

	c.TplName = "admin/elePowDR/treeInformation/awake.html"
}


func (c *TreeController) OnChain1() {

	id,err:=c.GetInt("id")
	if err!=nil{
		c.Error("非法请求", "/treeInformation")
		return
	}


	treeInfo:=models.TreeInfo{Id:id}
	models.DB.Find(&treeInfo)
	no := c.GetString("treeno")
	name := c.GetString("treename")
	date := c.GetString("date")
	quality := c.GetString("quality")
	healthy := c.GetString("status")
	soil := c.GetString("grade")
	abscissa := c.GetString("heng")
	ordinate := c.GetString("zong")
	open := c.GetString("open")
	temperature := c.GetString("tem")
	humidity := c.GetString("hum")
	phenology := c.GetString("phase")
	speed := c.GetString("speed")
	method := c.GetString("method")
	yellow := c.GetString("yellow")
	way := c.GetString("way")
	fromid := c.GetString("fromid")
    _ = fromid
    _ = way
	s1 := name
	s2 := string([]rune(no)[5:6])
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	cname := build.String()
	judge := string([]rune(no)[0:1])
	//panduan
	if judge == "A"{
		ccs, err7 := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err)
			return
		}
		/*treeInfo.TreeId=no
		error9:=models.DB.Save(&treeInfo).Error
		if error9!=nil{
			c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
			return
		}*/



		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(no))
		bodyBytes = append(bodyBytes, []byte(cname))
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
			c.Error(way, "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}
	//panduan2
	if judge =="B"{
		ccs, err7 := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err)
			return
		}
		/*treeInfo.TreeId=no
		error9:=models.DB.Save(&treeInfo).Error
		if error9!=nil{
			c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
			return
		}*/



		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(no))
		bodyBytes = append(bodyBytes, []byte(cname))
		bodyBytes = append(bodyBytes, []byte(date))
		bodyBytes = append(bodyBytes, []byte(quality))
		bodyBytes = append(bodyBytes, []byte(healthy))
		bodyBytes = append(bodyBytes, []byte(soil))
		bodyBytes = append(bodyBytes, []byte(abscissa))
		bodyBytes = append(bodyBytes, []byte(ordinate))
		bodyBytes = append(bodyBytes, []byte(open))
		bodyBytes = append(bodyBytes, []byte(fromid))
		bodyBytes = append(bodyBytes, []byte(temperature))
		bodyBytes = append(bodyBytes, []byte(humidity))
		bodyBytes = append(bodyBytes, []byte(phenology))
		bodyBytes = append(bodyBytes, []byte(speed))
		bodyBytes = append(bodyBytes, []byte(method))
		bodyBytes = append(bodyBytes, []byte(yellow))


		//调用智能合约

		data, errC := ccs.ChainCodeUpdate("setvalue", bodyBytes, "peer0.orgcompany.trace.com")
		if errC != nil {
			c.Error(no, "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}
	//pandaun3
	if judge =="C"{
		ccs, err7 := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err)
			return
		}
		/*treeInfo.TreeId=no
		error9:=models.DB.Save(&treeInfo).Error
		if error9!=nil{
			c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
			return
		}*/



		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(no))
		bodyBytes = append(bodyBytes, []byte(cname))
		bodyBytes = append(bodyBytes, []byte(date))
		bodyBytes = append(bodyBytes, []byte(quality))
		bodyBytes = append(bodyBytes, []byte(healthy))
		bodyBytes = append(bodyBytes, []byte(soil))
		bodyBytes = append(bodyBytes, []byte(abscissa))
		bodyBytes = append(bodyBytes, []byte(ordinate))
		bodyBytes = append(bodyBytes, []byte(open))
		bodyBytes = append(bodyBytes, []byte(fromid))
		bodyBytes = append(bodyBytes, []byte(temperature))
		bodyBytes = append(bodyBytes, []byte(humidity))
		bodyBytes = append(bodyBytes, []byte(phenology))
		bodyBytes = append(bodyBytes, []byte(speed))
		bodyBytes = append(bodyBytes, []byte(method))
		bodyBytes = append(bodyBytes, []byte(yellow))


		//调用智能合约

		data, errC := ccs.ChainCodeUpdate("setvalue", bodyBytes, "peer0.orghospital.trace.com")
		if errC != nil {
			c.Error(cname, "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}



}
func (c *TreeController) Plan() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("plan", [][]byte{[]byte(treeId)}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *TreeController) OnJudge1() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("judge", [][]byte{[]byte(treeId)}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
func (c *TreeController) OnJudge2() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("judge", [][]byte{[]byte(treeId)}, "peer0.orgcompany.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询东部林区数据
func (c *TreeController) OnJudge3() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("judge", [][]byte{[]byte(treeId)}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
func (c *TreeController) IsOpen1() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("isopen", [][]byte{[]byte(treeId)}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
func (c *TreeController) IsOpen2() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("isopen", [][]byte{[]byte(treeId)}, "peer0.orgcompany.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询东部林区数据
func (c *TreeController) IsOpen3() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeUpdate("isopen", [][]byte{[]byte(treeId)}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeUpdate.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}

//查询西部林区数据
func (c *TreeController) History1() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("gethistory", [][]byte{[]byte(treeId)}, "peer0.orgprocess.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询中部林区数据
func (c *TreeController) History2() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("gethistory", [][]byte{[]byte(treeId)}, "peer0.orgcompany.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
//查询东部林区数据
func (c *TreeController) History3() {
	treeId:= c.GetString("treeId")

	treeInfo := models.TreeInfo{}
	errInfo := models.DB.Where("tree_id=?",treeId).Find(&treeInfo).Error
	if errInfo != nil {
		c.Error("查询订单数据失败", "/treeInformation")
		return
	}
	ccs, err := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
	if err != nil {
		c.Error("err", "/treeInformation")
		beego.Info("Initialize.......", err)
		return
	}


	data, errC := ccs.ChainCodeQuery("gethistory", [][]byte{[]byte(treeId)}, "peer0.orghospital.trace.com")
	if errC != nil {
		c.Error("查询链上数据失败", "/treeInformation")
		beego.Info("ChainCodeQuery.......", errC)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Write(data)
}
func (c *TreeController) Add() {
	treeInfo := []models.TreeInfo{}
	models.DB.Find(&treeInfo)
	c.Data["treeInfo"] = treeInfo
	c.TplName = "admin/elePowDR/treeInformation/add.html"
}
func (c *TreeController) DoAdd() {


	no := c.GetString("treeno")
	name := c.GetString("treename")
	date := c.GetString("date")
	quality := c.GetString("quality")
	healthy := c.GetString("status")
	soil := c.GetString("grade")
	abscissa := c.GetString("heng")
	ordinate := c.GetString("zong")
	open := c.GetString("open")
	temperature := c.GetString("tem")
	humidity := c.GetString("hum")
	phenology := c.GetString("phase")
	speed := c.GetString("speed")
	method := c.GetString("method")
	yellow := c.GetString("yellow")
	way := c.GetString("way")
	fromid := c.GetString("fromid")
	_ = fromid
	_ = way
	ano :="AF-00"
	bno:="BF-00"
	cno:="CF-00"



	finalname := string([]rune(name)[0:6])
	judge := string([]rune(name)[0:1])
	//panduan
	if judge == "西"{
		s2 := string([]rune(name)[6:])
		var build strings.Builder
		build.WriteString(ano)
		build.WriteString(s2)
		finalno := build.String()
		treeInfo := models.TreeInfo{
			TreeId:     finalno,
			TreeName: finalname,

		}
		models.DB.Create(&treeInfo)
		c.Success("增加成功", "/treeInformation")

		ccs, err7 := blockchain.Initialize(channelId, chainCodeId, org_process, userId, "CORE_PROCESS_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err7)
			return
		}




		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(finalno))
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
			c.Error(way, "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}
	//panduan2
	if judge =="中"{
		s2 := string([]rune(name)[6:])
		var build strings.Builder
		build.WriteString(bno)
		build.WriteString(s2)
		finalno := build.String()
		treeInfo := models.TreeInfo{
			TreeId:     finalno,
			TreeName: finalname,

		}
		models.DB.Create(&treeInfo)
		c.Success("增加成功", "/treeInformation")
		ccs, err7 := blockchain.Initialize(channelId, chainCodeId1, org_company, userId, "CORE_COMPANY_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err7)
			return
		}
		/*treeInfo.TreeId=no
		error9:=models.DB.Save(&treeInfo).Error
		if error9!=nil{
			c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
			return
		}*/



		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(finalno))
		bodyBytes = append(bodyBytes, []byte(name))
		bodyBytes = append(bodyBytes, []byte(date))
		bodyBytes = append(bodyBytes, []byte(quality))
		bodyBytes = append(bodyBytes, []byte(healthy))
		bodyBytes = append(bodyBytes, []byte(soil))
		bodyBytes = append(bodyBytes, []byte(abscissa))
		bodyBytes = append(bodyBytes, []byte(ordinate))
		bodyBytes = append(bodyBytes, []byte(open))
		bodyBytes = append(bodyBytes, []byte(fromid))
		bodyBytes = append(bodyBytes, []byte(temperature))
		bodyBytes = append(bodyBytes, []byte(humidity))
		bodyBytes = append(bodyBytes, []byte(phenology))
		bodyBytes = append(bodyBytes, []byte(speed))
		bodyBytes = append(bodyBytes, []byte(method))
		bodyBytes = append(bodyBytes, []byte(yellow))


		//调用智能合约

		data, errC := ccs.ChainCodeUpdate("setvalue", bodyBytes, "peer0.orgcompany.trace.com")
		if errC != nil {
			c.Error(no, "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}
	//pandaun3
	if judge =="东"{
		s2 := string([]rune(name)[6:])
		var build strings.Builder
		build.WriteString(cno)
		build.WriteString(s2)
		finalno := build.String()
		treeInfo := models.TreeInfo{
			TreeId:     finalno,
			TreeName: finalname,

		}
		models.DB.Create(&treeInfo)
		c.Success("增加成功", "/treeInformation")
		ccs, err7 := blockchain.Initialize(channelId, chainCodeId2, org_hospital, userId, "CORE_HOSPITAL_CONFIG_FILE")
		if err7 != nil {
			c.Error("err", "/treeInformation")
			beego.Info("Initialize.......", err7)
			return
		}
		/*treeInfo.TreeId=no
		error9:=models.DB.Save(&treeInfo).Error
		if error9!=nil{
			c.Error("请检查您修改的数据格式", "/focus/edit?id="+strconv.Itoa(id))
			return
		}*/



		//function string, args [][]byte
		//'{"Args":["query","a"]}'
		var bodyBytes [][]byte
		bodyBytes = append(bodyBytes, []byte(finalno))
		bodyBytes = append(bodyBytes, []byte(name))
		bodyBytes = append(bodyBytes, []byte(date))
		bodyBytes = append(bodyBytes, []byte(quality))
		bodyBytes = append(bodyBytes, []byte(healthy))
		bodyBytes = append(bodyBytes, []byte(soil))
		bodyBytes = append(bodyBytes, []byte(abscissa))
		bodyBytes = append(bodyBytes, []byte(ordinate))
		bodyBytes = append(bodyBytes, []byte(open))
		bodyBytes = append(bodyBytes, []byte(fromid))
		bodyBytes = append(bodyBytes, []byte(temperature))
		bodyBytes = append(bodyBytes, []byte(humidity))
		bodyBytes = append(bodyBytes, []byte(phenology))
		bodyBytes = append(bodyBytes, []byte(speed))
		bodyBytes = append(bodyBytes, []byte(method))
		bodyBytes = append(bodyBytes, []byte(yellow))


		//调用智能合约

		data, errC := ccs.ChainCodeUpdate("setvalue", bodyBytes, "peer0.orghospital.trace.com")
		if errC != nil {
			c.Error("errC", "/treeInformation")
			beego.Info("ChainCodeQuery.......", errC)
			return
		}
		_ = data
		c.Success("数据上链成功", "/treeInformation")
	}



}
