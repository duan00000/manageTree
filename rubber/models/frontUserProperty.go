/**
  @Author ZXQ-Administrator
  @Date 2021-10-07
  @node: 用户资产情况 可根据自身需求添加更多属性
**/
package models

type FrontUserProperty struct {
	Id       int     `json:"id"`
	Craft    int `json:"craft"`
	AddTime  int     `json:"add_time"`
	West    int `json:"west"`
	Mid    int `json:"mid"`
	East    int `json:"east"`
	Trace    int `json:"trace"`
}

func (FrontUserProperty) TableName() string {
	return "frontuserproperty"
}
