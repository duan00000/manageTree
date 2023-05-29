/**
  @Author ZXQ-Administrator
  @Date 2021-08-19
  @node: 基类
**/
package admin

import (
	"elenewenergy/models"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
}

// 返回成功
func (c *BaseController) Success(message string, redirect string) {
	//http://localhost:8080/state_grid/goods?page=2
	//判断传入地址有没有 adminPath 地址
	c.Data["message"] = message
	if strings.Contains(redirect, beego.AppConfig.String("adminPath")) {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/success.html"
}

// 返回失败
func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	if strings.Contains(redirect, beego.AppConfig.String("adminPath")) {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/error.html"

}

// 302
func (c *BaseController) Goto(redirect string) {
	c.Redirect("/"+beego.AppConfig.String("adminPath")+redirect, 302)
}

// 公共上传图片的方法 title为传入的标题名 判断上传类型

func (c *BaseController) UploadImg(picName string, title string) (string, error) {

	ossStatus, _ := beego.AppConfig.Bool("ossStatus")
	if ossStatus == true {
		return c.OssUploadImg(picName,title)
	} else {
		return c.LocalUploadImg(picName,title)
	}
}

// 本地上传
func (c *BaseController) LocalUploadImg(picName string, title string) (string, error) {
	//1、获取上传的文件
	files, heads, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer files.Close()
	//3、获取后缀名 判断类型是否正确 .jpg .jpeg .png .gif
	extName := path.Ext(heads.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件类型不合法")
	}
	//4、创建图片保存目录 static/upPicture/20210902
	day := models.GetDay() //当前时间
	dir := "static/upPicture/" + day
	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", errors.New("创建目录失败")
	}
	//5、生成文件名称
	//	当前时间戳转换成string
	fileUnixName := strconv.FormatInt(models.GetUnixNano(), 10)
	// static/upPicture/20210902/title_1630572883.png
	saveDir := path.Join(dir, title+"_"+fileUnixName+extName)
	//6、保存图片 static/upPicture/20210902/1630572883.jpg
	if errSaveToFile := c.SaveToFile(picName, saveDir); errSaveToFile != nil {
		return "", errors.New("保存图片失败")
	}
	return saveDir, nil
}

// 本地上传
func (c *BaseController) OssUploadImg(picName string, title string) (string, error) {

	//获取系统信息
	setting := models.Setting{}
	models.DB.First(&setting)

	//1、获取上传的文件
	files, heads, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer files.Close()
	//3、获取后缀名 判断类型是否正确 .jpg .jpeg .png .gif
	extName := path.Ext(heads.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件类型不合法")
	}
	//4、把文件流上传到oss
	// 4.1创建OSSClient实例
	client, err := oss.New(setting.EndPoint, setting.Appid, setting.AppSecret)
	if err != nil {
		return "", err
	}
	// 4.2获取存储空间
	bucket,err := client.Bucket(setting.BucketName)
	if err != nil {
		return "", err
	}

	//4、创建图片保存目录 static/upPicture/20210902
	day := models.GetDay() //当前时间
	dir := "static/upPicture/" + day
	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", errors.New("创建目录失败")
	}
	//5、生成文件名称
	//	当前时间戳转换成string
	fileUnixName := strconv.FormatInt(models.GetUnixNano(), 10)
	// static/upPicture/20210902/title_1630572883.png
	saveDir := path.Join(dir, title+"_"+fileUnixName+extName)
	// 5.1上传文件流。
	err = bucket.PutObject(saveDir, files)
	if err != nil {
		return "", err
	}
	return saveDir, nil
}

func (c *BaseController) GetSetting() models.Setting {
	setting := models.Setting{}
	models.DB.First(&setting)
	return setting
}
//求和
func (c *BaseController) Sum(x [20]float64, size int) float64 {
	var i int
	var sum float64
	for i=0;i<size;i++{
		sum = sum + x[i]
	}
	return sum
}
//求平均值
func (c *BaseController) Avg(x [20]float64, size int) float64 {
	var i int
	var sum,avg float64
	for i=0;i<size;i++{
		sum = sum + x[i]
	}
	avg=sum/float64(size)
	return avg
}
