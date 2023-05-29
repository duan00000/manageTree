/**
  @Author ZXQ-Administrator
  @Date 2021-10-01
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type FrontUserController struct {
	BaseController
}


func (c *FrontUserController) Get() {
	//当前分页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每页显示数量
	pageSize := 10

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND username like \"%" + keyword + "%\""
	}
	//查询数据
	//查询数据
	var frontUserList []models.User
	//select * from tableName limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&frontUserList)

	//查询user表里面的用户
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("user").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(frontUserList) == 0  && count != 0{
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/user?page=" + strconv.Itoa(prvPageBack))
	}else{
		//	商品为空
	}
	c.Data["frontUserList"] = frontUserList
	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword
	c.TplName = "admin/frontUser/index.html"
}

func (c *FrontUserController) Add() {
	c.TplName = "admin/frontUser/add.html"
}

func (c *FrontUserController) DoAdd() {
	username := strings.Trim(c.GetString("username"), "")
	password := strings.Trim(c.GetString("password"), "")
	phone := strings.Trim(c.GetString("phone"), "")
	email := strings.Trim(c.GetString("email"), "")
	//对页面填入的数据进行校验
	if username ==""|| len(username) < 2 || len(password) < 6 {
		c.Error("用户名长度须大于2位不为空，密码长度须大于等于6", "/frontUser/add")
		return
	}
	if len(phone) != 11 {
		c.Error("电话号码长度为11位", "/frontUser/add")
		return
	}
	phoneExit := []models.User{}
	models.DB.Where("phone=?", phone).Find(&phoneExit)
	if len(phoneExit) > 0 {
		c.Error("手机号已注册", "/frontUser/add")
		return
	}
	//用户名是否已注册
	usernameExit := []models.User{}
	models.DB.Where("username=?", username).Find(&usernameExit)
	//验证用户名格式是否正确 中文、大小写字母和数字 字符长度2-50
	pattern := "[a-zA-Z0-9\u4e00-\u9fa5]{3,8}${2,50}$"
	reg := regexp.MustCompile(pattern)
	if len(usernameExit) > 0 {
		c.Error("用户名已存在", "/frontUser/add")
		return
	}else if !reg.MatchString(username) {
		c.Error("用户名格式不正确", "/frontUser/add")
		return
	}else{

	}
	//增加用户
	frontUser := models.User{
		Username: username,
		Password: models.Md5(password),
		Phone:   phone,
		Email:    email,
		Status : 0,
		AddTime : int(models.GetUnix()),
	}
	errFrontUser:=models.DB.Create(&frontUser).Error
	if errFrontUser != nil{
		c.Error("增加用户失败", "/frontUser/add")
		return
	}
	//给用户资产赋初始值值
	userProperty:=models.User{}
	models.DB.Where("phone=?", phone).Find(&userProperty)
	frontUserProperty:=models.FrontUserProperty{
		Id: userProperty.Id,
		Craft: userProperty.Craft,
		AddTime:  int(models.GetUnix()),
		West: userProperty.West,
		Mid: userProperty.Mid,
		East: userProperty.East,
		Trace: userProperty.Trace,

	}
	errInfo:=models.DB.Create(&frontUserProperty).Error
	if errInfo != nil{
		c.Error("增加用户资产失败", "/frontUser/add")
		return
	}
	c.Success("增加用户成功","/frontUser")

}
func (c *FrontUserController) Edit() {
	//获取管理员信息
	id,err:=c.GetInt("id")
	if err!=nil{
		c.Error("非法请求", "/frontUser")
		return
	}
	//获取user信息
	frontUser:=models.User{Id:id}
	models.DB.Find(&frontUser)
	c.Data["frontUser"]=frontUser
	//获取所有的角色
	c.TplName = "admin/frontUser/edit.html"
}

func (c *FrontUserController) DoEdit() {
	//1、获取要修改的用户数据
	id,errId:=c.GetInt("id")
	if errId !=nil{
		c.Error("非法请求","/frontUser")
		return
	}
	phone:=strings.Trim(c.GetString("phone"),"")
	email := strings.Trim(c.GetString("email"), " ")
	password := strings.Trim(c.GetString("password"), " ")
	username := strings.Trim(c.GetString("username"), "")

	//对页面填入的数据进行校验
	if username ==""|| len(username) < 2 || len(password) < 6 {
		c.Error("用户名长度须大于2位不为空，密码长度须大于等于6", "/frontUser/edit?id="+strconv.Itoa(id))
		return
	}

	if len(phone) != 11 {
		c.Error("电话号码长度为11位", "/frontUser/edit?id="+strconv.Itoa(id))
		return
	}
	phoneExit := []models.User{}
	models.DB.Where("phone=?", phone).Find(&phoneExit)
	if len(phoneExit) > 0 {
		c.Error("手机号已注册", "/frontUser/edit?id="+strconv.Itoa(id))
		return
	}
	//用户名是否已注册
	usernameExit := []models.User{}
	models.DB.Where("username=?", username).Find(&usernameExit)
	//验证用户名格式是否正确 中文、大小写字母和数字 字符长度2-50
	pattern := "[a-zA-Z0-9\u4e00-\u9fa5]{3,8}${2,50}$"
	reg := regexp.MustCompile(pattern)
	if len(usernameExit) > 0 {
		c.Error("用户名已存在", "/frontUser/edit?id="+strconv.Itoa(id))
		return
	}else if !reg.MatchString(username) {
		c.Error("用户名格式不正确", "/frontUser/edit?id="+strconv.Itoa(id))
		return
	}else{

	}
	//获取数据
	frontUser :=models.User{Id:id}
	models.DB.Find(&frontUser)
	frontUser.Username=username
	frontUser.Phone=phone
	frontUser.Email=email
	if password!=""{
		if len(password) < 6 {
			c.Error("密码长度不合法,密码长度不能小于6位", "/frontUser/edit?id="+strconv.Itoa(id))
			return
		}
		// 密码经MD5加密后传入
		frontUser.Password = models.Md5(password)
	}
	//执行修改
	err:=models.DB.Save(&frontUser).Error
	if err!=nil{
		c.Error("请检查您修改的数据格式", "/frontUser/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/frontUser")
	}

}

func (c *FrontUserController) Delete() {
	id,err :=c.GetInt("id")
	if err!=nil{
		c.Error("传入参数错误", "/frontUser")
		return
	}
	frontUser := models.User{Id: id}
	models.DB.Delete(&frontUser)
	frontUserProperty := models.FrontUserProperty{Id: id}
	models.DB.Delete(&frontUserProperty)

	c.Success("删除用户成功", "/frontUser")
}
