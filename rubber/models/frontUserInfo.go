/**
  @Author ZXQ-Administrator
  @Date 2021-10-07
  @node: 前端人员信息
**/
package models

type FrontUserInfo struct {
	Id       int
	Username string
	West int
	Phone    string
	AddTime  int
	Status   int
	Mid int
	East int
	Craft int
	Trace int
}

func (FrontUserInfo) TableName() string {
	return "frontUserInfo"
}
