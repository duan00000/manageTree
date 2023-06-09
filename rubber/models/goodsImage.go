/**
  @Author ZXQ-Administrator
  @Date 2021-09-10
  @node:
**/
package models

type GoodsImage struct {
	Id      int           `json:"id"`
	GoodsId int           `json:"goods_id"`
	ImgUrl  string        `json:"img_url"`
	ColorId int           `json:"color_id"`
	Sort    int           `json:"sort"`
	AddTime int           `json:"add_time"`
	Status  int           `json:"status"`
}
func (GoodsImage) TableName()string{
	return "goods_image"
}
