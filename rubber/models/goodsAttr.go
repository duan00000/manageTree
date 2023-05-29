/**
  @Author ZXQ-Administrator
  @Date 2021-09-10
  @node:
**/
package models

type GoodsAttr struct {
	Id              int
	GoodsId         int
	AttributeCateId int
	AttributeId     int
	AttributeTitle  string
	AttributeType   int
	AttributeValue  string
	Sort            int
	AddTime         int
	Status          int
}
func (GoodsAttr) TableName()string{
	return "goods_attr"
}