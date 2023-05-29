/**
  @Author ZXQ-Administrator
  @Date 2021-09-23
  @node: 注册用户 可根据自身需求添加更多属性
**/
package models



type User struct {
	Id       int
	Username string
	Phone    string
	Password string
	AddTime  int
	LastIp   string
	Email    string
	Status   int
	Craft int
	West int
	East int
	Mid int
	Trace int
}

func (User) TableName() string {
	return "user"
}

