/**
  @Author ZXQ-Administrator
  @Date 2021-10-26
  @node:减速机数据 测试联调前后端
**/
package models

type Reducer struct {
	Sno         int     `json:"id"`
	AddTime int     `json:"strRecTime"`
	Cha        float64 `json:"shape_factor"`
	Chb        float64 `json:"skw"`
	Chc        float64 `json:"crest"`
	Chd        float64 `json:"kur"`
}

func (Reducer) TableName() string {
	return "reducer"
}
