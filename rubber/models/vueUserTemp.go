/**
  @Author ZXQ-Administrator
  @Date 2021-12-29
  @node:VUE前端注册
**/
package models

type VueUserTemp struct {
	Id        int    `json:"id"`
	Ip        string `json:"ip"`
	Phone     string `json:"phone"`
	SendCount int    `json:"sendCount"`
	AddDay    string `json:"addDay"`
	AddTime   int    `json:"addTime"`
	Sign      string `json:"sign"`
	SmsCode      string `json:"smsCode"`

}

func (VueUserTemp) TableName() string {
	return "vue_user_temp"
}
