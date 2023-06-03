package routers

import (
	"github.com/astaxie/beego"
	"ikunsApi/controllers"
)

func init() {
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.VideoController{})
	beego.Include(&controllers.BaseController{})
	beego.Include(&controllers.CommentController{})
	beego.Include(&controllers.TopController{})
	beego.Include(&controllers.BarrageController{})
	beego.Include(&controllers.AliyunController{})
}
