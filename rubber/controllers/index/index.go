/**
  @Author ZXQ-Administrator
  @Date 2021-08-15
  @node:
**/
package index

//type IndexController struct {
//	beego.Controller
//}
//
//// 本次将代码下载到本地 调用该方法，直接引用地址参考：https://github.com/hunterhug/go_image
//func (c *IndexController) Get() {
//	//实现图片剪切 输入输出都是按宽度进行比例缩放
//	filename := "static/upPicture/aaa.jpg"
//	savepath := "static/upPicture/aaa_800.jpg"
//	err := go_image.ScaleF2F(filename, savepath, 800)
//	if err != nil {
//		beego.Error(err)
//	}
//
//	// filename := "static/upload/a.jpg"
//	// savepath := "static/upload/a_800_300.jpg"
//
//	// err := ThumbnailF2F(filename, savepath, 800, 300)
//	// if err != nil {
//	// 	beego.Error(err)
//	// }
//
//	//按宽度和高度进行比例缩放，输入和输出都是文件
//	// filename := "static/upload/b.png"
//	// savepath := "static/upload/b_400.png"
//	// err := ThumbnailF2F(filename, savepath, 400, 400)
//	// if err != nil {
//	// 	beego.Error(err)
//	// }
//
//	//https://github.com/skip2/go-qrcode
//	//生成二维码
//	// 链接地址 二维码质量 二维码宽度 输出路径
//	err5 := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 800, "static/upPicture/qr.png")
//	if err5 != nil {
//		beego.Error("生成二维码失败")
//	}
//	c.TplName = "front/index/index.html"
//}
