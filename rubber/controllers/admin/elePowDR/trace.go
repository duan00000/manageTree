package elePowDR

import (
	"elenewenergy/blockchain"
	"elenewenergy/controllers/admin"
	"elenewenergy/models"
	"github.com/astaxie/beego"
)


type TraceController struct {
	admin.BaseController
}
//查询西部林区数据
func (c *TraceController) QueryTree1() {
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

