/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type DsmUserDevice struct {
	Id            int       `json:"id"`
	NextDevice   string  `json:"next_device"`  //参与设备
	NextPoint    int      `json:"next_point"`	//监测点
	NextVolume   float64  `json:"next_volume"`	//设备容量
	NextSee      int      `json:"next_see"`	//是否可监测
	NextJoin     int 	   `json:"next_join"`	//是否参与
	NextMethod   int 	   `json:"next_method"`	//参与方式
	tUid         int 	   `json:"t_uid"`	//
	SubDeviceId int      `json:"sub_device_id"` //子设备id
	NextRespond  float64  `json:"next_respond"` //
	Status        int      `json:"status"`
	AddTime      int      `json:"add_time"`
}

func (DsmUserDevice) TableName() string {
	return "dsm_user_device"
}
