package controllers

import (
	"strconv"

	. "github.com/lock-upme/beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//添加blog
type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {
	if !this.isLogin {
		this.Redirect("/login", 302)
		return
	}
	/*userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.Redirect("/login", 302)
		return
	}
	*/
	var art Article
	art.Status = 1
	this.Data["art"] = art
	this.TplName = "article-form.tpl"
}

func (this *AddArticleController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请先登录"}
		this.ServeJSON()
		return
	}
	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")
	uri := this.GetString("uri")
	summary := this.GetString("summary")
	author := this.GetString("author")

	if "" == title {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写标题"}
		this.ServeJSON()
		return
	}

	if "" == content {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写内容"}
		this.ServeJSON()
		return
	}

	var art Article
	art.Title = title
	art.Keywords = keywords
	art.Uri = uri
	art.Summary = summary
	art.Content = content
	art.Author = author

	id, err := AddArticle(art)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "博客添加成功", "id": id}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "博客添加出错"}
	}
	this.ServeJSON()
}

//修改blog
type EditArticleController struct {
	BaseController
}

func (this *EditArticleController) Get() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请先登录"}
		this.ServeJSON()
		return
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)

	art, err := GetArticle(id)
	if err != nil {
		this.Redirect("/404.html", 302)
	}
	//this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	//this.ServeJSON()
	this.Data["art"] = art

	this.TplName = "article-form.tpl"
}

func (this *EditArticleController) Post() {
	id, err := this.GetInt("id")
	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")
	uri := this.GetString("uri")
	summary := this.GetString("summary")
	author := this.GetString("author")
	status, _ := this.GetInt("status")

	if "" == title {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写标题"}
		this.ServeJSON()
	}

	if "" == content {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写内容"}
		this.ServeJSON()
	}
	_, errAttr := GetArticle(id)
	if errAttr != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "博客不存在"}
		this.ServeJSON()
	}

	var art Article
	art.Title = title
	art.Keywords = keywords
	art.Uri = uri
	art.Summary = summary
	art.Content = content
	art.Author = author
	art.Status = status

	err = UpdateArticle(id, art)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "博客修改成功", "id": id}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "博客修改出错"}
	}
	this.ServeJSON()
}

//列表
type ListArticleController struct {
	BaseController
}

func (this *ListArticleController) Get() {
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
	countArticle := CountArticle(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countArticle)
	_, _, art := ListArticle(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["art"] = art
	//userLogin := this.GetSession("userLogin")
	//this.Data["isLogin"] = userLogin
	//this.Data["isLogin"] = this.isLogin

	this.TplName = "article.tpl"
}

//详情
type ShowArticleController struct {
	//beego.Controller
	BaseController
}

func (this *ShowArticleController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)

	art, err := GetArticle(id)
	if err != nil {
		this.Redirect("/404.html", 302)
	}
	if !this.isLogin {
		if art.Status == 0 {
			this.Redirect("/404.html", 302)
		}
	}
	this.Data["art"] = art

	//评论分页
	page, err1 := this.GetInt("p")
	if err1 != nil {
		page = 1
	}
	offset, err2 := beego.AppConfig.Int("pageoffset")
	if err2 != nil {
		offset = 9
	}
	condCom := make(map[string]string)
	condCom["article_id"] = idstr
	if !this.isLogin {
		condCom["status"] = "1"
	}
	countComment := CountComment(condCom)
	paginator := pagination.SetPaginator(this.Ctx, offset, countComment)
	_, _, coms := ListComment(condCom, page, offset)
	this.Data["paginator"] = paginator
	this.Data["coms"] = coms

	this.TplName = "article-detail.tpl"
}
