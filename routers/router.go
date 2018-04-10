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
		// 文件
		beego.NSRouter("/file", &controllers.FileController{}),
		// 检查更新
		beego.NSRouter("/update", &controllers.UpdateController{}),
	)

	beego.AddNamespace(ns)

}
