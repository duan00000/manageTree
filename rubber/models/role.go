/**
  @Author ZXQ-Administrator
  @Date 2021-08-20
  @node:
**/
package models

type Role struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime    int
}

func (Role) TableName() string {
	return "role"
}
