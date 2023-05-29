/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type VueDemandCapability struct {
	Id          int     `json:"id"`
	CUid        int     `json:"cUid"`
	OrderId     int     `json:"order_id"`
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
	AddTime     int     `json:"add_time"`
}

func (VueDemandCapability) TableName() string {
	return "vue_demand_cap"
}
