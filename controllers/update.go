package controllers

import (
	"FileServer/models"

	"github.com/astaxie/beego"
)

type UpdateController struct {
	beego.Controller
}

// 检查更新
func (u *UpdateController) Get() {
	filetype := u.GetString("filetype")

	switch filetype {
	case "APK":
		packagename := u.GetString("packagename")
		u.Data["json"] = models.GetApkInfo(packagename)
		break

	case "BIN":
		id := u.GetString("id")
		u.Data["json"] = models.GetBinInfo(id)
		break

	default:
		u.Data["json"] = models.GenerateError("未知的文件类型")
		break
	}
	u.ServeJSON()
}
