/**
  @Author ZXQ-Administrator
  @Date 2021-09-22
  @node: 前端 购物车 缓存
**/
package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

//定义结构体 缓存结构体 私有
type cookie struct {}


//写入数据的方法
func (c cookie) Set(ctx *context.Context, key string, value interface{}) {
	//app.conf 文件配置
	var cookieMaxTime, _ = beego.AppConfig.Int("cookieMaxTime")
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(beego.AppConfig.String("secureCookie"), key, string(bytes), cookieMaxTime)

}

//获取数据的方法
func (c cookie) Get(ctx *context.Context, key string, obj interface{}) bool {
	tempData, ok := ctx.GetSecureCookie(beego.AppConfig.String("secureCookie"), key)
	if !ok {
		return false
	}
	json.Unmarshal([]byte(tempData), obj)
	return true
}

// 去除数据的方法
func (c cookie) Remove(ctx *context.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(beego.AppConfig.String("secureCookie"), key, string(bytes), -1)

}


//实例化结构体
var Cookie = &cookie{}
