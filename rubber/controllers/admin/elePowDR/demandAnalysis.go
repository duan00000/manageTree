/**
  @Author ZXQ-Administrator
  @Date 2021-12-05
  @node:需求分析
**/
package elePowDR

import (
	"elenewenergy/controllers/admin"
)

type DemandAnalyController struct {
		admin.BaseController
}

func (c *DemandAnalyController) Get() {
	//预加载 将Pid=0的数据挂载到GoodsCateItem
	c.TplName = "admin/elePowDR/demandAnalysis.html"

}