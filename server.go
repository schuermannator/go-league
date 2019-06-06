package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// TimeSlice Define us a type so we can sort it
type TimeSlice []time.Time

// Forward request for length
func (p TimeSlice) Len() int {
	return len(p)
}

// Define compare
func (p TimeSlice) Less(i, j int) bool {
	return p[i].Before(p[j])
}

// Define swap over an array
func (p TimeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
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
		re := regexp.MustCompile("[0-9]+")
		historyLength, err := strconv.Atoi(re.FindAllString(c.PostForm("dropdown"), 1)[0])
		if err != nil {
			return
		}
		lengthMap, err := scrape(name, historyLength)
		if err != nil {
			return
		}
		// Convert map to slice of keys.
		times := []time.Time{}
		values := []float64{}
		labels := []string{}
		values_str := []string{}
		for key := range lengthMap {
			times = append(times, key)
		}
		sort.Sort(TimeSlice(times))
		for _, t := range times {
			values = append(values, lengthMap[t])
		}
		for t := range times {
			labels = append(labels, times[t].Format("01-02-2006"))
		}
		for v := range values {
			values_str = append(values_str, strconv.FormatFloat(values[v], 'f', 2, 64))
		}
		fmt.Println("Done scraping")
		c.HTML(http.StatusOK, "chart.html", gin.H{
			"title":  "Chart time!",
			"values": values_str,
			"labels": labels,
			"max":    10,
		})
	})
    router.Run() // listen and serve on 0.0.0.0:8080
}
