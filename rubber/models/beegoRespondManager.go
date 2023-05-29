/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node: 需求响应 订单信息
**/
package models

type RespondManager struct {
	Id            int     `json:"id"`
	TUid          int     `json:"t_uid"`  //总用户号
	CUid          int     `json:"c_uid"`  //对应user表的id
	CName         string  `json:"c_name"` //公司名称
	CPhone        string  `json:"c_phone"`
	UserType      int     `json:"userType"`    //用户类型：1,普通新用户 2,普通老用户，3负荷集成商
	CAddress      string  `json:"c_address"`   //公司地址
	ContractNo    string  `json:"contract_no"` //签约编号
	FirstTrial    int     `json:"first_trial"` //0 默认 1通过 2拒绝
	Review        int     `json:"review"`      //0 默认 1通过 2拒绝
	FirLiaison    string  `json:"firLiaison"`
	CapId         int     `json:"cap_id"`
	FirLiaisonWay string  `json:"firLiaisonWay"`
	LinkAddress   string  `json:"linkAddress"`
	ContractCap   float64 `json:"contractCap"`  //合同容量
	RunningCap    float64 `json:"runningCap"`   //运行容量
	MonIns        int     `json:"monAndIns"`    //是否接入
	BiddingPrice  float64 `json:"biddingPrice"` //竞价价格
	BeResponse    float64 `json:"be_response"`
	OrderId       int     `json:"order_id"`
	Status        int     `json:"status"`
	OrderStatus   int     `json:"order_status"`  //订单状态：用户申请（进行：0）（完成：1）","市级初审(（完成：2）)","省级复审(（完成：3）)","签订协议(（完成：4）)","进入用户首页( (完成：5）) || 1：已申请；2：初审通过；3：复审通过；4：签订完成；5：流程完成。
	SignContract  int     `json:"sign_contract"` //签订协议 //0 默认 1通过
	AddTime int `json:"addTime"`
}

func (RespondManager) TableName() string {
	return "respond_manager"
}
