/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node:
**/
package front

import (
	"elenewenergy/models"
	"fmt"
	"time"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	//调用功能功能
	c.SuperInit()

	//开始时间
	startTime := time.Now().UnixNano()

	// 	获取轮播图 以sort进行降序排列
	//获取轮播图 注意获取的时候要写地址
	var focus []models.Focus
	if hasFocus := models.CacheDb.Get("focus", &focus); hasFocus == true {
		c.Data["focusList"] = focus
	} else {
		models.DB.Where("status=1 AND focus_type=1").Order("sort desc").Find(&focus)
		c.Data["focusList"] = focus
		models.CacheDb.Set("focus", focus)
	}
	models.Cookie.Set(c.Ctx,"focus",focus)

	//获取楼层数据
	//风电 "hot"
	var redisWind []models.Goods
	if hasPhone := models.CacheDb.Get("wind", &redisWind); hasPhone == true {
		c.Data["windList"] = redisWind
	} else {
		wind := models.GetGoodsByCategory(34, "", 8)
		c.Data["windList"] = wind
		models.CacheDb.Set("wind", wind)
	}

	//光电
	var redisSun []models.Goods
	if hasPhone := models.CacheDb.Get("sun", &redisSun); hasPhone == true {
		c.Data["sunList"] = redisSun
	} else {
		sun := models.GetGoodsByCategory(35, "", 8)
		c.Data["sunList"] = sun
		models.CacheDb.Set("sun", sun)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)

	c.TplName = "front/index/index.html"

}
