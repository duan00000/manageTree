
//  @Author ZXQ-Administrator
//  @Date 2022-01-01
//  @node:使用JWT实现接口的安全验证 https://github.com/dgrijalva/jwt-go
//
//**/
package v1
//
//import (
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/dgrijalva/jwt-go"
//	"strings"
//	"time"
//)
//
//type TokenController struct {
//	beego.Controller
//}
//
//// 创建一个结构体存储生成token的信息
//type MyClaims struct {
//	Uid int
//	jwt.StandardClaims
//}
//
////签名字符串 自己定义
//var jwtKey = []byte("this is sign")
//
////过期时间 当前时间+24小时 返回Unix
//var expireTime = time.Now().Add(24 * time.Hour).Unix()
//
//func (c *TokenController) Get() {
//	c.Ctx.WriteString("API首页")
//}
//func (c *TokenController) Login() {
//
//	//1、实例化自定义的结构体
//	myClaimsObj := &MyClaims{
//		Uid: 50, //生成token的时候传值
//		StandardClaims: jwt.StandardClaims{
//			// 必须定义 过期时间
//			ExpiresAt: expireTime,
//			// 签发人
//			Issuer: "zxq",
//		},
//	}
//	//2、使用指定的签名方法创建签名对象 (签名算法加密方式，实例化结构体对象)
//	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaimsObj)
//	//3、使用指定的secret签名并获得完整的编码后的字符串token
//	tokenStr, err := tokenObj.SignedString(jwtKey)
//	if err != nil {
//		fmt.Println(err)
//		c.Data["json"] = map[string]interface{}{
//			"success": false,
//			"token":   "",
//			"message": "",
//		}
//		c.ServeJSON()
//		return
//	}
//	// c.Ctx.WriteString(tokenStr)
//	c.Data["json"] = map[string]interface{}{
//		"success": true,
//		"token":   tokenStr,
//		"message": "登录成功",
//	}
//	c.ServeJSON()
//}
//func (c *TokenController) AddAddress(){
////	1、获取token
//	tokenData := c.Ctx.Input.Header("Authorization")
//	fmt.Println(tokenData)
//	tokenString := strings.Split(tokenData," ")[1]
//
////	2、验证token
//	if tokenString == ""{
//		c.Data["json"] =map[string]interface{}{
//			"success": true,
//			"message": "登录成功",
//		}
//		c.ServeJSON()
//	}else {
//		token,claims,err := ParseToken(tokenString)
//		if err != nil || !token.Valid {
//			c.Data["json"] = map[string]interface{}{
//				"success": false,
//				"message": "token传入错误",
//			}
//			c.ServeJSON()
//		} else {
//			c.Data["json"] = map[string]interface{}{
//				"success": true,
//				"message": "获取数据成功",
//				"result":  "增加地址成功",
//				"uid":     claims.Uid,
//			}
//			c.ServeJSON()
//		}
//	}
//}
//
////验证token
//func ParseToken(tokenString string)(*jwt.Token,*MyClaims,error){
//	s:=&MyClaims{}
//	token,err:=jwt.ParseWithClaims(tokenString,s, func(token *jwt.Token) (i interface{}, e error) {
//		return jwtKey,nil
//	})
//	return token,s,err
//}