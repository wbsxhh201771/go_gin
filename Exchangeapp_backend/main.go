package main

import (
	"backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,info)
	})
	r.Run(config.Appconfig.App.Port) // listen and serve on 0.0.0.0:8080
}