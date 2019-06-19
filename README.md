# League-a-Lot
[![Go Report Card](https://goreportcard.com/badge/github.com/schuermannator/go-league)](https://goreportcard.com/report/github.com/schuermannator/go-league)  
Check how much time you (or your friends) spend playing League of Legends!  

Live at: https://lol.zvs.io  

## Building

Running locally:  
```bash 
$ make
$ ./league-a-lot
```

Deploy: (requires kubectl context)
```bash 
$ make docker
$ kubectl apply -f .
```

Or just restart pod on Kubernetes. (deployment has `ImagePullPolicy` set to `Always`)  

## Misc

Formatted with ```gofmt -s -w```
