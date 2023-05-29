/**
  @Author ZXQ-Administrator
  @Date 2021-09-05
  @node:
**/
package models

type GoodsType struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}

func (GoodsType) TableName() string {
	return "goods_type"
}
