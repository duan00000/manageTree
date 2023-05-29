/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node: 前端 路由
**/
package routers

import (
	"elenewenergy/controllers/front"
	"elenewenergy/middleware"
	"github.com/astaxie/beego"
)

func init() {
	// product 产品
	beego.Router("/", &front.IndexController{})
	beego.Router("/login", &front.LoginController{})
	beego.Router("/category_:id([0-9]+).html", &front.ProductController{}, "get:CategoryList")
	beego.Router("/item_:id([0-9]+).html", &front.ProductController{}, "get:ProductItem")
	beego.Router("/product/getImgList", &front.ProductController{}, "get:GetImgList")

	// cart 购物
	beego.Router("/cart", &front.CartController{})
	beego.Router("/cart/addCart", &front.CartController{}, "get:AddCart")
	beego.Router("/cart/decCart", &front.CartController{}, "get:DecCart")
	beego.Router("/cart/incCart", &front.CartController{}, "get:IncCart")
	beego.Router("/cart/changeOneCart", &front.CartController{}, "get:ChangeOneCart")
	beego.Router("/cart/changeAllCart", &front.CartController{}, "get:ChangeAllCart")
	beego.Router("/cart/delCart", &front.CartController{}, "get:DelCart")
	// pass 用户 注册
	beego.Router("/pass/login", &front.PassController{}, "get:Login")
	beego.Router("/pass/doLogin", &front.PassController{}, "post:DoLogin")
	beego.Router("/pass/loginOut", &front.PassController{}, "get:LoginOut")

	beego.Router("/pass/registerStep1", &front.PassController{}, "get:RegisterStep1")
	beego.Router("/pass/registerStep2", &front.PassController{}, "get:RegisterStep2")
	beego.Router("/pass/registerStep3", &front.PassController{}, "get:RegisterStep3")

	// 注册验证码
	beego.Router("/pass/sendCode", &front.PassController{}, "get:SendCode")
	beego.Router("/pass/validateSmsCode", &front.PassController{}, "get:ValidateSmsCode")
	beego.Router("/pass/validateUsername", &front.PassController{}, "get:ValidateUsername")
	beego.Router("/pass/doRegister", &front.PassController{}, "post:DoRegister")

	//	配置中间件判断权限
	beego.InsertFilter("/buy/*", beego.BeforeRouter, middleware.FrontAuth)
	//	结算
	beego.Router("/buy/checkout", &front.BuyController{}, "get:Checkout")
	beego.Router("/buy/doOrder", &front.BuyController{}, "post:DoOrder")
	beego.Router("/buy/confirm", &front.BuyController{}, "get:Confirm")

	//	配置中间件判断权限
	beego.InsertFilter("/address/*", beego.BeforeRouter, middleware.FrontAuth)
	//	收获地址
	beego.Router("/address/addAddress", &front.AddressController{}, "post:AddAddress")
	beego.Router("/address/getOneAddressList", &front.AddressController{}, "get:GetOneAddressList")
	beego.Router("/address/doEditAddressList", &front.AddressController{}, "post:DoEditAddressList")
	beego.Router("/address/changeDefaultAddress", &front.AddressController{}, "get:ChangeDefaultAddress")

	//配置中间件判断权限
	beego.InsertFilter("/user/*", beego.BeforeRouter, middleware.FrontAuth)
	beego.Router("/user", &front.UserController{})
	beego.Router("/user/edit", &front.UserController{},"get:Edit")
	beego.Router("/user/changeUserPhone", &front.UserController{}, "get:ChangeUserPhone")
	beego.Router("/user/changeUserInfo", &front.UserController{}, "post:ChangeUserInfo")

	beego.Router("/user/order", &front.UserController{}, "get:OrderList")
	beego.Router("/user/orderinfo", &front.UserController{}, "get:OrderInfo")
	//需求响应
	beego.Router("/need", &front.NeedController{})
	beego.Router("/need/doEdit", &front.NeedController{}, `post:DoEdit`)
	//需求页面
	beego.Router("/basic", &front.BasicController{})

}
