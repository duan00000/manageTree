/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: 后端 路由
**/
package routers

import (
	"elenewenergy/controllers/admin"
	"elenewenergy/controllers/admin/elePowDR"

	"elenewenergy/middleware"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/"+beego.AppConfig.String("adminPath"),
			//中间件:匹配路由前会执,可以用于权限验证
			//注意引入的包： github.com/astaxie/beego/context
			beego.NSBefore(middleware.AdminAuth),
			//后台管理
			beego.NSRouter("/", &admin.MainController{}),
			beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),
			//权限不足
			beego.NSRouter("/middleware", &admin.MainController{}, "get:Middleware"),
			//修改状态
			beego.NSRouter("/main/changeStatus", &admin.MainController{}, "get:ChangeStatus"),
			//修改数字
			beego.NSRouter("/main/editNum", &admin.MainController{}, "get:EditNum"),

			//登录管理
			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),
			beego.NSRouter("/login/loginOut", &admin.LoginController{}, "get:LoginOut"),

			//管理员管理
			beego.NSRouter("/manager", &admin.ManagerController{}),
			beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
			beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
			beego.NSRouter("/manager/doAdd", &admin.ManagerController{}, "post:DoAdd"),
			beego.NSRouter("/manager/doEdit", &admin.ManagerController{}, "post:DoEdit"),
			beego.NSRouter("/manager/delete", &admin.ManagerController{}, "get:Delete"),

			//需求响应管理
			beego.NSRouter("/demandAnalysis", &elePowDR.DemandAnalyController{}),
			beego.NSRouter("/responseManagement", &elePowDR.ResponseManageController{}),
			beego.NSRouter("/responseManagement/verify", &elePowDR.ResponseManageController{},"get:Verify"),
			beego.NSRouter("/responseManagement/doVerify", &elePowDR.ResponseManageController{},"post:DoVerify"),
			beego.NSRouter("/responseManagement/doConfirm", &elePowDR.ResponseManageController{},"post:DoConfirm"),
			beego.NSRouter("/responseManagement/delete", &elePowDR.ResponseManageController{},"get:Delete"),

			beego.NSRouter("/signContract", &elePowDR.SignContractController{}),
			beego.NSRouter("/signContract/sign", &elePowDR.SignContractController{},"get:Sign"),
			beego.NSRouter("/signContract/doSign", &elePowDR.SignContractController{},"post:DoSign"),



			//beego.NSRouter("/login", &admin.BaseController{},"get:Success"),
			//轮播图管理
			beego.NSRouter("/focus", &admin.FocusController{}),
			beego.NSRouter("/focus/add", &admin.FocusController{}, "get:Add"),
			beego.NSRouter("/focus/edit", &admin.FocusController{}, "get:Edit"),
			beego.NSRouter("/focus/doAdd", &admin.FocusController{}, "post:DoAdd"),
			beego.NSRouter("/focus/doEdit", &admin.FocusController{}, "post:DoEdit"),
			beego.NSRouter("/focus/delete", &admin.FocusController{}, "get:Delete"),

			//角色管理
			beego.NSRouter("/role", &admin.RoleController{}),
			beego.NSRouter("/role/add", &admin.RoleController{}, "get:Add"),
			beego.NSRouter("/role/edit", &admin.RoleController{}, "get:Edit"),
			beego.NSRouter("/role/doAdd", &admin.RoleController{}, "post:DoAdd"),
			beego.NSRouter("/role/doEdit", &admin.RoleController{}, "post:DoEdit"),
			beego.NSRouter("/role/delete", &admin.RoleController{}, "get:Delete"),
			beego.NSRouter("role/auth", &admin.RoleController{}, "get:Auth"),
			beego.NSRouter("role/doAuth", &admin.RoleController{}, "post:DoAuth"),

			//权限管理
			beego.NSRouter("/access", &admin.AccessController{}),
			beego.NSRouter("/access/add", &admin.AccessController{}, "get:Add"),
			beego.NSRouter("/access/edit", &admin.AccessController{}, "get:Edit"),
			beego.NSRouter("/access/doAdd", &admin.AccessController{}, "post:DoAdd"),
			beego.NSRouter("/access/doEdit", &admin.AccessController{}, "post:DoEdit"),
			beego.NSRouter("/access/delete", &admin.AccessController{}, "get:Delete"),
			beego.NSRouter("/access/doClickAccess", &admin.AccessController{}, "get:DoClickAccess"),


			//商品分类管理
			beego.NSRouter("/goodsCate", &admin.GoodsCateController{}),
			beego.NSRouter("/goodsCate/add", &admin.GoodsCateController{}, "get:Add"),
			beego.NSRouter("/goodsCate/edit", &admin.GoodsCateController{}, "get:Edit"),
			beego.NSRouter("/goodsCate/doAdd", &admin.GoodsCateController{}, "post:DoAdd"),
			beego.NSRouter("/goodsCate/doEdit", &admin.GoodsCateController{}, "post:DoEdit"),
			beego.NSRouter("/goodsCate/delete", &admin.GoodsCateController{}, "get:Delete"),

			//商品类型管理
			beego.NSRouter("/goodsType", &admin.GoodsTypeController{}),
			beego.NSRouter("/goodsType/add", &admin.GoodsTypeController{}, "get:Add"),
			beego.NSRouter("/goodsType/edit", &admin.GoodsTypeController{}, "get:Edit"),
			beego.NSRouter("/goodsType/doAdd", &admin.GoodsTypeController{}, "post:DoAdd"),
			beego.NSRouter("/goodsType/doEdit", &admin.GoodsTypeController{}, "post:DoEdit"),
			beego.NSRouter("/goodsType/delete", &admin.GoodsTypeController{}, "get:Delete"),

			//商品类型属性管理
			beego.NSRouter("/goodsTypeAttribute", &admin.GoodsTypeAttrController{}),
			beego.NSRouter("/goodsTypeAttribute/add", &admin.GoodsTypeAttrController{}, `get:Add`),
			beego.NSRouter("/goodsTypeAttribute/edit", &admin.GoodsTypeAttrController{}, `get:Edit`),
			beego.NSRouter("/goodsTypeAttribute/doAdd", &admin.GoodsTypeAttrController{}, `post:DoAdd`),
			beego.NSRouter("/goodsTypeAttribute/doEdit", &admin.GoodsTypeAttrController{}, `post:DoEdit`),
			beego.NSRouter("/goodsTypeAttribute/delete", &admin.GoodsTypeAttrController{}, `get:Delete`),

			//商品管理
			beego.NSRouter("/goods", &admin.GoodsController{}),
			beego.NSRouter("/goods/add", &admin.GoodsController{}, `get:Add`),
			beego.NSRouter("/goods/edit", &admin.GoodsController{}, `get:Edit`),
			beego.NSRouter("/goods/doAdd", &admin.GoodsController{}, `post:DoAdd`),
			beego.NSRouter("/goods/doEdit", &admin.GoodsController{}, `post:DoEdit`),
			beego.NSRouter("/goods/delete", &admin.GoodsController{}, `get:Delete`),
			beego.NSRouter("/goods/doUpload", &admin.GoodsController{}, `post:DoUpload`),
			beego.NSRouter("/goods/doUploadPhoto", &admin.GoodsController{}, `post:DoUploadPhoto`),
			beego.NSRouter("/goods/getGoodsTypeAttribute", &admin.GoodsController{}, `get:GetGoodsTypeAttribute`),
			beego.NSRouter("/goods/changeGoodsImageColor", &admin.GoodsController{}, `get:ChangeGoodsImageColor`),
			beego.NSRouter("/goods/removeGoodsImage", &admin.GoodsController{}, `get:RemoveGoodsImage`),

			//导航管理
			beego.NSRouter("/nav", &admin.NavController{}),
			beego.NSRouter("/nav/add", &admin.NavController{}, `get:Add`),
			beego.NSRouter("/nav/edit", &admin.NavController{}, `get:Edit`),
			beego.NSRouter("/nav/doAdd", &admin.NavController{}, `post:DoAdd`),
			beego.NSRouter("/nav/doEdit", &admin.NavController{}, `post:DoEdit`),
			beego.NSRouter("/nav/delete", &admin.NavController{}, `get:Delete`),

			//商店设置
			beego.NSRouter("/setting", &admin.SettingController{}),
			beego.NSRouter("/setting/doEdit", &admin.SettingController{}, `post:DoEdit`),

			//前端用户管理
			beego.NSRouter("/frontUser", &admin.FrontUserController{}),
			beego.NSRouter("/frontUser/add", &admin.FrontUserController{}, "get:Add"),
			beego.NSRouter("/frontUser/edit", &admin.FrontUserController{}, "get:Edit"),
			beego.NSRouter("/frontUser/doAdd", &admin.FrontUserController{}, "post:DoAdd"),
			beego.NSRouter("/frontUser/doEdit", &admin.FrontUserController{}, "post:DoEdit"),
			beego.NSRouter("/frontUser/delete", &admin.FrontUserController{}, "get:Delete"),

			//前端资产管理
			//前端资产管理
			beego.NSRouter("/frontUserProperty", &admin.FrontUserPropertyController{}),
			beego.NSRouter("/frontUserProperty/delete", &admin.FrontUserPropertyController{}, `get:Delete`),

			//qkl
			beego.NSRouter("/queryOrderItem", &elePowDR.ResponseManageController{},"get:QueryOrderItem"),
			beego.NSRouter("/queryOrderItem1", &elePowDR.ResponseManageController{},"get:QueryOrderItem1"),
			beego.NSRouter("/queryOrderItem2", &elePowDR.ResponseManageController{},"get:QueryOrderItem2"),
			beego.NSRouter("/createOrderItem", &elePowDR.ResponseManageController{},"get:CreateOrderItem"),
			//qkl1
			//qkl
			beego.NSRouter("/treeInformation", &elePowDR.TreeController{}),
			beego.NSRouter("/treeInformation/queryTree1", &elePowDR.TreeController{},"get:QueryTree1"),
			beego.NSRouter("/treeInformation/queryTree2", &elePowDR.TreeController{},"get:QueryTree2"),
			beego.NSRouter("/treeInformation/queryTree3", &elePowDR.TreeController{},"get:QueryTree3"),


			beego.NSRouter("/treeInformation/trace", &elePowDR.TreeController{},"get:Trace"),
			beego.NSRouter("/treeInformation/edit", &elePowDR.TreeController{},"get:Edit"),
			beego.NSRouter("/treeInformation/onChain1", &elePowDR.TreeController{},"post:OnChain1"),
			beego.NSRouter("/treeInformation/plan", &elePowDR.TreeController{},"get:Plan"),
			beego.NSRouter("/treeInformation/onJudge1", &elePowDR.TreeController{},"get:OnJudge1"),
			beego.NSRouter("/treeInformation/onJudge2", &elePowDR.TreeController{},"get:OnJudge2"),
			beego.NSRouter("/treeInformation/onJudge3", &elePowDR.TreeController{},"get:OnJudge3"),
			beego.NSRouter("/treeInformation/isOpen1", &elePowDR.TreeController{},"get:IsOpen1"),
			beego.NSRouter("/treeInformation/isOpen2", &elePowDR.TreeController{},"get:IsOpen2"),
			beego.NSRouter("/treeInformation/isOpen3", &elePowDR.TreeController{},"get:IsOpen3"),
			beego.NSRouter("/treeInformation/history1", &elePowDR.TreeController{},"get:History1"),
			beego.NSRouter("/treeInformation/history2", &elePowDR.TreeController{},"get:History2"),
			beego.NSRouter("/treeInformation/history3", &elePowDR.TreeController{},"get:History3"),
			beego.NSRouter("/treeInformation/add", &elePowDR.TreeController{},"get:Add"),
	        beego.NSRouter("/treeInformation/doAdd", &elePowDR.TreeController{},"post:DoAdd"),

			beego.NSRouter("/treeInformation/cheat", &elePowDR.TreeController{},"get:Trace"),

		)
	//注册 namespace
	beego.AddNamespace(ns)
}
