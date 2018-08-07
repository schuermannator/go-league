package main

import (
  "fmt"
	"log"
  "github.com/gin-gonic/gin"
  "net/http"
)

func main()  {
  log.Print("Starting server...")
  router := gin.Default()
  router.Static("/static", "./static")
  router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
  router.POST("/", func(c *gin.Context) {
    name := c.PostForm("text")
    fmt.Println(name)
    c.HTML(http.StatusOK, "chart.html", gin.H{
			"title": "Chart time!",
      "values": []int{2, 3, 5},
      "labels": []string{"1", "2", "3"},
      "max": 10,
		})
  })
	router.Run() // listen and serve on 0.0.0.0:8080
}
