package controllers

import (
	"time"

	. "github.com/lock-upme/beegoblog/models"
)

//添加评论
type AddCommentController struct {
	BaseController
}

func (this *AddCommentController) Post() {
	/*if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请先登录"}
		this.ServeJSON()
		return
	}*/
	nickname := this.GetString("nickname")
	article_id, _ := this.GetInt("article_id")
	content := this.GetString("content")
	uri := this.GetString("uri")

	if "" == nickname {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写昵称"}
		this.ServeJSON()
		return
	}

	if "" == content {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写内容"}
		this.ServeJSON()
		return
	}

	var com Comment
	com.Nickname = nickname
	com.ArticleId = article_id
	com.Uri = uri
	com.Content = content
	com.Status = 1
	com.Created = time.Now().Unix()

	id, err := AddComment(com)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "评论添加成功", "id": id}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "评论添加出错"}
	}
	this.ServeJSON()
}

//修改
type EditCommentController struct {
	BaseController
}

func (this *EditCommentController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": this.isLogin}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt("id")
	status, _ := this.GetInt("status")

	var com Comment
	com.Status = status

	err := UpdateComment(id, com)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "状态更新成功", "id": id}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "状态修改失败"}
	}
	this.ServeJSON()
}
