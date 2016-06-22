package models

import (
	//"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	Id       int
	Title    string
	Picture  string
	Keywords string
	Summary  string
	Created  int64
	Viewnum  int
	Status   int
}

func (this *Album) TableName() string {
	return "album"
}

func init() {
	orm.RegisterModel(new(Album))
}

func GetAlbum(id int) (Album, error) {
	o := orm.NewOrm()
	o.Using("default")
	alb := Album{Id: id}
	err := o.Read(&alb)

	//if err == orm.ErrNoRows {
	//return alb, nil
	//}
	return alb, err
}

func UpdateAlbum(id int, updAlb Album) error {
	o := orm.NewOrm()
	o.Using("default")
	alb := Album{Id: id}

	alb.Title = updAlb.Title
	alb.Summary = updAlb.Summary
	alb.Status = updAlb.Status
	_, err := o.Update(&alb, "title", "summary", "status")
	return err
}

func AddAlbum(updAlb Album) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	alb := new(Album)

	alb.Title = updAlb.Title
	alb.Picture = updAlb.Picture
	alb.Keywords = updAlb.Keywords
	alb.Summary = updAlb.Summary
	alb.Created = time.Now().Unix()
	alb.Viewnum = 1
	alb.Status = updAlb.Status

	id, err := o.Insert(alb)
	return id, err
}

func ListAlbum(condArr map[string]string, page int, offset int) (num int64, err error, alb []Album) {
	o := orm.NewOrm()
	qs := o.QueryTable("album")
	cond := orm.NewCondition()
	if condArr["title"] != "" {
		cond = cond.And("title__icontains", condArr["title"])
	}
	if condArr["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condArr["keywords"])
	}
	if condArr["status"] != "" {
		//status, _ := strconv.Atoi(condArr["status"])
		cond = cond.And("status", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}
	start := (page - 1) * offset
	var albums []Album
	num, err1 := qs.Limit(offset, start).All(&albums)
	return num, err1, albums
}

func CountAlbum(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("album")
	cond := orm.NewCondition()
	if condArr["title"] != "" {
		cond = cond.And("title__icontains", condArr["title"])
	}
	if condArr["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condArr["keywords"])
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
