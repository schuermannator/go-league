# League-a-Lot
[![Go Report Card](https://goreportcard.com/badge/github.com/schuermannator/league-a-lot)](https://goreportcard.com/report/github.com/schuermannator/league-a-lot)  
Check how much time you (or your friends) spend playing League of Legends!  

Live at: https://lol.zvs.io  

## Running
```
# docker build -t league-a-lot .
# docker run -p 8080:8080 --env GIN_MODE=release --env RIOTAPIKEY=*** league-a-lot
```

## Building

Running locally:  
```bash 
$ go build && ./league-a-lot
```

## Todo
- [ ] Timing validation/tests
- [ ] Integration test to ping final site and check for 200 OK
- [ ] Visual feedback for waiting in line/rate limiting
- [ ] Expand to better app that can 'track' certain ID's and continuously scrape, etc.
- [ ] Ceiling for queries at 100 
