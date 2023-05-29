/**
  @Author ZXQ-Administrator
  @Date 2022-01-23
  @node:
**/
package v1

import (
	"elenewenergy/models"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) SubmitAllInfo() {
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
		if err != nil || !token.Valid {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "token传入错误",
			}
			c.ServeJSON()
			return

		} else {

			//存入对应order表
			id := claims.Uid
			year := time.Now().Year()
			vueCompanyInfo := models.VueCompany{}
			models.DB.Where("c_uid=?", id).First(&vueCompanyInfo)
			orderId := id + year + int(models.GetUnix())

			// 将数据转移到新表 basic
			vueDemandBasic := models.VueDemandBasic{}
			models.DB.Where("company_id=?", id).First(&vueDemandBasic)
			//beego.Info("vueDemandBasic:", vueDemandBasic)
			orderDemandBasic := models.OrderDemandBasic{
				CompanyId:     vueDemandBasic.CompanyId,
				VueBasicId:    vueDemandBasic.Id,
				JoinY:         year,
				FirLiaison:    vueDemandBasic.FirLiaison,
				FirLiaisonWay: vueDemandBasic.FirLiaisonWay,
				SecLiaison:    vueDemandBasic.SecLiaison,
				SecLiaisonWay: vueDemandBasic.SecLiaisonWay,
				ContractCap:   vueDemandBasic.ContractCap,
				RunningCap:    vueDemandBasic.RunningCap,
				LinkAddress:   vueDemandBasic.LinkAddress,
				TRC:           vueDemandBasic.TRC,
				MonIns:        vueDemandBasic.MonIns,
				BiddingPrice:  vueDemandBasic.BiddingPrice,
				Status:        0,
				OrderId:       orderId,
				AddTime:       int(models.GetUnix()),
			}
			errBasic := models.DB.Create(&orderDemandBasic).Error
			if errBasic != nil {
				c.Data["json"] = map[string]interface{}{
					"success":    false,
					"resultCode": "errBasic",
					"message":    errBasic,
				}
				c.ServeJSON()
				return
			}

			// 将数据转移到新表 cap
			vueDemandCap := models.VueDemandCapability{}
			models.DB.Where("c_uid=?", id).First(&vueDemandCap)
			orderDemandCap := models.OrderDemandCapability{
				VueCapId:    vueDemandCap.Id,
				CUid:        id,
				MTFz:        vueDemandCap.MTFz,
				MTFm:        vueDemandCap.MTFm,
				MTFw:        vueDemandCap.MTFw,
				MTFn:        vueDemandCap.MTFn,
				WFz:         vueDemandCap.WFz,
				WFm:         vueDemandCap.WFm,
				WFw:         vueDemandCap.WFw,
				WFn:         vueDemandCap.WFn,
				CanResponse: vueDemandCap.CanResponse,
				Status:      0,
				OrderId:     orderId,
				AddTime:     int(models.GetUnix()),
			}
			errCap := models.DB.Create(&orderDemandCap).Error
			if errCap != nil {
				c.Data["json"] = map[string]interface{}{
					"success":    false,
					"resultCode": "errCap",
					"message":    errCap,
				}
				c.ServeJSON()
				return
			}
			// 将数据转移到新表 dev
			vueDemandDevice := models.VueDemandDevice{}
			models.DB.Where("t_uid=?", vueCompanyInfo.TUid).First(&vueDemandDevice)
			orderDemandDevice := models.OrderDemandDevice{
				VueDeviceId:  vueDemandDevice.Id,
				CName:        vueDemandDevice.CName,
				CAddress:     vueDemandDevice.CAddress,
				CPhone:       vueDemandDevice.CPhone,
				TUid:         vueCompanyInfo.TUid,
				UserType:     vueDemandDevice.UserType,
				BeResponse:   vueDemandDevice.BeResponse,
				SubId:        vueDemandDevice.SubId,
				SubDeviceId:  vueDemandDevice.SubDeviceId,
				DeviceId:     vueDemandDevice.DeviceId,
				TimeRespond1: vueDemandDevice.TimeRespond1,
				TimeRespond2: vueDemandDevice.TimeRespond2,
				TimeRespond3: vueDemandDevice.TimeRespond3,
				TimeRespond4: vueDemandDevice.TimeRespond4,
				Status:       0,
				OrderId:      orderId,
				AddTime:      int(models.GetUnix()),
			}
			errDevice := models.DB.Create(&orderDemandDevice).Error
			if errDevice != nil {
				c.Data["json"] = map[string]interface{}{
					"success":    false,
					"resultCode": "errDevice",
					"message":    errDevice,
				}
				c.ServeJSON()
				return
			}

			// 存入订单数据 order_demand
			//获取对应表 id
			vueBasic := models.VueDemandBasic{}
			models.DB.Where("company_id=?", id).First(&vueBasic)
			vueCap := models.VueDemandCapability{}
			models.DB.Where("c_uid=?", id).First(&vueCap)
			vueDevice := models.VueDemandDevice{}
			models.DB.Where("t_uid=?", vueCompanyInfo.TUid).First(&vueDevice)

			//存入对应order表
			orderDemandInfo := models.OrderDemand{
				CompanyId:    id,
				Uid:          vueCompanyInfo.TUid,
				UserType:     vueCompanyInfo.UserType,
				CName:        vueCompanyInfo.CName,
				Phone:        vueCompanyInfo.CPhone,
				Address:      vueCompanyInfo.CAddress,
				FirstTrial:   0,
				Review:       0,
				BasicId:      vueBasic.Id, //对应basic数据
				DeviceId:     vueDevice.Id,
				CapId:        vueCap.Id,
				FirTime:      int(models.GetUnix()),
				SecTime:      0,
				ThiTime:      0,
				FouTime:      0,
				OrderId:      orderId,
				ProcessTime:  0,
				FinishTime:   0,
				OrderStatus:  1,
				SignContract: 0,
				Status:       0,
			}
			errInfo := models.DB.Create(&orderDemandInfo).Error
			//存入到后端数据库respondManager库中（后加内容）0
			respondManager := models.RespondManager{
				TUid:          vueCompanyInfo.TUid,
				CUid:          vueCompanyInfo.CUid,
				CName:         vueCompanyInfo.CName,
				CPhone:        vueCompanyInfo.CPhone,
				UserType:      vueCompanyInfo.UserType,
				CAddress:      vueCompanyInfo.CAddress,
				ContractNo:    vueCompanyInfo.ContractNo,
				FirstTrial:    0,
				Review:        0,
				CapId:         vueCap.Id,
				FirLiaison:    vueDemandBasic.FirLiaison,
				FirLiaisonWay: vueDemandBasic.FirLiaisonWay,
				LinkAddress:   vueDemandBasic.LinkAddress,
				ContractCap:   vueDemandBasic.ContractCap,
				RunningCap:    vueDemandBasic.RunningCap,
				MonIns:        vueDemandBasic.MonIns,
				BiddingPrice:  vueDemandBasic.BiddingPrice,
				BeResponse:    vueDemandDevice.BeResponse,
				OrderId:       orderId,
				OrderStatus:   1,
				Status:        0,
				SignContract:  0,
				AddTime:       int(models.GetUnix()),
			}
			errManager := models.DB.Create(&respondManager).Error
			if errManager != nil {
				c.Data["json"] = map[string]interface{}{
					"success":    false,
					"resultCode": "errManager",
					"message":    errManager,
				}
				c.ServeJSON()
				return
			}

			if errInfo != nil {
				c.Data["json"] = map[string]interface{}{
					"success":    false,
					"resultCode": "errInfo",
					"message":    errInfo,
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success": true,
					"message": "数据提交成功！",
				}
				c.ServeJSON()
				return
			}

		}
	}

}
