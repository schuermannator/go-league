package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"sync"
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

func apiCounter(apiReqs *[]time.Time) {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			// clear requests before current time
			then := time.Now().Add(-10 * time.Minute)
			//then := time.Now().Add(-2 * time.Minute)
			// deleteIndex := -1
			for i, t := range *apiReqs {
				if t.After(then) {
					// deleteIndex = i - 1
					if i > 0 {
						*apiReqs = (*apiReqs)[i-1:]
					}
					break
				}
			}

		}
	}()
}

func main() {

	// atomic counter for number of API requests
	var apiReqs []time.Time
	var mu = &sync.Mutex{}

	go apiCounter(&apiReqs)

	log.Print("Starting server...")
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "League-a-Lot",
		})
	})
	router.POST("/", func(c *gin.Context) {
		name := c.PostForm("text")
		re := regexp.MustCompile("[0-9]+")
		historyLength, err := strconv.Atoi(re.FindAllString(c.PostForm("dropdown"), 1)[0])
		log.Println(historyLength)
		if err != nil {
			return
		}
		lengthMap, err := scrape(name, historyLength, &apiReqs, mu)
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
		c.HTML(http.StatusOK, "chart.html", gin.H{
			"title":  "League-a-Lot: " + name,
			"values": values_str,
			"labels": labels,
			"max":    10,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
