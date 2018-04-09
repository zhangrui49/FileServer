package routers

import (
	"FileServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSRouter("/file", &controllers.FileController{}),
		beego.NSRouter("/update", &controllers.UpdateController{}),
	)

	beego.AddNamespace(ns)

}
