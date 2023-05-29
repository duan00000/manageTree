/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node: 需求响应 订单信息
**/
package models

type OrderDemand struct {
	Id          int    `json:"id"`
	CompanyId   int    `json:"company_id"`
	Uid         int    `json:"uid"`      //总用户id
	UserType    int    `json:"userType"` //用户类型：1,普通新用户 2,普通老用户，3负荷集成商
	CName       string  `json:"c_name"`
	Phone       string `json:"phone"`    //电话
	Address     string `json:"address"`
	FirstTrial  int    `json:"first_trial"` //0 默认 1通过 2拒绝
	Review      int    `json:"review"`      //0 默认 1通过 2拒绝
	BasicId     int    `json:"basic_id"`
	DeviceId    int    `json:"device_id"`
	CapId       int    `json:"cap_id"`
	FirTime     int    `json:"fir_time"`
	SecTime     int    `json:"sec_time"`
	ThiTime     int    `json:"thi_time"`
	FouTime     int    `json:"fou_time"`
	OrderId     int    `json:"order_id"`
	ProcessTime int    `json:"process_Time"`
	FinishTime  int    `json:"finish_Time"`
	SignContract int   `json:"sign_contract"` //签订协议 //0 默认 1通过
	Status      int    `json:"status"`
	OrderStatus int    `json:"orderStatus"` //订单状态：用户申请（进行：0）（完成：1）","市级初审(（完成：2）)","省级复审(（完成：3）)","签订协议(（完成：4）)","进入用户首页( (完成：5）) || 1：已申请；2：初审通过；3：复审通过；4：签订完成；5：流程完成。
}

func (OrderDemand) TableName() string {
	return "order_demand"
}
