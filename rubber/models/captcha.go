/**
  @Author ZXQ-Administrator
  @Date 2021-09-23
  @node: 验证码 生成
**/
package models

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var Cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data 验证码
	store := cache.NewMemoryCache()
	Cpt = captcha.NewWithFilter("/captcha/", store)
	Cpt.ChallengeNums = 4  //生成验证码个数
	//设置验证码高、宽
	Cpt.StdWidth = 100
	Cpt.StdHeight = 60
	// 验证码失效时间
	//cpt.Expiration= 3 * time.Minute
}
