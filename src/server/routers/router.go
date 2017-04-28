package routers

import (
	"server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/tests", &controllers.MainController{})
}
