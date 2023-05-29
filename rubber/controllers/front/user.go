/**
  @Author ZXQ-Administrator
  @Date 2021-09-20
  @node:
**/
package front

import (
	"elenewenergy/models"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.SuperInit()
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	models.DB.Where("iphone",user.Phone).Find(&user)
	c.Data["userInfo"]=user
	c.TplName = "front/user/welcome.html"
}
func (c *UserController) Edit() {
	c.SuperInit()
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	models.DB.Where("iphone",user.Phone).Find(&user)
	c.Data["userInfo"]=user
	c.TplName = "front/user/user_info.html"
}
func (c *UserController) OrderList() {
	c.SuperInit()
	//1、获取当前用户
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)

	//2、获取当前用户下面的订单信息 并 分页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每页显示数量
	pageSize := 2

	//获取搜索关键词
	where := "uid=?"
	keywords := c.GetString("keywords")
	if keywords != "" {
		orderItem := []models.OrderItem{}
		models.DB.Where("product_title like ?", "%"+keywords+"%").Find(&orderItem)
		// 拼接
		var str string
		for i := 0; i < len(orderItem); i++ {
			if i == 0 {
				str += strconv.Itoa(orderItem[i].OrderId)
			} else {
				str += "," + strconv.Itoa(orderItem[i].OrderId)
			}
		}
		// 2,3,4
		// in 查询
		where += " And id in (" + str + ")"

	}
	//获取筛选条件
	orderStatus, statusError := c.GetInt("order_status")
	if statusError == nil {
		where += " And order_status=" + strconv.Itoa(orderStatus)
		c.Data["orderStatus"] = orderStatus
	} else {
		c.Data["orderStatus"] = "nil"
	}

	//总数量
	var count int
	models.DB.Where(where, user.Id).Table("order").Count(&count)

	order := []models.Order{}
	// 关联查询 分页 排序
	models.DB.Where(where, user.Id).Offset((page - 1) * pageSize).Limit(pageSize).Preload("OrderItem").Order("add_time desc").Find(&order)

	c.Data["order"] = order
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["keywords"] = keywords
	c.Data["pageSize"] = pageSize

	c.TplName = "front/user/order.html"
}
func (c *UserController) OrderInfo() {
	c.SuperInit()
	id, _ := c.GetInt("id")
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	order := models.Order{}
	models.DB.Where("id=? And uid=?", id, user.Id).Preload("OrderItem").Find(&order)
	c.Data["order"] = order
	if order.OrderId == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法请求用户详情，跳转到首页",
		}
		c.ServeJSON()
		c.Redirect("/", 302)
	}
	c.TplName = "front/user/order_info.html"
}
func (c *UserController) ChangeUserPhone() {
	phone := strings.Trim(c.GetString("phone"), "")
	phoneExit := []models.User{}
	models.DB.Where("phone=?", phone).Find(&phoneExit)
	if len(phoneExit) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号已存在",
		}
		c.ServeJSON()
		return
	}else{
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"msg":     "手机号可以使用",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) ChangeUserInfo() {
	fmt.Printf("--------------")

	id, _ := c.GetInt("id")
	fmt.Printf("id--------------",id)

	phone := strings.Trim(c.GetString("phone"), "")
	username := strings.Trim(c.GetString("username"), "")
	password := strings.Trim(c.GetString("password"), "")
	fmt.Printf("id--------------",id)
	userInfo := models.User{Id: id}
	models.DB.Find(&userInfo)

	if password == "" {
		userInfo.Phone = phone
		userInfo.Username = username
		userInfo.AddTime = int(models.GetUnix())


	}else{
		userInfo.Phone = phone
		userInfo.Username = username
		userInfo.Password = models.Md5(password)
		userInfo.AddTime = int(models.GetUnix())
	}

	err := models.DB.Save(&userInfo).Error
	if err != nil{
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "修改信息失败！",
		}
		c.ServeJSON()
		return
	}else{
		c.Data["json"] = map[string]interface{}{
			"success": true ,
			"msg":     "修改信息成功！",
		}
		c.ServeJSON()
		return
	}

}

