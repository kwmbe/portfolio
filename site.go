package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(static.Serve("/", static.LocalFile("./static", true)))
  r.LoadHTMLGlob("templates/*")
  // r.GET("/test", func(c *gin.Context) {
  //   c.HTML(200, "template.html", gin.H{
  //     "content": "test",
  //   })
  // })
  r.Run()
}

