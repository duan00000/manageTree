/**
  @Author ZXQ-Administrator
  @Date 2021-12-29
  @node:Vue用户信息
**/
package models

import (
	_ "github.com/jinzhu/gorm"
)

type VueUser struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	CompanyName string `json:"companyName"`
	CompanyDesc string `json:"companyDesc"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	AddTime     int    `json:"addTime"`
	LastIp      string `json:"lastIp"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
	Captcha     string `gorm:"-" json:"captcha"`
	InputCap    string `gorm:"-" json:"inputCap"`
}

func (VueUser) TableName() string {
	return "vue_user"
}
