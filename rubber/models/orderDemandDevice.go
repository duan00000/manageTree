/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type OrderDemandDevice struct {
	Id           int     `json:"id"`
	VueDeviceId    int     `json:"vue_device_id"`  //与vue_demand_basic 中的id绑定
	CName        string  `json:"c_name"`
	CAddress     string  `json:"c_address"`
	CPhone       string  `json:"c_phone"`
	TUid         int     `json:"t_uid"`
	UserType     int     `json:"user_type"`
	BeResponse   float64 `json:"be_response"`
	SubId        int     `json:"sub_id"`
	SubDeviceId  int     `json:"sub_device_id"`
	DeviceId     int     `json:"device_id"`
	TimeRespond1 float64 `json:"time_respond1"`
	TimeRespond2 float64 `json:"time_respond2"`
	TimeRespond3 float64 `json:"time_respond3"`
	TimeRespond4 float64 `json:"time_respond4"`
	Status       int     `json:"status"`
	OrderId       int     `json:"order_id"`
	AddTime      int     `json:"addTime"`
}

func (OrderDemandDevice) TableName() string {
	return "order_demand_device"
}
