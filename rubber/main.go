/**
  @Author ZXQ-Administratorgit上传文件到gitlab
  @Date 2021-08-18
  @node:
**/
package main

import (
	"elenewenergy/models"
	_ "elenewenergy/routers"
	"encoding/gob"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)
// https://beego.me/docs/mvc/controller/session.md
//因为 session 内部采用了 gob 来注册存储的对象，例如 struct，所以如果你采用了非 memory 的引擎，请自己在 main.go 的 init 里面注册需要保存的这些结构体，不然会引起应用重启之后出现无法解析的错误
// 错误：gob: name not registered for interface: "xxxxxxx/models.Manager"
func init()  {
	gob.Register(models.Manager{})
}
func main() {
	//注册模板函数
	beego.AddFuncMap("unixToDate",models.UnixToDate)
	beego.AddFuncMap("setting",models.GetSettingFromColumn)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("formatAttr", models.FormatAttr)
	beego.AddFuncMap("mul", models.Mul)

	//后台配置允许跨域
	//https://github.com/astaxie/beego/tree/master/plugins/cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowOrigins:配置 "*" 表示所有的域都可以访问后台
		//				AllowOrigins:     []string{"http://localhost:8081","http://172.16.111.67:8081"},
		//AllowOrigins:     []string{"*"},
		AllowAllOrigins: true,		// 允许请求的域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//允许头部信息

		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//允许暴露的头信息
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true, //是否允许cookie跨域
	}))

	//配置session保存在redis
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
	defer models.DB.Close()
}
