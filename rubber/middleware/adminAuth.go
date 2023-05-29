/**
  @Author ZXQ-Administrator
  @Date 2021-09-01
  @node:判断用户有没有权限访问地址，防止通过手动输入地址进入到不同权限的地址
**/
package middleware

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
	"strings"
)


//context注意引入的包： github.com/astaxie/beego/context
func AdminAuth(ctx *context.Context){
	pathname := ctx.Request.URL.String()
	userinfo, ok := ctx.Input.Session("userInfo").(models.Manager) //类型断言
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" && pathname != "/"+beego.AppConfig.String("adminPath")+"/login/doLogin" {
			ctx.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	}else{
		pathname=strings.Replace(pathname,"/"+beego.AppConfig.String("adminPath"),"",1)
		urlPath,_:=url.Parse(pathname)
		//判断管理员是不是超级管理员以及判断排除的url地址

		//	判断管理员是不是超级管理员
		if userinfo.IsSuper ==0 && !excludeAuthPath(string(urlPath.Path)){
			// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
			roleId:=userinfo.RoleId
			roleAccess :=[]models.RoleAccess{}
			models.DB.Where("role_id=?",roleId).Find(&roleAccess)
			roleAccessMap:=make(map[int]int)
			for _,v:=range roleAccess{
				roleAccessMap[v.AccessId]=v.AccessId
			}
			// 2、获取当前访问的url对应的权限id
			/*
				/beego_admin/manager 替换成 /manager
				beego_admin/role/edit?id=11   替换成  /role/edit

				pathname = strings.Replace(pathname, "/"+beego.AppConfig.String("adminPath"), "", 1)
				urlPath, _ := url.Parse(pathname) //urlPath.Path  /role/edit
			*/
			access := models.Access{}
			models.DB.Where("url=?", urlPath.Path).Find(&access)

			//3、判断当前访问的url对应的权限id 是否在权限列表的id中
			if _, ok := roleAccessMap[access.Id]; !ok {
				ctx.Redirect(302,"/"+beego.AppConfig.String("adminPath")+"/middleware")
				//ctx.WriteString("您好，您没有该权限")
				return
			}

		}
	}
}
//判断一个url是否在排除的地址里面
func excludeAuthPath(urlPath string) bool{
	excludeAuthPathSlice :=strings.Split(beego.AppConfig.String("excludeAuthPath"),",")
	for _,v:=range excludeAuthPathSlice{
		if v == urlPath {
			return true
		}	}
	return false
}
