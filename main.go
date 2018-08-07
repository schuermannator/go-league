package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
  "encoding/json"
  "io/ioutil"
)

func get_account_id(summoner string) (int, error) {
  api_key := url.QueryEscape("RGAPI-3c40b941-1cbe-4bc6-8b0a-c2bb3d69d118")
  endpt := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s", summoner, api_key)

	// Build the request
	req, err := http.NewRequest("GET", endpt, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return -1, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

  // Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return -1, err
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

  log.Print(data["accountId"])

  return int(data["accountId"].(float64)), nil
}

func get_matches(id int) (map[string]interface{}, error) {
  api_key := url.QueryEscape("RGAPI-3c40b941-1cbe-4bc6-8b0a-c2bb3d69d118")
  endpt := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v3/matchlists/by-account/%d?api_key=%s", id, api_key)
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
	client := &http.Client{}

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

  return data, nil
}

func main() {
	id, err := get_account_id("fuqqboi")
	if err != nil {
    return
  }

  var data map[string]interface{}
  data, err = get_matches(id)

  log.Print(data)

}
