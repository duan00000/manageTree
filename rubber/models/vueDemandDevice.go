/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type VueDemandDevice struct {
	Id           int     `json:"id"`
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
	AddTime      int     `json:"addTime"`
}

func (VueDemandDevice) TableName() string {
	return "vue_demand_device"
}
