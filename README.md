# League-a-Lot
[![Go Report Card](https://goreportcard.com/badge/github.com/schuermannator/go-league)](https://goreportcard.com/report/github.com/schuermannator/go-league)  
Check how much time you (or your friends) spend playing League of Legends!  

Live at: https://league.zvs.io  

## Known Issues

Limited API Key - rate limited and code does not handle well (just returns 404 error after attempting to chart)


## Building

Running locally:  
```bash 
$ go build
$ ./go-league
```

Build image for deployment:
```bash 
$ docker build -t schuermannator/league .
$ docker push schuermannator/league
```

Then restart pod on Kubernetes. (deployment has `ImagePullPolicy` set to `Always`)  

## Misc

Formatted with ```gofmt -s -w```
