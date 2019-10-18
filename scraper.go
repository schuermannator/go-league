package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"crypto/tls"
	"os"
	"sync"
	"time"
)

func getAccountID(summoner string) (string, error) {
	apiKey := url.QueryEscape(os.Getenv("RIOTAPIKEY"))
	endpt := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summoner, apiKey)
	//log.Print(endpt)
	// Build the request
	req, err := http.NewRequest("GET", endpt, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "", err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request fatal - Do: ", err)
		return "", err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	//var record Numverify
	//var data string

	// Use json.Decode for reading streams of JSON data
	// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	// 	log.Println(err)
	// }
	body, err := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	//log.Print(data["accountId"])

	return data["accountId"].(string), nil
}

func getMatches(id string) ([]int64, error) {
	apiKey := url.QueryEscape(os.Getenv("RIOTAPIKEY"))
	endpt := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s?api_key=%s", id, apiKey)
	//log.Print(endpt)
	// Build the request
	req, err := http.NewRequest("GET", endpt, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	//log.Print(data)
	var list = make([]int64, 0)
	matchList := data["matches"].([]interface{})
	for _, match := range matchList {
		matchMap := match.(map[string]interface{})
		list = append(list, int64(matchMap["gameId"].(float64)))
	}
	//log.Print(list)
	return list, nil
}

func getMatchTimes(matchID int64) (time.Time, float64, error) {
	apiKey := url.QueryEscape(os.Getenv("RIOTAPIKEY"))
	endpt := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matches/%d?api_key=%s", matchID, apiKey)
	//log.Print(endpt)
	// Build the request
	req, err := http.NewRequest("GET", endpt, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return time.Time{}, 0, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return time.Time{}, 0, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Do: ", err)
		return time.Time{}, 0, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	if data["gameCreation"] == nil {
		panic(data)
	}

	creation := time.Unix(int64(data["gameCreation"].(float64)/1000), 0)
	duration := data["gameDuration"].(float64) / 3600.
	rounded := time.Date(creation.Year(), creation.Month(), creation.Day(), 0, 0, 0, 0, creation.Location())
	return rounded, duration, nil
}

func scrape(name string, length int) (map[time.Time]float64, error) {
	id, err := getAccountID(name)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
    var mu = &sync.Mutex{}

	var matchList []int64
	matchList, err = getMatches(id)

	lengthMap := make(map[time.Time]float64)
	for i, match := range matchList {
		if i > length {
			break
		}
		wg.Add(1)
		log.Print(match)
		go func(match int64) {
			defer wg.Done()
			create, dur, _ := getMatchTimes(match)
            mu.Lock()
			if val, ok := lengthMap[create]; ok {
				// already in map
				lengthMap[create] = val + dur
			} else {
		 		// not in map
				lengthMap[create] = dur
		 	}
            mu.Unlock()
			//if er != nil {
			//	return nil, err
			//}
			log.Print(create, dur)
		}(match)
	}
	wg.Wait()
	return lengthMap, nil
}
