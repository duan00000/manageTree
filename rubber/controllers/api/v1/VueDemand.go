/**
  @Author ZXQ-Administrator
  @Date 2022-01-13
  @node:vue数据
**/
package v1

import (
	"elenewenergy/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type TempBasicInfo struct {
	CompanyId     int    `json:"company_id"` //与vue_company 中的id绑定
	FirLiaison    string `json:"firLiaison"`
	FirLiaisonWay string `json:"firLiaisonWay"`
	SecLiaison    string `json:"secLiaison"`
	SecLiaisonWay string `json:"secLiaisonWay"`
	ContractCap   string `json:"contractCap"`
	RunningCap    string `json:"runningCap"`
	LinkAddress   string `json:"linkAddress"`
	TRC           string `json:"trc"`          //税务登记证
	MonIns        int    `json:"monAndIns"`    //是否接入
	BiddingPrice  string `json:"biddingPrice"` //竞价价格
	Status        int    `json:"status"`
	OrderID       int    `json:"order_id"`
	AddTime       int    `json:"add_time"`
}

type TempUserDevice struct {
	TUid         int    `json:"tUid"` //总用户号
	BeResponse   string `json:"beResponse"`
	TimeRespond1 string `json:"timeRespond1"`
	TimeRespond2 string `json:"timeRespond2"`
	TimeRespond3 string `json:"timeRespond3"`
	TimeRespond4 string `json:"timeRespond4"`
}
type TempRespCap struct {
	CanResponse int    `json:"can_response"`
	MTFz        string `json:"mTFz"`
	MTFm        string `json:"mTFm"`
	MTFw        string `json:"mTFw"`
	MTFn        string `json:"mTFn"`
	WFz         string `json:"wFz"`
	WFm         string `json:"wFm"`
	WFw         string `json:"wFw"`
	WFn         string `json:"wFn"`
}

type VueDemandController struct {
	beego.Controller
}

func (c *VueDemandController) GetDemandBasic() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			//2、去数据库匹配
			demandBasic := []models.VueCompany{}
			models.DB.Where("c_uid=? AND status=?", id, 0).Find(&demandBasic)
			if demandBasic[0].TUid == 0 {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "请先去用户信息界面补全信息！",
				}
				c.ServeJSON()
				return
			} else {
				demandBasicList := models.VueCompany{
					CUid:       demandBasic[0].CUid,
					TUid:       demandBasic[0].TUid,
					UserType:   demandBasic[0].UserType,
					ContractNo: demandBasic[0].ContractNo,
					CName:      demandBasic[0].CName,
				}
				fdemandDevice := [] models.VueDemandDevice{}
				models.DB.Where("t_uid=? AND status=?", demandBasic[0].TUid, 0).Find(&fdemandDevice)
				if len(fdemandDevice) < 1 {
					demandDeviceList := models.VueDemandDevice{
						CName:        demandBasic[0].CName,
						CAddress:     demandBasic[0].CAddress,
						CPhone:       demandBasic[0].CPhone,
						TUid:         demandBasic[0].TUid,
						UserType:     demandBasic[0].UserType,
						BeResponse:   0,
						TimeRespond1: 0,
						TimeRespond2: 0,
						TimeRespond3: 0,
						TimeRespond4: 0,
					}
					c.Data["json"] = map[string]interface{}{
						"success":          true,
						"message":          "获取数据成功",
						"companyInfoList":  demandBasicList,
						"demandDeviceList": demandDeviceList,
					}
					c.ServeJSON()
					return
				} else {
					demandDeviceList := models.VueDemandDevice{TUid: demandBasic[0].TUid}
					models.DB.Where("t_uid=? AND status=?", demandBasic[0].TUid, 0).Find(&demandDeviceList)
					c.Data["json"] = map[string]interface{}{
						"success":          true,
						"message":          "获取数据成功",
						"companyInfoList":  demandBasicList,
						"demandDeviceList": demandDeviceList,
					}
					c.ServeJSON()
					return
				}

			}
		}
	}
}

func (c *VueDemandController) GetDemandBasicInfo() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			//2、去数据库匹配
			demandBasic := []models.VueDemandBasic{}
			models.DB.Where("company_id=? AND status=?", id, 0).Find(&demandBasic)
			if len(demandBasic) < 1 {
				c.Data["json"] = map[string]interface{}{
					"success": true,
					"msg":     "当前没有数据！",
					"result":  0,
				}
			} else {
				vueDemandBasicInfo := models.VueDemandBasic{
					FirLiaison:    demandBasic[0].FirLiaison,
					FirLiaisonWay: demandBasic[0].FirLiaisonWay,
					SecLiaison:    demandBasic[0].SecLiaison,
					SecLiaisonWay: demandBasic[0].SecLiaisonWay,
					ContractCap:   demandBasic[0].ContractCap,
					RunningCap:    demandBasic[0].RunningCap,
					LinkAddress:   demandBasic[0].LinkAddress,
					TRC:           demandBasic[0].TRC,
					MonIns:        demandBasic[0].MonIns,
					BiddingPrice:  demandBasic[0].BiddingPrice,
				}
				c.Data["json"] = map[string]interface{}{
					"success":            true,
					"message":            "获取数据成功",
					"vueDemandBasicInfo": vueDemandBasicInfo,
				}
				c.ServeJSON()
			}

		}
	}
}

func (c *VueDemandController) DoEditUserDevice() {
	//获取前端数据VueCompany
	data := c.Ctx.Input.RequestBody
	demandDeviceInfo := TempUserDevice{}

	err := json.Unmarshal(data, &demandDeviceInfo)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Unmarshall解析错误!",
		}
		c.ServeJSON()
		return
	}
	beResponse, _ := strconv.ParseFloat(demandDeviceInfo.BeResponse, 64)
	timeRespond1, _ := strconv.ParseFloat(demandDeviceInfo.TimeRespond1, 64)
	timeRespond2, _ := strconv.ParseFloat(demandDeviceInfo.TimeRespond2, 64)
	timeRespond3, _ := strconv.ParseFloat(demandDeviceInfo.TimeRespond3, 64)
	timeRespond4, _ := strconv.ParseFloat(demandDeviceInfo.TimeRespond4, 64)

	companyInfo := models.VueCompany{}
	models.DB.Where("t_uid=? AND status=?", demandDeviceInfo.TUid, 0).Find(&companyInfo)
	demandDeviceList := []models.VueDemandDevice{}
	models.DB.Where("t_uid=?", demandDeviceInfo.TUid).Find(&demandDeviceList)
	if len(demandDeviceList) < 1 {
		cDemandDevice := models.VueDemandDevice{
			CName:        companyInfo.CName,
			CPhone:       companyInfo.CPhone,
			CAddress:     companyInfo.CAddress,
			TUid:         demandDeviceInfo.TUid,
			UserType:     companyInfo.UserType,
			BeResponse:   beResponse,
			TimeRespond1: timeRespond1,
			TimeRespond2: timeRespond2,
			TimeRespond3: timeRespond3,
			TimeRespond4: timeRespond4,
			AddTime:      int(models.GetUnix()),
			Status:       0,
		}
		errInfo := models.DB.Create(&cDemandDevice).Error
		if errInfo != nil {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": errInfo,
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"message": "公司信息修改成功!",
			}
			c.ServeJSON()
			return
		}
	} else {
		sDemandDevice := models.VueDemandDevice{TUid: demandDeviceInfo.TUid}
		models.DB.Find(&sDemandDevice)

		sDemandDevice.BeResponse = beResponse
		sDemandDevice.TimeRespond1 = timeRespond1
		sDemandDevice.TimeRespond2 = timeRespond2
		sDemandDevice.TimeRespond3 = timeRespond3
		sDemandDevice.TimeRespond4 = timeRespond4
		errInfo := models.DB.Save(&sDemandDevice).Error
		if errInfo != nil {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": errInfo,
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"message": "信息修改成功!",
			}
			c.ServeJSON()
			return
		}
	}

}

func (c *VueDemandController) GetDsmDevice() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			vueCompany := models.VueCompany{}
			models.DB.Where("c_uid=? AND status=?", id, 0).Find(&vueCompany)

			//2、去数据库匹配
			dsmUserDevice := []models.DsmUserDevice{}
			models.DB.Where("t_uid=? AND status=?", vueCompany.TUid, 0).Find(&dsmUserDevice)
			if len(dsmUserDevice) < 1 {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "联系管理员添加设备数据！",
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success":       true,
					"message":       "设备数据返回成功！",
					"dsmUserDevice": dsmUserDevice,
				}
				c.ServeJSON()
				return
			}

		}
	}
}

func (c *VueDemandController) GetRespInfo() {
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			vueRespCap := []models.VueDemandCapability{}
			models.DB.Where("c_uid=? AND status=?", id, 0).Find(&vueRespCap)
			if len(vueRespCap) < 1 {
				c.Data["json"] = map[string]interface{}{
					"success": true,
					"msg":     "当前没有数据！",
					"result":  0,
				}
			} else {
				vueRespCapList := models.VueDemandCapability{
					MTFn: vueRespCap[0].MTFn,
					MTFz: vueRespCap[0].MTFz,
					MTFw: vueRespCap[0].MTFw,
					MTFm: vueRespCap[0].MTFm,
					WFm:  vueRespCap[0].WFm,
					WFz:  vueRespCap[0].WFz,
					WFw:  vueRespCap[0].WFw,
					WFn:  vueRespCap[0].WFn,
				}
				c.Data["json"] = map[string]interface{}{
					"success":        true,
					"message":        "获取数据成功",
					"vueRespCapList": vueRespCapList,
				}
				c.ServeJSON()
			}

		}
	}

}

func (c *VueDemandController) PutRespInfo() {
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			vueCompanyInfo := []models.VueCompany{}
			models.DB.Where("c_uid=?", id).Find(&vueCompanyInfo)

			dataBody := c.Ctx.Input.RequestBody
			//beego.Info("dataBody", string(dataBody))
			tempRespCap := TempRespCap{}
			err := json.Unmarshal(dataBody, &tempRespCap)
			//beego.Info("companyInfo", tempRespCap)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "Unmarshall解析错误!",
				}
				c.ServeJSON()
				return
			} else {
				id := claims.Uid
				MTFm, _ := strconv.ParseFloat(tempRespCap.MTFm, 64)
				MTFn, _ := strconv.ParseFloat(tempRespCap.MTFn, 64)
				MTFw, _ := strconv.ParseFloat(tempRespCap.MTFw, 64)
				MTFz, _ := strconv.ParseFloat(tempRespCap.MTFz, 64)
				WFm, _ := strconv.ParseFloat(tempRespCap.WFm, 64)
				WFn, _ := strconv.ParseFloat(tempRespCap.WFn, 64)
				WFw, _ := strconv.ParseFloat(tempRespCap.WFw, 64)
				WFz, _ := strconv.ParseFloat(tempRespCap.WFz, 64)

				vueRespCap := []models.VueDemandCapability{}
				models.DB.Where("c_uid=? AND status=?", id, 0).Find(&vueRespCap)
				if len(vueRespCap) < 1 {
					sVueRespCap := models.VueDemandCapability{
						CUid:        id,
						OrderId:     0,
						MTFz:        MTFz,
						MTFm:        MTFm,
						MTFw:        MTFw,
						MTFn:        MTFn,
						WFz:         WFz,
						WFm:         WFm,
						WFw:         WFw,
						WFn:         WFn,
						CanResponse: tempRespCap.CanResponse,
						AddTime:     int(models.GetUnix()),
						Status:      0,
					}
					err := models.DB.Create(&sVueRespCap).Error
					if err != nil {
						c.Data["json"] = map[string]interface{}{
							"success": false,
							"message": "错误信息：响应能力数据存储报错，请联系管理员！",
						}
						c.ServeJSON()
						return
					} else {
						c.Data["json"] = map[string]interface{}{
							"success": true,
							"message": "响应能力数据存储成功！",
						}
						c.ServeJSON()
					}

				} else {
					vueRespCap[0].MTFm = MTFm
					vueRespCap[0].MTFn = MTFn
					vueRespCap[0].MTFw = MTFw
					vueRespCap[0].MTFz = MTFz
					vueRespCap[0].WFm = WFm
					vueRespCap[0].WFn = WFn
					vueRespCap[0].WFw = WFw
					vueRespCap[0].WFz = WFz
					errInfo := models.DB.Save(&vueRespCap[0]).Error
					if errInfo != nil {
						c.Data["json"] = map[string]interface{}{
							"success": false,
							"message": errInfo,
						}
						c.ServeJSON()
						return
					} else {
						c.Data["json"] = map[string]interface{}{
							"success": true,
							"message": "响应能力数据修改成功!",
						}
						c.ServeJSON()
						return
					}
				}

			}

		}
	}

}

func (c *VueDemandController) PutBasicInfo() {
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid

			dataBody := c.Ctx.Input.RequestBody
			beego.Info("dataBody", string(dataBody))
			tempRespCap := TempBasicInfo{}
			err := json.Unmarshal(dataBody, &tempRespCap)
			beego.Info("companyInfo", tempRespCap)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "Unmarshall解析错误!",
				}
				c.ServeJSON()
				return
			} else {
				contractCap, _ := strconv.ParseFloat(tempRespCap.ContractCap, 64)
				runningCap, _ := strconv.ParseFloat(tempRespCap.RunningCap, 64)
				biddingPrice, _ := strconv.ParseFloat(tempRespCap.BiddingPrice, 64)
				year := time.Now().Year()

				vueBasic := []models.VueDemandBasic{}
				models.DB.Where("company_id=? AND status=?", id, 0).Find(&vueBasic)
				if len(vueBasic) < 1 {
					sVueBasic := models.VueDemandBasic{
						CompanyId:     id,
						JoinY:         year,
						FirLiaison:    tempRespCap.FirLiaison,
						FirLiaisonWay: tempRespCap.FirLiaisonWay,
						SecLiaison:    tempRespCap.SecLiaison,
						SecLiaisonWay: tempRespCap.SecLiaisonWay,
						ContractCap:   contractCap,
						RunningCap:    runningCap,
						BiddingPrice:  biddingPrice,
						TRC:           tempRespCap.TRC,
						Status:        0,
						OrderId:       0,
						LinkAddress:   tempRespCap.LinkAddress,
						AddTime:       int(models.GetUnix()),
					}
					err := models.DB.Create(&sVueBasic).Error
					if err != nil {
						c.Data["json"] = map[string]interface{}{
							"success": false,
							"message": "错误信息：基础信息存储报错，请联系管理员！",
						}
						c.ServeJSON()
						return
					} else {
						c.Data["json"] = map[string]interface{}{
							"success": true,
							"message": "基础信息数据存储成功！",
						}
						c.ServeJSON()
					}

				} else {
					vueBasic[0].FirLiaison = tempRespCap.FirLiaison
					vueBasic[0].FirLiaisonWay = tempRespCap.FirLiaisonWay
					vueBasic[0].SecLiaison = tempRespCap.SecLiaison
					vueBasic[0].SecLiaisonWay = tempRespCap.SecLiaisonWay
					vueBasic[0].ContractCap = contractCap
					vueBasic[0].RunningCap = runningCap
					vueBasic[0].TRC = tempRespCap.TRC
					vueBasic[0].MonIns = tempRespCap.MonIns
					vueBasic[0].BiddingPrice = biddingPrice
					vueBasic[0].LinkAddress = tempRespCap.LinkAddress
					errInfo := models.DB.Save(&vueBasic[0]).Error
					if errInfo != nil {
						c.Data["json"] = map[string]interface{}{
							"success": false,
							"message": errInfo,
						}
						c.ServeJSON()
						return
					} else {
						c.Data["json"] = map[string]interface{}{
							"success": true,
							"message": "基本信息数据修改成功!",
						}
						c.ServeJSON()
						return
					}
				}

			}

		}
	}

}

func (c *VueDemandController) GetOrderDemand() {
	// 1、获取token
	tokenData := c.Ctx.Input.Header("Authorization")
	tokenString := strings.Split(tokenData, " ")[1]
	//2、验证token
	if tokenString == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "token传入错误",
		}
		c.ServeJSON()
		return

	} else {
		token, claims, err := ParseToken(tokenString)
		beego.Info("claims", claims)
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {
			id := claims.Uid
			//2、去数据库匹配
			orderDemand := []models.OrderDemand{}
			models.DB.Where("company_id=? AND status=?", id, 0).Find(&orderDemand)

			if len(orderDemand) < 1 {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "不存在申报的流程！",
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success":         true,
					"message":         "流程数据返回成功！",
					"orderDemandList": orderDemand,
				}
				c.ServeJSON()
				return
			}
		}
	}
}

func (c *VueDemandController) AbandonOrder() {
	//tokenData := c.Ctx.Input.Header("Authorization")
	data := c.Ctx.Input.RequestBody
	beego.Info("data", string(data))

	orderDemand := models.OrderDemand{}
	err := json.Unmarshal(data, &orderDemand)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Unmarshall解析错误!",
		}
		c.ServeJSON()
		return
	}
	orderId := orderDemand.OrderId
	errInfo := models.DB.Where("order_id=?", orderId).Find(&orderDemand).Error
	if errInfo != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "数据库没有该订单，请联系管理员!",
		}
		c.ServeJSON()
		return
	}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderDemand)

	orderBasic := models.OrderDemandBasic{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderBasic)

	orderCap := models.OrderDemandCapability{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderCap)

	orderDevice := models.OrderDemandDevice{}
	models.DB.Where("order_Id = ?", orderId).Delete(&orderDevice)

	respM := models.RespondManager{}
	models.DB.Where("order_Id = ?", orderId).Delete(&respM)
	c.Data["json"] = map[string]interface{}{
		"success":     true,
		"message":     "订单删除成功！",
		"orderDemand": orderDemand,
	}
	c.ServeJSON()
	return
}
