package main

import "github.com/gin-gonic/gin"

func main() {
	e := gin.New()
	e.GET("/index", index)
	e.Run(":1701")
}

func index(c *gin.Context) {
	c.File("index.html")
}
