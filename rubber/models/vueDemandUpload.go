/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:
**/
package models

type VueDemandUpload struct {
	Id          int    `json:"id"`
	Uid         int    `json:"uid"` //总用户id
	Sid         int    `json:"sid"` //子用户id
	Did         int    `json:"did"` //设备id
	BL          string `json:"businessLicence"`
	CompanyInfo string `json:"companyInfo"`
	haveUp      int    `json:"uploadContract"`
	AddTime     int    `json:"addTime"`
}

func (VueDemandUpload) TableName() string {
	return "vue_demand_upload"
}
