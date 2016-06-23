package controllers

import (
	//"fmt"

	"io"

	"github.com/lock-upme/beegoblog/utils"
	//"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "www.beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (this *MainController) Go404() {
	this.TplName = "404.tpl"
	return
}

//单文件上传
type UploadController struct {
	BaseController
}

func (this *UploadController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": "你没有权限上传"}
		this.ServeJSON()
		return
	}
	//imgFile
	f, h, err := this.GetFile("imgFile")
	defer f.Close()

	//生成上传路径
	now := time.Now()
	dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}
	//生成新的文件名
	filename := h.Filename
	ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
	filename = utils.GetGuid() + ext

	if err != nil {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": err}
	} else {
		//this.SaveToFile("imgFile", "./static/uploadfile/"+h.Filename)
		this.SaveToFile("imgFile", dir+"/"+filename)
		this.Data["json"] = map[string]interface{}{"error": 0, "url": strings.Replace(dir, ".", "", 1) + "/" + filename}
	}
	this.ServeJSON()
}

//多文件上传
type UploadMultiController struct {
	BaseController
}

func (this *UploadMultiController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "你没有权限上传"}
		this.ServeJSON()
		return
	}

	files, err := this.GetFiles("uploadFiles")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}

	//生成上传路径
	now := time.Now()
	dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}

	resfilestr := ""
	resfilename := ""
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}

		//生成新的文件名
		filename := files[i].Filename
		resfilename += utils.GetFileSuffix(filename) + "||"

		ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
		filename = utils.GetGuid() + ext
		dst, err := os.Create(dir + "/" + filename)

		defer dst.Close()
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}
		if _, err := io.Copy(dst, file); err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}
		resfilestr += strings.Replace(dir, ".", "", 1) + "/" + filename + "||"
	}
	this.SetSession("uploadMultiPic", resfilestr)
	this.SetSession("uploadMultiName", resfilename)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功", "url": resfilestr}
	this.ServeJSON()
	return
}
