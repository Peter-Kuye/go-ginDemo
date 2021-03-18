package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//router.GET("/", func(c *gin.Context){
	//	c.String(http.StatusOK,"HEllO World")
	//})
	//err := router.Run(":8081")
	//http.ListenAndServe(":8081",router)
	//if err != nil {
	//	log.Panicln("服务器启动失败: ",err.Error())
	//}
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8081")
}
