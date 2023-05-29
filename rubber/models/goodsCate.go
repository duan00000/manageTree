/**
  @Author ZXQ-Administrator
  @Date 2021-09-04
  @node:
**/
package models

type GoodsCate struct {
	Id          int
	Title       string
	CateImg     string
	Link        string
	Template    string
	Pid         int
	SubTitle    string
	Keywords    string
	Description string
	Status      int
	Sort        int
	AddTime     int
	GoodsCateItem  []GoodsCate `gorm:"foreignkey:Pid;association_foreignkey:Id"`

}

func (GoodsCate) TableName() string{
	return "goods_cate"
}