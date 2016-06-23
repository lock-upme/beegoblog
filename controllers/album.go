package controllers

import (
	"fmt"
	"strings"

	. "github.com/lock-upme/beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//upload
type UploadAlbumController struct {
	BaseController
}

func (this *UploadAlbumController) Get() {
	if !this.isLogin {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "album-upload.tpl"
}

//修改blog
type EditAlbumController struct {
	BaseController
}

func (this *EditAlbumController) Post() {
	id, err := this.GetInt("id")
	title := this.GetString("title")
	summary := this.GetString("summary")
	status, _ := this.GetInt("status")

	if "" == title {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写标题"}
		this.ServeJSON()
	}
	_, errAttr := GetAlbum(id)
	if errAttr != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "相册不存在"}
		this.ServeJSON()
	}

	var alb Album
	alb.Title = title
	alb.Summary = summary
	alb.Status = status

	err = UpdateAlbum(id, alb)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "相册修改成功", "id": id}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "相册修改出错"}
	}
	this.ServeJSON()
}

//列表
type ListAlbumController struct {
	BaseController
}

func (this *ListAlbumController) Get() {

	//添加上传的图片到表
	str := this.GetSession("uploadMultiPic")
	if str != nil {
		str = strings.Trim(str.(string), "||")
		strPic := strings.Split(str.(string), "||")

		strn := this.GetSession("uploadMultiName")
		strn = strings.Trim(strn.(string), "||")
		strName := strings.Split(strn.(string), "||")

		//fmt.Println(strName)

		for i, pic := range strPic {
			var alb Album
			alb.Picture = pic
			alb.Title = strName[i]
			alb.Status = 1

			_, err2 := AddAlbum(alb)
			fmt.Println(err2)
		}
		this.DelSession("uploadMultiName")
		this.DelSession("uploadMultiPic")
	}

	page, err1 := this.GetInt("p")
	title := this.GetString("title")
	keywords := this.GetString("keywords")
	status := this.GetString("status")
	if err1 != nil {
		page = 1
	}

	offset, err2 := beego.AppConfig.Int("pageoffset")
	if err2 != nil {
		offset = 9
	}

	condArr := make(map[string]string)
	condArr["title"] = title
	condArr["keywords"] = keywords
	if !this.isLogin {
		condArr["status"] = "1"
	} else {
		condArr["status"] = status
	}
	countAlbum := CountAlbum(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countAlbum)
	_, _, alb := ListAlbum(condArr, page, offset)

	fmt.Println(countAlbum)

	this.Data["paginator"] = paginator
	this.Data["alb"] = alb
	this.TplName = "album.tpl"
}
