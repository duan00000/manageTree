/**
  @Author ZXQ-Administrator
  @Date 2021-09-21
  @node: 公共抽离
**/
package front

import (
	"elenewenergy/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"net/url"
	"strings"
)

type BaseController struct {
	beego.Controller
}
// 定义公共方法
func (c *BaseController) SuperInit() {
	//	获取顶部导航 以sort进行降序排列
	var topNav []models.Nav
	if hasTopNav := models.CacheDb.Get("topNav", &topNav); hasTopNav == true {
		c.Data["topNavList"] = topNav
	} else {
		models.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topNav)
		c.Data["topNavList"] = topNav
		models.CacheDb.Set("topNav", topNav)
	}

	//	获取左侧分类   https://gorm.io/zh_CN/docs/preload.html 自定义预加载 SQL 排序
	var goodsCate []models.GoodsCate
	if hasGoodsCate := models.CacheDb.Get("goodsCate", &goodsCate); hasGoodsCate == true {
		c.Data["goodsCateList"] = goodsCate
	} else {
		models.DB.Preload("GoodsCateItem", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Where("pid=0 AND status=1").Order("sort desc", true).Find(&goodsCate)
		c.Data["goodsCateList"] = goodsCate
		models.CacheDb.Set("goodsCate", goodsCate)
	}

	//	获取中间导航 以sort进行降序排列
	//获取中间导航的数据
	var middleNav []models.Nav
	if hasMiddleNav := models.CacheDb.Get("middleNav", &middleNav); hasMiddleNav == true {
		c.Data["middleNavList"] = middleNav
	} else {
		models.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleNav)

		for i := 0; i < len(middleNav); i++ {
			//获取关联商品
			// middleNav[i].Relation  19,20,21
			middleNav[i].Relation = strings.ReplaceAll(middleNav[i].Relation, "，", ",")
			relation := strings.Split(middleNav[i].Relation, ",")
			goods := []models.Goods{}
			models.DB.Where("id in (?)", relation).Select("id,title,goods_img,price").Find(&goods)
			middleNav[i].GoodsItem = goods
		}
		c.Data["middleNavList"] = middleNav
		models.CacheDb.Set("middleNav", middleNav)
	}

	//判断用户是否登录
	user := models.User{}
	// 获取cookie中，用户信息
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	if len(user.Phone) == 11 {
		str := fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>		

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="#">喜欢</a></li>

					<li><a href="/pass/loginOut">退出登录</a></li>
				</ol>
			
			</li>
		</ul> `, user.Phone)
		c.Data["userinfo"] = str
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/pass/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/pass/registerStep1" target="_blank" >注册</a></li>							
		</ul>`)
		c.Data["userinfo"] = str
	}
	// 个人中心 左侧名
	urlPath, _ := url.Parse(c.Ctx.Request.URL.String())
	c.Data["pathname"] = urlPath.Path
}

func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	if strings.Contains(redirect, beego.AppConfig.String("adminPath")) {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/error.html"

}
func (c *BaseController) Success(message string, redirect string) {
	//http://localhost:8080/state_grid/goods?page=2
	//判断传入地址有没有 adminPath 地址
	c.Data["message"] = message
	if strings.Contains(redirect, beego.AppConfig.String("adminPath")) {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/success.html"
}
