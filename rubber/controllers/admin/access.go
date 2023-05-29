/**
  @Author ZXQ-Administrator
  @Date 2021-08-24
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type AccessController struct {
	BaseController
}

func (c *AccessController) Get() {
	access := []models.Access{}
	//预加载 将module_id=0的数据挂载到AccessItem
	//仅搜索status=1的显示数据
	//models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	models.DB.Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Order("sort desc").Where("module_id=?", 0).Find(&access)
	//加入checkedList属性
	for i := 0; i < len(access); i++ {
		access[i].CheckedList = false
	}
	c.Data["accessList"] = access
	c.TplName = "admin/access/index.html"
}
func (c *AccessController) Add() {
	//加载顶级模块
	access := []models.Access{}
	models.DB.Where("module_id=0").Find(&access)

	c.Data["accessList"] = access
	c.TplName = "admin/access/add.html"
}
func (c *AccessController) DoAdd() {
	//获取页面数据
	moduleName := c.GetString("module_name")
	iType, errType := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, errMI := c.GetInt("module_id")
	sort, errSort := c.GetInt("sort")
	description := c.GetString("description")
	status, errStatus := c.GetInt("status")
	if errType != nil || errMI != nil || errSort != nil || errStatus != nil {
		c.Error("传入参数错误", "/access/add")
		return
	}
	//对页面填入的数据进行校验
	if len(moduleName) < 2 || moduleName == "" {
		c.Error("模块名称须大于2位且不为空", "/access/add")
		return
	}
	if len(actionName) < 2 || actionName == "" {
		c.Error("模块名称须大于2位且不为空", "/access/add")
		return
	}
	if moduleId != 0 && iType == 1 {
		c.Error("只有顶级模块才能增加模块节点！", "/access/add")
		return
	}
	// 判断当前module_name是否存在数据库
	var moduleList []models.Access
	models.DB.Where("module_id=? AND module_name=?", 0, moduleName).Find(&moduleList)
	if len(moduleList) > 1 {
		c.Error("该顶级模块已已存在", "/access/add")
		return
	}
	//增加权限
	access := models.Access{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		c.Error("增加数据失败", "/access/add")
	}
	//c.Goto("/")
	c.Success("增加数据成功", "/access")
}
func (c *AccessController) Edit() {
	//需要修改的数据
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	c.Data["access"] = access

	//获取顶级模块
	var accessList []models.Access
	models.DB.Where("module_id=0").Find(&accessList)
	c.Data["accessList"] = accessList
	c.TplName = "admin/access/edit.html"
}
func (c *AccessController) DoEdit() {
	// 获取修改数据
	id, errId := c.GetInt("id")
	moduleName := c.GetString("module_name")
	iType, errType := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, errMI := c.GetInt("module_id")
	sort, errSort := c.GetInt("sort")
	description := c.GetString("description")
	status, errStatus := c.GetInt("status")
	if errId != nil || errType != nil || errMI != nil || errSort != nil || errStatus != nil {
		c.Error("传入参数错误", "/access/add")
		return
	}
	//对页面填入的数据进行校验
	if len(moduleName) < 2 || moduleName == "" {
		c.Error("模块名称须大于2位且不为空", "/access/add")
		return
	}
	if len(actionName) < 2 || actionName == "" {
		c.Error("模块名称须大于2位且不为空", "/access/add")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = iType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	access.AddTime = int(models.GetUnix())

	err := models.DB.Save(&access).Error
	if err != nil {
		c.Error("修改失败", "/access/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/access/")
}
func (c *AccessController) Delete() {
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	//获取当前数据
	access := models.Access{Id: id}
	models.DB.Find(&access)
	// 顶级模块 判断下层是否存在
	if access.ModuleId == 0 {
		accessed := []models.Access{}
		models.DB.Where("module_id=?", access.Id).Find(&accessed)
		if len(accessed) > 0 {
			c.Error("当前模块下面还有菜单或者操作，无法删除", "/access")
			return
		}
	}
	accessDel := models.Access{Id: id}
	models.DB.Delete(&accessDel)
	c.Success("删除成功", "/access")
}

func (c *AccessController) DoClickAccess() {
	id, errId := c.GetInt("id")
	if errId != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := []models.Access{}
	//预加载 将module_id=0的数据挂载到AccessItem
	//仅搜索status=1的显示数据
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	for i := 0; i < len(access); i++ {
		if id == access[i].Id {
			access[i].CheckedList = !access[i].CheckedList
		}

	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  access,
	}
	c.ServeJSON()
	//获取当前数据
	c.Data["accessList"] = access
	//c.TplName = "admin/access/index.html"

}
