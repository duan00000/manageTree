/**
  @Author ZXQ-Administrator
  @Date 2021-09-23
  @node:注册用户 判断发送短信次数
**/
package models

type UserTemp struct {
	Id        int
	Ip        string
	Phone     string
	SendCount int
	AddDay    string
	AddTime   int
	Sign      string
}

func (UserTemp) TableName() string {
	return "user_temp"
}

