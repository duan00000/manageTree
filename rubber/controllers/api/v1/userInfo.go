/**
  @Author ZXQ-Administrator
  @Date 2022-01-09
  @node:
**/
package v1

import (
	"elenewenergy/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"regexp"
	"strings"
)

type UserInfoController struct {
	beego.Controller
}
type UserInfo struct {
	Id          int    `json:"id"`
	CompanyDesc string `json:"companyDesc"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	AddTime     string `json:"addTime"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}
type EditUser struct {
	Id        int    `json:"id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	CheckPass string `json:"checkPass"`
}

func (c *UserInfoController) GetUserInfo() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
	} else {
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
		} else {
			id := claims.Uid
			//2、去数据库匹配
			vueUserInfo := []models.VueUser{}
			models.DB.Where("id=? AND status=?", id, 0).Find(&vueUserInfo)

			userInfo := UserInfo{
				Id:          vueUserInfo[0].Id,
				Username:    vueUserInfo[0].Username,
				CompanyDesc: vueUserInfo[0].CompanyDesc,
				Phone:       vueUserInfo[0].Phone,
				AddTime:     models.UnixToDate(vueUserInfo[0].AddTime), //时间戳转换成日期
				Email:       vueUserInfo[0].Email,
				Status:      vueUserInfo[0].Status,
			}

			c.Data["json"] = map[string]interface{}{
				"success":      true,
				"message":      "获取数据成功",
				"userInfoList": userInfo,
			}
			c.ServeJSON()
		}
	}
}

func (c *UserInfoController) PutUserInfo() {
	// 自定义数据类型
	editUser := EditUser{}
	//获取前端数据
	data := c.Ctx.Input.RequestBody
	//beego.Info("vueUser", string(data))

	err := json.Unmarshal(data, &editUser)
	//beego.Info("vueUser", editUser)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Unmarshal解析失败!",
		}
		c.ServeJSON()
		return
	} else {
		id := editUser.Id
		//2、去数据库匹配
		vueUserInfo := models.VueUser{}
		models.DB.Where("id=? AND status=?", id, 0).Find(&vueUserInfo)
		// 防止Id和phone 不匹配
		if vueUserInfo.Phone != editUser.Phone {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "错误信息：数据错误!",
			}
			c.ServeJSON()
			return
		} else {
			if editUser.Password != editUser.CheckPass {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "错误信息：密码不一致!",
				}
				c.ServeJSON()
				return
			} else {
				vueUserInfo.Password = models.Md5(editUser.Password)
				errInfo := models.DB.Save(&vueUserInfo).Error
				if errInfo != nil {
					c.Data["json"] = map[string]interface{}{
						"success": false,
						"message": "错误信息：修改密码失败!",
					}
					c.ServeJSON()
					return
				}
			}
		}
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "密码修改成功!",
		}
		c.ServeJSON()
		return
	}

}

// 公司信息
func (c *UserInfoController) GetCompanyInfo() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
		} else {
			id := claims.Uid
			//2、去数据库匹配
			vueCompanyInfo := []models.VueCompany{}
			models.DB.Where("c_uid=?", id).Find(&vueCompanyInfo)
			if len(vueCompanyInfo) < 1 {
				companyInfoList := models.VueCompany{
					CUid:     id,
					TUid:     id + int(models.GetUnix()),
					UserType: 1,
					AddTime:  int(models.GetUnix()),
					Status:   0,
				}
				err := models.DB.Create(&companyInfoList).Error
				if err != nil {
					c.Data["json"] = map[string]interface{}{
						"success": false,
						"msg":     "错误信息：公司信息存取错误，请联系管理员！",
					}
					c.ServeJSON()
					return
				} else {
					c.Data["json"] = map[string]interface{}{
						"success":         true,
						"message":         "获取数据成功",
						"companyInfoList": companyInfoList,
					}
					c.ServeJSON()
				}
			} else {
				companyInfoList := models.VueCompany{
					CKind:      vueCompanyInfo[0].CKind,
					CCount:     vueCompanyInfo[0].CCount,
					CName:      vueCompanyInfo[0].CName,
					CAddress:   vueCompanyInfo[0].CAddress,
					CPhone:     vueCompanyInfo[0].CPhone,
					CLaw:       vueCompanyInfo[0].CLaw,
					CCapital:   vueCompanyInfo[0].CCapital,
					UserType:   vueCompanyInfo[0].UserType,
					ContractNo: vueCompanyInfo[0].ContractNo,
					TUid:       vueCompanyInfo[0].TUid,
				}
				c.Data["json"] = map[string]interface{}{
					"success":         true,
					"message":         "获取数据成功",
					"companyInfoList": companyInfoList,
				}
				c.ServeJSON()
			}
		}
	}
}
func (c *UserInfoController) PutCompanyInfo() {
	//获取前端数据VueCompany
	data := c.Ctx.Input.RequestBody
	vueCompanyInfo := models.VueCompany{}
	err := json.Unmarshal(data, &vueCompanyInfo)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Unmarshall解析错误!",
		}
		c.ServeJSON()
		return
	}
	if vueCompanyInfo.CPhone == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号为空",
		}
		c.ServeJSON()
		return
	}
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(vueCompanyInfo.CPhone) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号的格式不正确",
		}
		c.ServeJSON()
		return
	}
	if vueCompanyInfo.CName == "" || vueCompanyInfo.CAddress == "" || vueCompanyInfo.TUid == 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "前端公司信息传入数据错误!",
		}
		c.ServeJSON()
		return
	} else {
		vueUserInfo := models.VueUser{}
		models.DB.Where("phone=? AND status=?", vueCompanyInfo.CPhone, 0).Find(&vueUserInfo)

		companyInfo := []models.VueCompany{}
		models.DB.Where("c_uid=? AND t_uid=?", vueUserInfo.Id, vueCompanyInfo.TUid).Find(&companyInfo)
		beego.Info("len(companyInfo):...",len(companyInfo))
		if len(companyInfo) < 1 {
			createCompany:= models.VueCompany{
				CKind:      vueCompanyInfo.CKind,
				CUid:       vueUserInfo.Id,
				TUid:       vueCompanyInfo.TUid,
				UserType:   vueCompanyInfo.UserType,
				ContractNo: vueCompanyInfo.ContractNo,
				CPhone:     vueCompanyInfo.CPhone,
				CName:      vueCompanyInfo.CName,
				CAddress:   vueCompanyInfo.CAddress,
				CLaw:       vueCompanyInfo.CLaw,
				CCount:     vueCompanyInfo.CCount,
				CCapital:   vueCompanyInfo.CCapital,
				Status:     0,
				AddTime:    int(models.GetUnix()),
			}
			errCreate := models.DB.Create(&createCompany).Error
			if errCreate != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "errCreate,添加数据失败",
				}
				c.ServeJSON()
				return
			}else {
				c.Data["json"] = map[string]interface{}{
					"success": true,
					"message": "公司信息添加成功!",
				}
				c.ServeJSON()
				return
			}
		}else{
			companyInfo[0].CKind = vueCompanyInfo.CKind
			companyInfo[0].UserType = vueCompanyInfo.UserType
			companyInfo[0].ContractNo = vueCompanyInfo.ContractNo
			companyInfo[0].CPhone = vueCompanyInfo.CPhone
			companyInfo[0].CName = vueCompanyInfo.CName
			companyInfo[0].CAddress = vueCompanyInfo.CAddress
			companyInfo[0].CLaw = vueCompanyInfo.CLaw
			companyInfo[0].CCount = vueCompanyInfo.CCount
			companyInfo[0].CCapital = vueCompanyInfo.CCapital

			errInfo := models.DB.Save(&companyInfo[0]).Error
			if errInfo != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": errInfo,
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success": true,
					"message": "公司信息修改成功!",
				}
				c.ServeJSON()
				return
			}
		}
		}
}
