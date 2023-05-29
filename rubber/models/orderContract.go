/**
  @Author ZXQ-Administrator
  @Date 2022-02-19
  @node:
**/
package models

type OrderContract struct {
	Id         int    `json:"id"`
	CPhone     string `json:"c_phone"`     //联系电话
	CName      string `json:"c_name"`      //公司名称
	CAddress   string `json:"c_address"`   //公司地址
	CLaw       string `json:"c_law"`       //公司法人
	Fax        string `json:"fax"`         //传真
	Email      string  `json:"email"`
	LinkUser   string  `json:"link_user"`
	OrderId     int    `json:"order_id"`
	SignContract int   `json:"sign_contract"` //签订协议 //0 默认 1通过 2拒绝
	AddTime    int    `json:"add_time"`
}

func (OrderContract) TableName() string {
	return "order_contract"
}

