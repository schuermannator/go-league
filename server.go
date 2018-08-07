package main

import (
  "fmt"
	"log"
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
  "sort"
)

// Define us a type so we can sort it
type TimeSlice []time.Time

// Forward request for length
func (p TimeSlice) Len() int {
    return len(p) }

// Define compare
func (p TimeSlice) Less(i, j int) bool {
    return p[i].Before(p[j]) }

// Define swap over an array
func (p TimeSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i] }

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
    length_map, err := scrape(name)
    if err != nil {
      return
    }
    // Convert map to slice of keys.
    labels := []time.Time{}
    values := []float64{}
    for key, _ := range length_map {
        labels = append(labels, key)
    }
    sort.Sort(TimeSlice(labels))
    for _, label := range labels {
        values = append(values, length_map[label])
    }
    fmt.Println("Done scraping")
    c.HTML(http.StatusOK, "chart.html", gin.H{
			"title": "Chart time!",
      "values": values,
      "labels": labels,
      "max": 10,
		})
  })
	router.Run() // listen and serve on 0.0.0.0:8080
}
