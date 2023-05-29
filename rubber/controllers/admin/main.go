/**
  @Author ZXQ-Administrator
  @Date 2021-08-17
  @node:
**/
package admin

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	userinfo, ok := c.GetSession("userInfo").(models.Manager)
	if ok {
		//获取用户名
		c.Data["username"] = userinfo.Username
		//1.获取角色id
		roleId := userinfo.RoleId
		//2.获取全部的权限（排序） 参考：：https://gorm.io/zh_CN/docs/preload.html
		//	//仅搜索status=1的 显示数据
		access := []models.Access{}
		models.DB.Preload("AccessItem", "status=1",func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort desc").Where("module_id=? AND status=?", 0,1).Find(&access)
		//3.获取当前角色拥有的权限，并把权限id放在一个map对象中
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", roleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
		for i := 0; i < len(access); i++ {
			if _, ok := roleAccessMap[access[i].Id]; ok {
				access[i].Checked = true
			}
			for j := 0; j < len(access[i].AccessItem); j++ {
				if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
					access[i].AccessItem[j].Checked = true
				}
			}
		}
		//5、渲染权限数据以及角色 Id
		c.Data["accessList"] = access
		// 管理员
		c.Data["isSuper"] = userinfo.IsSuper
	}
	c.TplName = "admin/main/index.html"
}
func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}

//用户权限不足时 跳转界面
func (c *MainController) Middleware() {
	c.TplName = "admin/main/middleware.html"
}

//修改公共状态
func (c *MainController) ChangeStatus() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法请求",
		}
		c.ServeJSON()
		return
	}
	table := c.GetString("table")
	field := c.GetString("field")
	//eg.	models.DB.Exec("update focus set status=ABS(status-1) where id =10)
	errInfo := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if errInfo != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "更新数据失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "更新数据成功",
	}
	c.ServeJSON()
}

//直接点击后 修改数量的方法
func (c *MainController) EditNum() {
	id := c.GetString("id")
	table := c.GetString("table")
	field := c.GetString("field")
	num := c.GetString("num")

	err := models.DB.Exec("update " + table + " set " + field + " = " + num + " where id =" + id).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "修改数量失败",
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "修改数量成功",
		}
		c.ServeJSON()
	}
}
