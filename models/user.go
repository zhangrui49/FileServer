package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Name     string `orm:"pk"`
	Password string `json:"password"` //659a845b2eb99d11788fbdb0199539bc
	Role     string `json:"role"`
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u User) Response {
	created, _, err := Orm.ReadOrCreate(&u, "Name")
	if err == nil {
		if created {
			return GenerateSuccess("用户已添加", u)
		} else {
			return GenerateError("用户已存在")
		}
	} else {
		return GenerateError(err.Error())
	}
}

func GetAllUsers() Response {
	var users []*User
	_, err := Orm.QueryTable("user").All(&users)
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("SUCCESS", users)
}

// 登录
func Login(username, password string) Response {
	user := User{Name: username, Password: password}
	fmt.Println(user)
	err := Orm.Read(&user, "Name")
	if err != nil {
		if err == orm.ErrNoRows {
			return GenerateError("NOT EXIST")
		} else if err == orm.ErrMissPK {
			return GenerateError(err.Error())
		} else {
			return GenerateError(err.Error())
		}
	}

	return GenerateSuccess("登录成功", user)
}

func DeleteUser(uid string) Response {
	var user User
	_, err := Orm.QueryTable("user").Filter("id", uid).Delete()
	if err != nil {
		return GenerateError(err.Error())
	}

	return GenerateSuccess("删除成功", user)
}

func Search(name string) *User {
	user := User{Name: name}
	err := Orm.Read(&user, "Name")
	if err != nil {
		return &user
	}

	return &user
}

//func UpdateUser(uid string, user *User) Response {
//	_, err := Orm.QueryTable("user").Filter("id", uid).Delete()
//	if err != nil {
//		return Response{1, err.Error(), nil}
//	}

//	return Response{0, "update success", user}
//}
