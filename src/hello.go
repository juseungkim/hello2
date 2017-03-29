package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init(){
	app := gin.Default()
	app.GET("/", index)
	http.Handle("/", app)
}

func index (c *gin.Context){
	content := gin.H{"message" : "HELL!O GO"}
	c.JSON(200, content)

}

