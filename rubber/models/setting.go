/**
  @Author ZXQ-Administrator
  @Date 2021-09-15
  @node:
**/
package models

type Setting struct {
	Id              int    `form:"id"`               //自增长 主键
	SiteTitle       string `form:"site_title"`       //网站标题
	SiteLogo        string `form:"site_logo"`        //网站logo
	SiteKeywords    string `form:"site_keywords"`    //网站关键词
	SiteDescription string `form:"site_description"` //网站描述
	NoPicture       string `form:"no_picture"`       //图片加载不出来时，显示的图片
	SiteIcp         string `form:"site_icp"`         //备案信息
	SiteTel         string `form:"site_tel"`         //客服电话
	SearchKeywords  string `form:"search_keywords"`  //搜索关键词
	TongjiCode      string `form:"tongji_code"`      //统计（访问）的代码
	Appid           string `form:"appid"`            //云存储
	AppSecret       string `form:"app_secret"`       //
	EndPoint        string `form:"end_point"`        //
	BucketName      string `form:"bucket_name"`      //
	OssStatus       int    `form:"oss_status"`

}

func (Setting) TableName() string {
	return "setting"
}
