package models

import (
	"github.com/astaxie/beego/orm"
)

type Version interface {
	GetVersion() int
}

type FileUpload struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Filetype string `json:"filetype"`
	Datetime string `json:"datetime"`
	User     *User  `orm:"rel(fk)"`
}

type Apk struct {
	FileUpload
	Versioncode int    `json:"versioncode"`
	Versionname string `json:"versionname"`
	Packagename string `orm:"pk"`
}

type Bin struct {
	FileUpload
	Id       string `orm:"pk"`
	Hardcode int    `json:"hardcode"`
	Softcode int    `json:"softcode"`
}

type Doc struct {
	FileUpload
	Id      string `orm:"pk"`
	Version int    `json:"version"`
}

func init() {
	orm.RegisterModel(new(Bin))
	orm.RegisterModel(new(Apk))
	orm.RegisterModel(new(Doc))
}

//func SaveToJson(apk *Apk) {
//	fmt.Println(apk)
//	b, err := json.Marshal(apk)
//	if err != nil {
//		panic(err)
//	} else {
//		path := "./static/config/apkinfo.json"
//		//		absPath, _ := filepath.Abs(path)
//		_, err := os.Stat(path)
//		if err == nil {
//			os.Remove(path)
//		}

//		f, osErr := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
//		if osErr != nil {
//			panic(osErr)
//		}
//		fmt.Println("写配置文件成功")
//		f.Write(b)
//	}
//}

func InsertOrUpdateApk(apk Apk) Response {

	_, err := Orm.InsertOrUpdate(&apk)

	if err == nil {
		return GenerateSuccess("APK已添加", apk)
	} else {
		return GenerateError(err.Error())
	}
}

func InsertOrUpdateBin(bin Bin) Response {

	_, err := Orm.InsertOrUpdate(&bin)

	if err == nil {
		return GenerateSuccess("BIN已添加", bin)
	} else {
		return GenerateError(err.Error())
	}
}

func InsertOrUpdateDoc(doc Doc) Response {

	_, err := Orm.InsertOrUpdate(&doc)

	if err == nil {
		return GenerateSuccess("DOC已添加", doc)
	} else {
		return GenerateError(err.Error())
	}
}

func (upload *FileUpload) InsertOrUpdate() Response {

	_, err := Orm.InsertOrUpdate(&upload)

	if err == nil {
		return GenerateSuccess("APK已添加", upload)
	} else {
		return GenerateError(err.Error())
	}
}

func GetApkFiles() []*Apk {
	var files []*Apk
	Orm.QueryTable("apk").All(&files)
	return files
}

func DeleteApk(packagename string) Response {
	var apk Apk
	_, err := Orm.QueryTable("apk").Filter("packagename", packagename).Delete()
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("删除成功"+packagename, apk)
}

func DeleteDoc(id string) Response {
	var doc Doc
	_, err := Orm.QueryTable("doc").Filter("id", id).Delete()
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("删除成功", doc)
}

func DeleteBin(id string) Response {
	var bin Bin
	_, err := Orm.QueryTable("bin").Filter("id", id).Delete()
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("删除成功", bin)
}

func GetBinFiles() []*Bin {
	var files []*Bin
	Orm.QueryTable("bin").All(&files)
	return files
}

func GetDocFiles() []*Doc {
	var files []*Doc
	Orm.QueryTable("doc").All(&files)
	return files
}

func (apk *Apk) GetVersion() int {
	return apk.Versioncode
}
func (bin *Bin) GetVersion() int {
	return bin.Softcode
}
func (doc *Doc) GetVersion() int {
	return doc.Version
}
