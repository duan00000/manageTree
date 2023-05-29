/**
  @Author ZXQ-Administrator
  @Date 2021-10-07
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"math"
	"strconv"
)

type FrontUserPropertyController struct {
	BaseController
}

func (c *FrontUserPropertyController) Get() {
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
	frontUserInfoList := []models.FrontUserInfo{}
	//propertyList := []models.FrontUserProperty{}
	//select * from tableName limit ((page-1)*pageSize),pageSize
	//models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&propertyList)
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Table("user").Select("user.id as id,user.username as username, frontuserproperty.craft as craft, user.phone as phone,frontuserproperty.west as west,frontuserproperty.mid as mid,frontuserproperty.east as east,frontuserproperty.trace as trace,user.status as status,frontuserproperty.add_time as add_time").Joins("left join frontuserproperty on  frontuserproperty.id=user.id ").Scan(&frontUserInfoList)
	//查询goods表里面的数量
	var count int
	//统计数量赋值给count
	models.DB.Where(where).Table("frontuserproperty").Count(&count)
	//判断是否有数据，没有数据跳转到上一页
	if len(frontUserInfoList) == 0  && count != 0 {
		prvPageBack := page - 1
		if prvPageBack == 0 {
			prvPageBack = 1
		}
		c.Goto("/frontUserProperty?page=" + strconv.Itoa(prvPageBack))
	} else {
		//	信息为空
	}
	c.Data["propertyList"] = frontUserInfoList

	c.Data["totalCount"] = count
	//总数据向上取整（9.8 即取为10）除每页数量
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["pageSize"] = pageSize
	c.Data["keyword"] = keyword
	c.TplName = "admin/frontUserProperty/index.html"

}

func (c *FrontUserPropertyController) Delete() {
	id,err :=c.GetInt("id")
	if err!=nil{
		c.Error("传入参数错误", "/frontUserProperty")
		return
	}

	frontUser := models.User{Id: id}
	models.DB.Delete(&frontUser)

	frontuserproperty := models.FrontUserProperty{Id: id}
	models.DB.Delete(&frontuserproperty)

	c.Success("删除用户成功", "/frontUserProperty")
}


