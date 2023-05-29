/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"strconv"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (c *ManagerController) Get() {
	manager := []models.Manager{}
	models.DB.Preload("Role").Find(&manager)
	c.Data["managerList"] = manager

	c.TplName = "admin/manager/index.html"
}
func (c *ManagerController) Add() {
	//获取所有的角色
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	c.TplName = "admin/manager/add.html"
}
func (c *ManagerController) DoAdd() {
	//获取数据
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.Error("非法请求", "/manager/add")
	}
	username := strings.Trim(c.GetString("username"), "")
	password := strings.Trim(c.GetString("password"), "")
	mobile := strings.Trim(c.GetString("mobile"), "")
	email := strings.Trim(c.GetString("email"), "")
	//对页面填入的数据进行校验
	if username ==""|| len(username) < 2 || len(password) < 6 {
		c.Error("用户名长度须大于2位不为空，密码长度须大于等于6", "/manager/add")
		return
	}
	if len(mobile) != 11 {
		c.Error("电话号码长度为11位", "/manager/add")
		return
	}
	//判断当前用户是否在数据库
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		c.Error("用户名已存在", "/manager/add")
		return
	}
	//增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status : 1,
		AddTime : int(models.GetUnix()),
		RoleId : roleId,
	}
	errManager:=models.DB.Create(&manager).Error
	if errManager != nil{
		c.Error("增加管理员失败", "/manager/add")
		return
	}
	c.Success("增加管理员成功","/manager")
}
func (c *ManagerController) Edit() {
	//获取管理员信息
	id,err:=c.GetInt("id")
	if err!=nil{
		c.Error("非法请求", "/manager")
		return
	}
	//获取manager信息
	manager:=models.Manager{Id:id}
	models.DB.Find(&manager)
	c.Data["manager"]=manager
	//获取所有的角色
	role:=[]models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"]=role
	c.TplName = "admin/manager/edit.html"
}
func (c *ManagerController) DoEdit() {
	id,errId:=c.GetInt("id")
	if errId !=nil{
		c.Error("非法请求","/manager")
		return
	}
	roleId,errRoleId :=c.GetInt("role_id")
	if errRoleId != nil{
		c.Error("非法请求", "/manager")
		return
	}
	mobile:=strings.Trim(c.GetString("mobile"),"")
	email := strings.Trim(c.GetString("email"), " ")
	password := strings.Trim(c.GetString("password"), " ")

	//获取数据
	manager :=models.Manager{Id:id}
	models.DB.Find(&manager)
	manager.RoleId=roleId
	manager.Mobile=mobile
	manager.Email=email
	if password!=""{
		if len(password) < 6 {
			c.Error("密码长度不合法,密码长度不能小于6位", "/manager/edit?id="+strconv.Itoa(id))
			return
		}
		// 密码经MD5加密后传入
		manager.Password = models.Md5(password)
	}
	//执行修改
	err:=models.DB.Save(&manager).Error
	if err!=nil{
		c.Error("请检查您修改的数据格式", "/manager/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/manager")
	}
}
func (c *ManagerController) Delete() {
	id,err :=c.GetInt("id")
	if err!=nil{
		c.Error("传入参数错误", "/manager")
		return
	}
	manager := models.Manager{Id: id}

	models.DB.Delete(&manager)
	c.Success("删除管理员成功", "/manager")
}
