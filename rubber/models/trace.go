/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node: 需求响应 订单信息
**/
package models

type Trace struct {

	Info       string  `json:"info"`

}

func (Trace) TableName() string {
	return "trace"
}
