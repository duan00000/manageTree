/**
  @Author ZXQ-Administrator
  @Date 2021-09-14
  @node:
**/
package models

type Nav struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Link      string   `json:"link"`
	Position  int      `json:"position"`
	IsOpennew int      `json:"is_opennew"`
	Relation  string   `json:"relation"`
	Sort      int      `json:"sort"`
	Status    int      `json:"status"`
	AddTime   int      `json:"add_time"`
	GoodsItem []Goods `gorm:"-" json:"goods_item"`         //关联商品的数据

}


func (Nav) TableName() string {
	return "nav"
}
