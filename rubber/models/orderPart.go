/**
  @Author ZXQ-Administrator
  @Date 2022-02-19
  @node:
**/
package models

type OrderPart struct {
	Id        int    `json:"id"`
	PartA     string `json:"phone"`
	PartB    string `json:"password"`
}

func (OrderPart) TableName() string {
	return "order_part"
}
