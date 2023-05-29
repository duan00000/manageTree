/**
  @Author ZXQ-Administrator
  @Date 2021-08-24
  @node:
**/
package models

type Access struct {
	Id          int`json:"id"`
	ModuleName  string 	`json:"module_name"`	//模块名称
	Type        int		`json:"type"`	//节点类型：1.表示模块  2.表示菜单  3.操作
	ActionName  string	`json:"action_name"`	//操作名称
	Url         string	`json:"url"`	//路由跳转地址
	ModuleId    int		`json:"module_id"`	//此module_id和当前模型的_id关联  module_id=0 表示模块
	Sort        int		`json:"sort"`	//排序
	Description string`json:"description"`
	AddTime     int`json:"add_time"`
	Status      int`json:"status"`
	AccessItem  []Access `gorm:"foreignkey:ModuleId;association_foreignkey:Id" json:"access_item"`
	Checked       bool	`gorm:"-" json:"checked"` 	//忽略本字段 main.go
	CheckedList   bool	`gorm:"-" json:"checked_list"` 	//忽略本字段 access.go

}

func (Access) TableName() string {
	return "access"
}
