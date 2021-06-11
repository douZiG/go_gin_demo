package main

import (
	"./log"
	"./routers"
	"fmt"
)

func init() {
	log.InitLog()
}

func main() {
	fmt.Println("a")
	router := routers.ConfigRouters()
	router.Run(":8080")
}
