/**
  @Author ZXQ-Administrator
  @Date 2021-09-23
  @node: 注册用户 可根据自身需求添加更多属性
**/
package models



type TreeInfo struct {
	Id       int     `json:"id"`
	TreeId   string `json:"treeid"`
	TreeName string `json:"treename"`

}

func (TreeInfo) TableName() string {
	return "tree_info"
}

