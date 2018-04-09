// upload
package controllers

import (
	"FileServer/models"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// 文件
type FileController struct {
	beego.Controller
}

// 接收上传的文件

func (u *FileController) Post() {
	f, h, err := u.GetFile("file_name")
	if err != nil {
		log.Fatal(" UPLOAD FAILED")
	}
	defer f.Close()
	//APK文件
	if strings.HasSuffix(h.Filename, "apk") {
		fmt.Println("apk")
		u.SaveToFile("file_name", "static/apk/"+h.Filename)
		var apk models.Apk
		apk.Packagename = u.GetString("packagename", "")
		apk.Name = h.Filename
		apk.Url = "http://127.0.0.1:8080/apk/" + h.Filename
		apk.Versioncode, _ = u.GetInt("versioncode", 0)
		apk.Versionname = u.GetString("versionname", "")
		apk.User = models.Search(u.GetString("user", ""))
		apk.Filetype = "APK"
		apk.Datetime = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(apk)
		u.Data["json"] = models.InsertOrUpdateApk(apk)
		// BIN文件
	} else if strings.HasSuffix(h.Filename, "bin") {
		fmt.Println("bin")
		var bin models.Bin
		bin.Name = h.Filename
		bin.Url = "http://127.0.0.1:8080/bin/" + h.Filename
		bin.Hardcode, _ = u.GetInt("hardcode", 0)
		bin.Softcode, _ = u.GetInt("softcode", 0)
		bin.User = models.Search(u.GetString("user", ""))
		bin.Filetype = "BIN"
		bin.Id = u.GetString("id", "")
		bin.Datetime = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(bin)
		u.Data["json"] = models.InsertOrUpdateBin(bin)
		u.SaveToFile("file_name", "static/bin/"+h.Filename)
	} else {
		// DOC文档
		fmt.Println("doc")
		var doc models.Doc
		doc.Name = h.Filename
		doc.Url = "http://127.0.0.1:8080/bin/" + h.Filename
		doc.Version, _ = u.GetInt("version", 0)
		doc.User = models.Search(u.GetString("user", ""))
		doc.Filetype = "DOC"
		doc.Datetime = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(doc)
		u.Data["json"] = models.InsertOrUpdateDoc(doc)
		u.SaveToFile("file_name", "static/doc/"+h.Filename)
	}
	//u.Data["json"] = models.GenerateSuccess("上传成功")
	//	go func() {
	//		versionCode, versionName := util.ParseApk("./cache/" + h.Filename)
	//		fmt.Println(versionName)
	//		apkInfo := models.ApkInfo{versionCode, versionName, "http://127.0.0.1:8080/static/apk/" + h.Filename}
	//		models.SaveToJson(&apkInfo)
	//	}()
	u.ServeJSON()

}

func (u *FileController) Get() {
	apkFiles := models.GetApkFiles()
	binFiles := models.GetBinFiles()
	docFiles := models.GetDocFiles()
	slice := make([]interface{}, len(apkFiles)+len(binFiles)+len(docFiles))
	for key, value := range apkFiles {
		slice[key] = value
	}
	for key, value := range binFiles {
		slice[key+len(apkFiles)] = value
	}
	for key, value := range docFiles {
		slice[key+len(apkFiles)+len(binFiles)] = value
	}
	//	slice := make([]interface{})
	//	copy(slice, apkFiles)
	//	copy(slice[len(apkFiles):], binFiles)
	//	copy(slice[len(apkFiles)+len(binFiles):], docFiles)
	//files := models.GetFiles()
	u.Data["json"] = models.Response{Msg: "查询成功", Data: slice, Success: 0}
	u.ServeJSON()
}

func (u *FileController) Delete() {
	filetype := u.GetString("filetype")

	switch filetype {
	case "APK":
		packagename := u.GetString("packagename")
		u.Data["json"] = models.DeleteApk(packagename)
		break

	case "BIN":
		id := u.GetString("id")
		u.Data["json"] = models.DeleteBin(id)
		break

	case "DOC":
		id := u.GetString("id")
		u.Data["json"] = models.DeleteDoc(id)
		break

	}
	u.ServeJSON()
}
