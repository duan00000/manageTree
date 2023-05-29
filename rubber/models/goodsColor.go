/**
  @Author ZXQ-Administrator
  @Date 2021-09-09
  @node:
**/
package models

type GoodsColor struct {
	Id         int
	ColorName  string
	ColorValue string
	Status     int
	Checked    bool	`gorm:"-"` 	//忽略本字段
}
func (GoodsColor) TableName() string{
	return "goods_color"
}