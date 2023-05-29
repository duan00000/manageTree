/**
  @Author ZXQ-Administrator
  @Date 2021-10-27
  @node:
**/
package v1

import (
	"elenewenergy/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type VueLoginController struct {
	beego.Controller
}

// 创建一个结构体存储生成token的信息
type MyClaims struct {
	Uid int
	jwt.StandardClaims
}

//签名字符串
var jwtKey = []byte("this is sign")

//过期时间
var expireTime = time.Now().Add(24 * time.Hour).Unix()

func (c *VueLoginController) Get() {

}
func (c *VueLoginController) SaveIdentifyCodes() {

}
func (c *VueLoginController) DoLogin() {
	vueUser := models.VueUser{}
	//获取前端数据
	data := c.Ctx.Input.RequestBody
	//beego.Info("data：", string(data))
	err := json.Unmarshal(data, &vueUser)
	//beego.Info("vueUser：", vueUser)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"token":   "",
			"message": "解析失败!",
		}
		c.ServeJSON()
		return
	}
	//1、验证用户输入的验证码是否正确strings.EqualFold("Go", "go") 忽略大小写
	if strings.EqualFold(vueUser.InputCap, vueUser.Captcha) {
		username := vueUser.Phone
		password := models.Md5(vueUser.Password)
		//2、去数据库匹配
		vueUserInfo := []models.VueUser{}
		models.DB.Where("phone=? AND password=? AND status=?", username, password, 0).Find(&vueUserInfo)
		beego.Info("length:", vueUserInfo)

		if len(vueUserInfo) > 0 { //登录成功
			// 1、实例化自定义的结构体
			myClaimsObj := &MyClaims{
				vueUserInfo[0].Id,
				jwt.StandardClaims{
					ExpiresAt: expireTime,
					Issuer:    "zxq",
				},
			}
			//2、 使用指定的签名方法创建签名对象
			tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaimsObj)
			//3、使用指定的secret签名并获得完整的编码后的字符串token
			tokenStr, errToken := tokenObj.SignedString(jwtKey)
			if errToken != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"token":   "",
					"msg":     errToken,
				}
				c.ServeJSON()
				return
			}
			// 把用户信息保存在cookie里面
			//models.Cookie.Set(c.Ctx, "userinfo", user[0])
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"token":   tokenStr,
				"msg":     "登录成功",
			}
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"token":   "",
			"msg":     "用户名或者密码错误",
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"token":   "",
			"msg":     "输入的图形验证码不正确",
		}
		c.ServeJSON()
		return
	}
}

func (c *VueLoginController) GetProfile() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	fmt.Println("tokenData", tokenData)
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return
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
			vueUserInfo := []models.VueUser{}
			models.DB.Where("id=? AND status=?", id, 0).Find(&vueUserInfo)
			//beego.Info(vueUserInfo)

			c.Data["json"] = map[string]interface{}{
				"success": true,
				"message": "获取数据成功",
				"username":vueUserInfo[0].Username,
			}
			c.ServeJSON()
		}
	}
	//// 1、获取token
	//tokenData := c.Ctx.Input.Header("Authorization")
	//fmt.Println("tokenData",tokenData)
	//tokenString := strings.Split(tokenData, " ")[1]
	//
	////2、验证token
	//if tokenString == "" {
	//	c.Data["json"] = map[string]interface{}{
	//		"success": false,
	//		"message": "token传入错误",
	//	}
	//	c.ServeJSON()
	//} else {
	//	token, claims, err := ParseToken(tokenString)
	//	beego.Info("claims",claims)
	//	if err != nil || !token.Valid {
	//		c.Data["json"] = map[string]interface{}{
	//			"success": false,
	//			"message": "token传入错误",
	//		}
	//		c.ServeJSON()
	//	} else {
	//		c.Data["json"] = map[string]interface{}{
	//			"success": true,
	//			"message": "获取数据成功",
	//			"result":  "已经获取到了用户的信息",
	//			"uid":     claims.Uid,
	//		}
	//		c.ServeJSON()
	//	}
	//
	//}

}

// 验证token
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	s := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, s, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, s, err
}
