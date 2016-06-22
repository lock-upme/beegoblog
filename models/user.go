package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
)

type User struct {
	Id          int
	Phone       string
	UserProfile *UserProfile `orm:"rel(one)"`
	Password    string
	Status      int
	Created     int64
	Changed     int64
}
type UserProfile struct {
	Id       int
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Hobby    string
	Intro    string
	User     *User `orm:"reverse(one)"`
}

func (this *User) TableName() string {
	return "user"
}
func init() {
	orm.RegisterModel(new(User), new(UserProfile)) //
}

func LoginUser(phone string, password string) (err error, user []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	cond := orm.NewCondition()

	cond = cond.And("phone", phone)
	pwdmd5 := com.Md5(password)
	cond = cond.And("password", pwdmd5)
	cond = cond.And("status", 1)

	qs = qs.SetCond(cond)
	var users []User
	err1 := qs.Limit(1).One(&users)
	return err1, users
}

func GetUser(id int) (User, error) {
	var user User
	var err error
	o := orm.NewOrm()

	user = User{Id: id}
	err = o.Read(&user)

	if err == orm.ErrNoRows {
		return user, nil
	}
	return user, err
}

func GetProfile(id int) (UserProfile, error) {
	var pro UserProfile
	var err error
	o := orm.NewOrm()

	pro = UserProfile{Id: id}
	err = o.Read(&pro)

	if err == orm.ErrNoRows {
		return pro, nil
	}
	return pro, err
}

func UpdateProfile(id int, updPro UserProfile) error {
	var pro UserProfile
	o := orm.NewOrm()
	pro = UserProfile{Id: id}

	pro.Realname = updPro.Realname
	pro.Sex = updPro.Sex
	pro.Birth = updPro.Birth
	pro.Email = updPro.Email
	pro.Phone = updPro.Phone
	pro.Address = updPro.Address
	pro.Hobby = updPro.Hobby
	pro.Intro = updPro.Intro
	_, err := o.Update(&pro)
	return err

}

func UpdatePassword(id int, oldPawd string, newPwd string) error {
	o := orm.NewOrm()
	//salt := com.RandString(10)

	user := User{Id: id}
	err := o.Read(&user)
	if nil != err {
		return err
	} else {
		if user.Password == com.Md5(oldPawd) {
			user.Password = com.Md5(newPwd)
			_, err := o.Update(&user)
			return err
		} else {
			return fmt.Errorf("验证出错")
		}
	}
}
