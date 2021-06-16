package controllers

import (
	"../log"
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"reflect"
)

var logs = log.Logs

func QueryUser(c *gin.Context) {
	//name := c.Param("name")
	user, err := models.QueryUser(models.TestDb)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  0,
			"message": err,
			"data":    []string{},
		})
	}
	fmt.Println(reflect.TypeOf(user))

	c.JSON(200, gin.H{
		"status":  1,
		"message": "ok",
		"data":    &user,
	})
}

func AddUser(c *gin.Context) {
	fmt.Println("origin:", c.Request.Header.Get("Origin"))

	requestData := make(map[string]interface{}) //注意该结构接受的内容
	jsonErr := c.BindJSON(&requestData)
	if jsonErr != nil {
		c.JSON(500, gin.H{
			"status":  0,
			"message": "接收请求参数出错",
		})
		return
	}
	// 判断参数是否存在
	_, name := requestData["name"]
	_, phone := requestData["phone"]
	if name != true || phone != true {
		c.JSON(504, gin.H{
			"status":  0,
			"message": "缺少参数",
		})
		return
	}
	fmt.Println(requestData)
	var user models.User
	requestData["status"] = 1
	if err := mapstructure.Decode(requestData, &user); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  0,
			"message": err.Error(),
		})
		return
	}
	insertResult, errMessage := models.AddUser(models.TestDb, user)
	fmt.Println(insertResult)
	if insertResult != true {
		logs.Info("创建人员失败, 失败原因为:" + errMessage)
		c.JSON(200, gin.H{
			"status":  0,
			"message": errMessage,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "创建人员成功",
			"data":    "",
		})
	}
}

func UpdateUser(c *gin.Context) {
	requestData := make(map[string]interface{}) //注意该结构接受的内容
	jsonErr := c.BindJSON(&requestData)
	if jsonErr != nil {
		c.JSON(500, gin.H{
			"status":  0,
			"message": "接收请求参数出错",
		})
		return
	}
	// 判断参数是否存在
	_, name := requestData["name"]
	_, phone := requestData["phone"]
	if name != true || phone != true {
		c.JSON(504, gin.H{
			"status":  0,
			"message": "缺少参数",
		})
		return
	}
	var user models.User
	if err := mapstructure.Decode(requestData, &user); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  0,
			"message": err.Error(),
		})
		return
	}
	insertResult, errMessage := models.UpdateUser(models.TestDb, user)
	fmt.Println(insertResult)
	if insertResult != true {
		logs.Info("更新人员信息失败, 失败原因为:" + errMessage)
		c.JSON(200, gin.H{
			"status":  0,
			"message": errMessage,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "更新人员信息成功",
			"data":    "",
		})
	}
}

func DeleteUser(c *gin.Context) {
	requestData := make(map[string]interface{}) //注意该结构接受的内容
	jsonErr := c.BindJSON(&requestData)
	if jsonErr != nil {
		c.JSON(500, gin.H{
			"status":  0,
			"message": "接收请求参数出错",
		})
		return
	}
	// 判断参数是否存在
	_, name := requestData["name"]
	if name != true {
		c.JSON(504, gin.H{
			"status":  0,
			"message": "缺少参数name",
		})
		return
	}
	var user models.User
	if err := mapstructure.Decode(requestData, &user); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  0,
			"message": err.Error(),
		})
		return
	}
	insertResult, errMessage := models.DeleteUser(models.TestDb, user)
	fmt.Println(insertResult)
	if insertResult != true {
		logs.Info("删除人员信息失败, 失败原因为:" + errMessage)
		c.JSON(200, gin.H{
			"status":  0,
			"message": errMessage,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "删除人员信息成功",
			"data":    "",
		})
	}

}
