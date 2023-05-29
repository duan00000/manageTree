/**
  @Author ZXQ-Administrator
  @Date 2022-01-10
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

type VuePassController struct {
	beego.Controller
}
type RegSms struct {
	Phone string `json:"phone"`
	SmsCode   string `json:"smsCode"`
}
type RegInfo struct {
	Company         string `json:"company"`
	Mobile          string `json:"mobile"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	ConfirmPassword string `json:"confirmPassword"`
	VueSign         string `json:"vueSign"`
	VueSmsCode      string `json:"verificationCode"`
}

func (c *VuePassController) SendCode() {
	// 自定义数据类型
	phone := c.Input().Get("phone")

	if phone == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号为空",
		}
		c.ServeJSON()
		return
	}
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号的格式不正确",
		}
		c.ServeJSON()
		return
	}
	//3、验证手机号是否注册过
	vueUserT := []models.VueUser{}
	models.DB.Where("phone=?", phone).Find(&vueUserT)
	if len(vueUserT) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此用户已经存在",
		}
		c.ServeJSON()
		return
	}
	//4、验证手机注册的次数
	addDay := models.GetDay()                             //年月日
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0] //ip 127.0.0.1:64912

	sign := models.Md5(phone + addDay)  //签名
	vueSmsCode := models.GetRandomNum() //发送短信的随机码
	beego.Info("VueSmsCode............", vueSmsCode)
	userTemp := []models.VueUserTemp{}
	models.DB.Where("add_day=? AND phone=?", addDay, phone).Find(&userTemp)
	// 隔天 刷新SendCount 记录
	if len(userTemp) > 0 {
		userTempNew := models.VueUserTemp{}
		models.DB.Where("add_day=? AND phone=?", addDay, phone).Find(&userTempNew)
		exitDay := userTemp[0].AddDay
		if addDay != exitDay {
			userTempNew.SendCount = 0
			models.DB.Save(&userTempNew)
			userTemp[0].SendCount = 0
		}
	}
	var sendCount int
	models.DB.Where("add_day=? AND ip=?", addDay, ip).Table("vue_user_temp").Count(&sendCount)

	//当前ip地址今天发送的次数
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				//1、发送验证码
				models.VueSendMsg(vueSmsCode)
				//2、保存验证码
				c.SetSession("vue_sms_code", vueSmsCode)
				//3、更新发送短信信息
				oneUserTemp := models.VueUserTemp{}
				models.DB.Where("id=?", userTemp[0].Id).Find(&oneUserTemp)
				oneUserTemp.SendCount = oneUserTemp.SendCount + 1
				oneUserTemp.SmsCode = vueSmsCode
				models.DB.Save(&oneUserTemp)

				c.Data["json"] = map[string]interface{}{
					"success": true,
					"msg":     "短信发送成功",
					"sign":    sign,
					// "sms_code": sms_code, //测试期间返回数据
				}
				c.ServeJSON()
				return

			} else {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"msg":     "当前手机号今天发送短信的次数太多了，明天再试",
				}
				c.ServeJSON()
				return
			}

		} else {
			//1、发送验证码
			models.VueSendMsg(vueSmsCode)
			//2、保存验证码
			c.SetSession("vue_sms_code", vueSmsCode)
			//3、保存发送的信息
			oneUserTemp := models.VueUserTemp{
				Ip:        ip,
				Phone:     phone,
				SendCount: 1,
				AddDay:    addDay,
				AddTime:   int(models.GetUnix()),
				Sign:      sign,
				SmsCode:   vueSmsCode,
			}
			errCode := models.DB.Create(&oneUserTemp).Error
			if errCode != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"msg":     "错误信息：数据库存储错误，请联系管理员！",
				}
				c.ServeJSON()
				return
			}
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"msg":     "短信发送成功",
				"sign":    sign,
				// "sms_code": sms_code, //测试期间返回数据
			}
			c.ServeJSON()
			return

		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此ip今天发送次数已经达到上限，明天再试",
		}
		c.ServeJSON()
		return
	}
}

// 注册 验证验证码
func (c *VuePassController) RegisterSms() {
	RegSmsUser := RegSms{}
	//获取前端数据
	data := c.Ctx.Input.RequestBody
	beego.Info("data", string(data))

	err := json.Unmarshal(data, &RegSmsUser)
	beego.Info("RegSmsUser", RegSmsUser)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "错误信息：Unmarshal解析失败!",
		}
		c.ServeJSON()
		return
	}

	phoneExit := []models.VueUser{}
	models.DB.Where("phone=?", RegSmsUser.Phone).Find(&phoneExit)
	if len(phoneExit) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号已存在",
		}
		c.ServeJSON()
		return
	}

	//1、验证数据是否合法 验证签名,验证短信验证码是否正确 与数据库里比较
	userTemp := []models.VueUserTemp{}
	models.DB.Where("phone=? AND sms_code=?", RegSmsUser.Phone, RegSmsUser.SmsCode).Find(&userTemp)
	if len(userTemp) == 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "错误信息：smsCode验证错误！",
		}
		c.ServeJSON()
		return
	}
	//2、判断验证码有没有过期   时间限制15分
	nowTime := models.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "当前验证码已经过期！",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "手机号和验证码验证通过！",
	}
	c.ServeJSON()
	return
}

// 注册
func (c *VuePassController) RegisterStep() {
	// 自定义数据类型
	regInfo := RegInfo{}
	//获取前端数据
	data := c.Ctx.Input.RequestBody
	//beego.Info("data", string(data))

	err := json.Unmarshal(data, &regInfo)
	//beego.Info("vueUser", regInfo)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "错误信息：Unmarshal解析失败!",
		}
		c.ServeJSON()
		return
	}
	//前端传入数据
	phone := regInfo.Mobile
	vueSign := regInfo.VueSign
	vueSmsCode := regInfo.VueSmsCode

	phoneExit := []models.VueUser{}
	models.DB.Where("phone=?", phone).Find(&phoneExit)
	if len(phoneExit) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号已存在",
		}
		c.ServeJSON()
		return
	}

	//1、验证数据是否合法 验证签名,验证短信验证码是否正确 与数据库里比较
	userTemp := []models.VueUserTemp{}
	models.DB.Where("sign=? AND sms_code=?", vueSign, vueSmsCode).Find(&userTemp)
	if len(userTemp) == 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "错误信息：sign或smsCode验证错误",
		}
		c.ServeJSON()
		return
	}
	//2、判断验证码有没有过期   时间限制15分
	nowTime := models.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "当前验证码已经过期",
		}
		c.ServeJSON()
		return
	}
	if regInfo.Password != regInfo.ConfirmPassword {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "错误信息：密码和确认密码不一致",
		}
		c.ServeJSON()
		return
	} else {
		ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0] //获取客户端ip
		vueUser := models.VueUser{
			Username:    regInfo.Username,
			CompanyName: regInfo.Company,
			Phone:       regInfo.Mobile,
			Password:    models.Md5(regInfo.Password),
			AddTime:     int(models.GetUnix()),
			LastIp:      ip,
			Email:       regInfo.Email,
			Status:      0,
		}
		err := models.DB.Create(&vueUser).Error
		if err != nil {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"msg":     "错误信息：数据库存储错误，请联系管理员！",
			}
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"msg":     "用户创建成功，请登录",
		}
		c.ServeJSON()
		return
	}

}
