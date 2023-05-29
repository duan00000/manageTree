/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node: Vue需求响应基础信息
**/
package models

type OrderDemandBasic struct {
	Id            int     `json:"id"`
	CompanyId     int     `json:"company_id"` //与vue_company 中的id绑定
	VueBasicId    int     `json:"vue_basic_id"`  //与vue_demand_basic 中的id绑定
	JoinY         int     `json:"joinY"`     //加入年份 当年
	FirLiaison    string  `json:"firLiaison"`
	FirLiaisonWay string  `json:"firLiaisonWay"`
	SecLiaison    string  `json:"secLiaison"`
	SecLiaisonWay string  `json:"secLiaisonWay"`
	ContractCap   float64 `json:"contractCap"` //合同容量
	RunningCap    float64 `json:"runningCap"`  //运行容量
	LinkAddress   string  `json:"linkAddress"`
	TRC           string  `json:"trc"`          //税务登记证
	MonIns        int     `json:"monAndIns"`    //是否接入
	BiddingPrice  float64 `json:"biddingPrice"` //竞价价格
	Status        int     `json:"status"`
	OrderId       int     `json:"order_id"`
	AddTime       int     `json:"add_time"`
}

func (OrderDemandBasic) TableName() string {
	return "order_demand_basic"
}
