package main

import (
	"backend/config"
	"backend/router"
	"fmt"
)

func main() {
	config.Initconfig()
	fmt.Println(config.Appconfig.App.Port)

	type Msg_info struct {
		Msg string
	}

    info := Msg_info{
		Msg: "Hello World",
	}
	fmt.Println(info) 
	r := router.SetupRouter()

	port := config.Appconfig.App.Port

	if port ==""{
		port = ":8080"
		fmt.Println("配置文件中端口号为空")
	}
	r.Run(config.Appconfig.App.Port) // listen and serve on 0.0.0.0:8080
}