/**
  @Author ZXQ-Administrator
  @Date 2021-12-29
  @node:Vue公司信息
**/
package models

import (
	_ "github.com/jinzhu/gorm"
)

type VueCompany struct {
	Id         int    `json:"id"`
	CKind      string `json:"c_kind"`      //公司类型
	CUid       int    `json:"c_uid"`       //对应user表的id
	TUid       int    `json:"t_uid"`       //总用户号
	UserType   int    `json:"user_type"`   //用户类型
	ContractNo string `json:"contract_no"` //签约编号
	CPhone     string `json:"c_phone"`     //联系电话
	CName      string `json:"c_name"`      //公司名称
	CAddress   string `json:"c_address"`   //公司地址
	CLaw       string `json:"c_law"`       //公司法人
	CCount     int    `json:"c_count"`     //公司数量
	CCapital   string `json:"c_capital"`   //注册资本
	Status	int `json:"status"`
	AddTime    int    `json:"add_time"`
}

func (VueCompany) TableName() string {
	return "vue_company"
}
