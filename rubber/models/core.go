/**
  @Author ZXQ-Administrator
  @Date 2021-08-18
  @node: 配置数据库的连接
**/
package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//声明全局  数据库连接
var DB *gorm.DB
var err error

func init() {
	mysqladmin:=beego.AppConfig.String("mysqladmin")
	mysqlpwd:=beego.AppConfig.String("mysqlpwd")
	mysqldb:=beego.AppConfig.String("mysqldb")
	/*
	如果想指定远程主机，你需要使用 (). 例如:
		user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	*/
	DB,err = gorm.Open("mysql",mysqladmin+":"+mysqlpwd+"@/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		beego.Error(err)
	}
	DB.LogMode(true)
}
