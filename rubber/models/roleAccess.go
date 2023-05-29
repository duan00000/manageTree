/**
  @Author ZXQ-Administrator
  @Date 2021-08-28
  @node:
**/
package models

type RoleAccess struct {
	AccessId int
	RoleId   int
}
func (RoleAccess) TableName()string{
	return "role_access"
}
