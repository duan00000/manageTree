/**
  @Author ZXQ-Administrator
  @Date 2022-01-16
  @node: DSM 子用户信息
**/
package models

type DsmSubUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	UserId    int    `json:"user_id"`     //加入年份 当年
	Area       string `json:"area"`
	SignedInt string `json:"signed_int"`
	CompanyId string `json:"company_id"`
	DeviceId  int    `json:"device_id"`
	AddTime   int    `json:"add_time"` //合同容量
	Status     int    `json:"status"`  //运行容量
}

func (DsmSubUser) TableName() string {
	return "dsm_sub_user"
}
