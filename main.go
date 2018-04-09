package main

import (
	_ "FileServer/routers"

	"FileServer/models"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.SetStaticPath("/apk", "static/apk")
	beego.SetStaticPath("/bin", "static/bin")
	beego.SetStaticPath("/doc", "static/doc")
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Credentials", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	models.Init()
	beego.Run()
}

func TransparentStatic(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}
