package main

import (
	"fmt"
	"newProject/log"
	"newProject/routers"
)

func init() {
	log.InitLog()
}

func main() {
	fmt.Println("a")
	router := routers.ConfigRouters()
	router.Run(":8080")
}
