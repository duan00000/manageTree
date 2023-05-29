/**
  @Author ZXQ-Administrator
  @Date 2021-09-02
  @node:
**/
package models

type Focus struct {
	Id        int
	Title     string
	FocusType int
	FocusImg  string
	Link      string
	Sort      int
	Status    int
	AddTime   int
}

func (Focus) TableName() string {
	return "focus"
}
