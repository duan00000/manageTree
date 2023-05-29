/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: api 路由
**/
package routers

import (
	v1 "elenewenergy/controllers/api/v1"
	"elenewenergy/controllers/api/v2"
	"github.com/astaxie/beego"
)

func init() {

}
func init() {
	ns1 :=
		beego.NewNamespace("/api/v1",
			//版本1
			//beego.NSRouter("/", &api.V1Controller{}),
			//beego.NSRouter("/nav", &api.V1Controller{}, "get:Nav"),
			beego.NSRouter("/login", &v1.VueLoginController{}),
			//设备属性
			beego.NSRouter("/reducer", &v1.ReducerController{}, "get:Reducer"),
			beego.NSRouter("/doLogin", &v1.VueLoginController{}, "post:DoLogin"),
			beego.NSRouter("/getProfile", &v1.VueLoginController{}, "post:GetProfile"),
			beego.NSRouter("/sendCode", &v1.VuePassController{}, "get:SendCode"),
			beego.NSRouter("/registerSms", &v1.VuePassController{}, "post:RegisterSms"),

			beego.NSRouter("/registerStep", &v1.VuePassController{}, "post:RegisterStep"),

			//用户信息
			beego.NSRouter("/getUserInfo", &v1.UserInfoController{}, "get:GetUserInfo"),
			beego.NSRouter("/putUserInfo", &v1.UserInfoController{}, "post:PutUserInfo"),
			//公司信息
			beego.NSRouter("/getCompanyInfo", &v1.UserInfoController{}, "get:GetCompanyInfo"),
			beego.NSRouter("/putCompanyInfo", &v1.UserInfoController{}, "post:PutCompanyInfo"),

			//需求信息填写
			beego.NSRouter("/getDemandBasic", &v1.VueDemandController{}, "get:GetDemandBasic"),
			beego.NSRouter("/getDemandBasicInfo", &v1.VueDemandController{}, "get:GetDemandBasicInfo"),
			beego.NSRouter("/putBasicInfo", &v1.VueDemandController{}, "post:PutBasicInfo"),
			beego.NSRouter("/doEditUserDevice", &v1.VueDemandController{}, "post:DoEditUserDevice"),
			beego.NSRouter("/getDsmDevice", &v1.VueDemandController{}, "get:GetDsmDevice"),
			beego.NSRouter("/getRespInfo", &v1.VueDemandController{}, "get:GetRespInfo"),
			beego.NSRouter("/putRespInfo", &v1.VueDemandController{}, "post:PutRespInfo"),
			beego.NSRouter("/submitAllInfo", &v1.OrderController{}, "post:SubmitAllInfo"),

			beego.NSRouter("/getOrderDemand", &v1.VueDemandController{}, "get:GetOrderDemand"),
			beego.NSRouter("/abandonOrder", &v1.VueDemandController{}, "post:AbandonOrder"),


			//beego.NSRouter("/doEdit", &v1.VueLoginController{}, "put:DoEdit"),
			//beego.NSRouter("/deleteNav", &v1.VueLoginController{}, "delete:DeleteNav"),
			//beego.NSRouter("/setUser", &v1.VueLoginController{}, "get:SetUser"),
			//beego.NSRouter("/getUser", &v1.VueLoginController{}, "get:GetUser"),
		)
	ns2 :=
		beego.NewNamespace("/api/v2",
			//版本2
			beego.NSRouter("/", &v2.V2Controller{}),
		)
	//注册 namespace
	beego.AddNamespace(ns1, ns2)
}