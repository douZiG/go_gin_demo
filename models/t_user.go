package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"newProject/log"

	"gopkg.in/mgo.v2/bson"
	"time"
)

var logs = log.Logs

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Role       string        `json:"role" bson:"role"`
	Phone      string        `json:"phone" bson:"phone"`
	Mail       string        `json:"mail" bson:"mail"`
	Age        int           `json:"age" bson:"age"`
	Status     int           `json:"status" bson:"status"`
	CreateTime time.Time     `json:"create_time" bson:"create_time"`
}

func QueryUser(database *mgo.Database) ([]User, error) {
	logs.Info("开始查询数据库", "a")

	var userList []User
	err := database.C("t_user").Find(bson.M{}).All(&userList)
	if err != nil {
		logs.Error("查询数据库失败, 失败信息为： ", err)
	}
	return userList, err
}

func AddUser(database *mgo.Database, user User) (bool, string) {
	// 判断用户是否存在
	var userList []User
	fmt.Println(user)
	finErr := database.C("t_user").Find(bson.M{"name": user.Name}).All(&userList)
	if finErr != nil {
		logs.Error("查询数据库失败, 失败信息为： ", finErr)
		return false, finErr.Error()
	}
	if len(userList) == 0 {
		client := database.C("t_user")
		// 插入数据
		fmt.Println(user)
		user.Id = bson.NewObjectId()
		insertErr := client.Insert(&user)
		if insertErr != nil {
			return false, insertErr.Error()
		}
		return true, "创建人员成功"
	} else {
		return false, "该用户已存在"
	}
	//client := database.C("t_user")
	//// 插入数据
	//fmt.Println(user)
	//insertErr := client.Insert(&user)
	//if insertErr != nil {
	//	return false, insertErr.Error()
	//}
	//return true, "创建人员成功"
}

func UpdateUser(database *mgo.Database, data User) (bool, string) {
	// 判断用户是否存在
	var userList []User
	fmt.Println(data)
	finErr := database.C("t_user").Find(bson.M{"name": data.Name}).All(&userList)
	if finErr != nil {
		logs.Error("查询数据库失败, 失败信息为： ", finErr)
		return false, finErr.Error()
	}
	if len(userList) == 0 {
		return false, "该用户不存在"
	} else {
		fmt.Println(userList)
		client := database.C("t_user")
		for _, user := range userList {
			fmt.Println(data.Id)
			data.Id = user.Id
			fmt.Println(data.Id)
			updateErr := client.Update(bson.M{"_id": user.Id}, bson.M{"$set": &data})
			if updateErr != nil {
				return false, updateErr.Error()
			}
		}
		return true, "更新人员信息成功"
	}
}

func DeleteUser(database *mgo.Database, data User) (bool, string) {
	var userList []User
	finErr := database.C("t_user").Find(bson.M{"_id": data.Id}).All(&userList)
	if finErr != nil {
		logs.Error("查询数据库失败, 失败信息为： ", finErr)
		return false, finErr.Error()
	}
	if len(userList) == 0 {
		return false, "该用户不存在"

	} else {
		fmt.Println(2)
		client := database.C("t_user")
		for _, user := range userList {
			fmt.Println(user.Id)
			deleteErr := client.Remove(bson.M{"_id": user.Id})
			if deleteErr != nil {
				fmt.Println(deleteErr.Error())
				return false, deleteErr.Error()
			}
		}
		return true, "删除人员信息成功"
	}
}
