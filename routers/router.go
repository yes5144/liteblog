package routers

import (
	"liteblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 注解路由 需要调用Include
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	// beego.Router("/", &controllers.MainController{})
}
