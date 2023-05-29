/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type OrderDemandCapability struct {
	Id          int     `json:"id"`
	CUid        int     `json:"cUid"`
	VueCapId    int     `json:"vue_cap_id"`  //与vue_demand_cap 中的id绑定
	MTFz        float64 `json:"mTFz"`         //4小时以上凌晨
	MTFm        float64 `json:"mTFm"`         //4小时以上早
	MTFw        float64 `json:"mTFw"`         //4小时以上腰
	MTFn        float64 `json:"mTFn"`         //4小时以上晚
	WFz         float64 `json:"wFz"`          //4小时以上凌晨
	WFm         float64 `json:"wFm"`          //4小时内早
	WFw         float64 `json:"wFw"`          //4小时内腰
	WFn         float64 `json:"wFn"`          //4小时内晚
	CanResponse int     `json:"can_response"` //能够实施响应
	Status      int     `json:"status"`
	OrderId      int     `json:"order_id"`
	AddTime     int     `json:"add_time"`
}

func (OrderDemandCapability) TableName() string {
	return "order_demand_cap"
}
