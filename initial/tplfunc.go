package initial

import (
	"github.com/lock-upme/beegoblog/utils"

	"github.com/astaxie/beego"
)

func InitTplFunc() {
	beego.AddFuncMap("date_mh", utils.GetDateMH)
	beego.AddFuncMap("date", utils.GetDate)
	beego.AddFuncMap("avatar", utils.GetGravatar)

}
