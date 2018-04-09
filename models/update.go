package models

func GetBinInfo(id string) Response {
	var doc Doc
	err := Orm.QueryTable("doc").Filter("id", id).One(&doc)
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("查询成功", doc)
}

func GetApkInfo(packagename string) Response {
	var apk Apk
	err := Orm.QueryTable("apk").Filter("packagename", packagename).One(&apk)
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("查询成功"+packagename, apk)
}
