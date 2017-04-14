package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"use_gin/api"
)

func main() {
	serv := gin.New()
	serv.Use(gin.Recovery())
	serv.Use(midWare())
	api.Install(serv)
	//serv.GET("/home/", index)
	serv.Run(":8888")
}

func index (c *gin.Context)  {
	fmt.Println(c.HandlerName())
	c.JSON(1, "ok")
}

func midWare()gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println(c.HandlerName())
		handlerPath := strings.Split(c.HandlerName(), "/")
		var moduleName, funcName string
		if len(handlerPath) == 5 {
			handlerName := strings.Split(handlerPath[4], ".")
			if len(handlerName) == 2 {
				moduleName, funcName = handlerName[0], handlerName[1]
			}
		}
		fmt.Println("xxxx")
		fmt.Println(moduleName)
		fmt.Println(funcName)
		fmt.Println("xxxx2")
	}
}