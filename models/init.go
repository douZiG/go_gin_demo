package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MongoDBHosts = "0.0.0.0:27017"
	AuthDatabase = "test"
	AuthUserName = "root"
	AuthPassword = "123456"
	MaxCon       = 300
)

var (
	TestDb *mgo.Database
)

const test string = "test"

func init() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	fmt.Println([]string{MongoDBHosts})
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal("CreateSession failed:%\n", err)
	}

	//设置连接池的大小
	session.SetPoolLimit(MaxCon)
	session.SetMode(mgo.Eventual, true)
	TestDb = session.DB(test)
	indexInit()
	//defer session.Close()
}

func indexInit() {
	index1 := mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Sparse:     true,
		Background: true,
	}

	index2 := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Sparse:     true,
		Background: true,
	}
	for _, index := range []mgo.Index{index1, index2} {
		err := TestDb.C("t_user").EnsureIndex(index)
		fmt.Println(err)
		if err != nil {
			logs.Error("创建索引失败")
			break
		}
	}

}
