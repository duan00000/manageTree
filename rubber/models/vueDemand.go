/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node: 需求响应 订单信息
**/
package models

type VueDemand struct {
	Id          int    `json:"id"`
	Uid         int    `json:"uid"`      //总用户id
	Did         int    `json:"did"`      //设备id
	UserType    string `json:"userType"` //用户类型：1,普通新用户 2,普通老用户，3负荷集成商
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	BasicId     int    `json:"basic_id"`
	DeviceId    int    `json:"device_id"`
	CapId       int    `json:"cap_id"`
	FirTime     int    `json:"fir_id"`
	SecTime     int    `json:"sec_id"`
	ThiTime     int    `json:"thi_id"`
	FouTime     int    `json:"fou_id"`
	orderId     int    `json:"order_id"`
	ProcessTime int    `json:"processTime"`
	FinishTime  int    `json:"finishTime"`
	OrderStatus string `json:"orderStatus"` //订单状态：用户申请（进行：0）（完成：1）（修改：10）","市级初审(（进行：2）（完成：3）（修改：20）)","省级复审(（进行：4）（完成：5）（修改：40）)","签订协议(（进行：6）（完成：7）（修改：60）)","进入用户首页(（进行：8）（完成：9）（修改：80）)
	AddTime     int    `json:"addTime"`
}

func (VueDemand) TableName() string {
	return "vue_demand"
}
