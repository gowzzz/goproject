package main

import (
	"fmt"
	"goproject/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var confFile = map[string]string{"feidan": "feidan.yml", "homey": "homey.yml"}
	for k, v := range confFile {
		fmt.Printf("%s:%s\n", k, v)
	}

	return
	router := gin.Default()
	// router.Static("/static", "./static")
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/onefile", handler.OneFile)
	router.Run(":8080")
	handler.UseLog().Info("in project")

}
