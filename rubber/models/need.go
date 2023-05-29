/**
  @Author ZXQ-Administrator
  @Date 2021-09-15
  @node:
**/
package models

type Need struct {
	Id                int    `form:"id"`               //自增长 主键
	BasicName         string `form:"basic_name"`       //网站标题
	BasicContract     string `form:"basic_contract"`        //网站logo
	BasicFirst        string `form:"basic_first"`    //网站关键词
	BasicVolume       int `form:"basic_volume"` //网站描述
	BasicAddress      string `form:"basic_address"`       //图片加载不出来时，显示的图片


}

func (Need) TableName() string {
	return "need"
}
