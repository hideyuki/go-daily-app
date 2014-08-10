# go-dialy-app

====

GO Dialy Application

<img height="200" src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" />

# Go App
## Installation

```
$ go get github.com/kr/godep
$ godep restore
```

## Run

```
$ go get github.com/codegangsta/gin
$ gin run server.go   # go run server.go
```

## Test

- [Ginkgo](http://github.com/onsi/ginkgo): BDD Testing Framework
- [Gomega](http://github.com/onsi/gomega): Matcher library

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
$ go test
```

---

# Docker

This repository connect the [Docker Hub repository](https://registry.hub.docker.com/u/hideyuki/go-daily/) for automated build of docker image.

## Build

```
$ docker build --rm -t hideyuki/go-daily:0.1 ./
``` 

With no-cache flag

```
$ docker build --rm --no-cache -t hideyuki/go-daily:0.1 ./
``` 

## Run

```
$ docker run -i -t -p 13000:3000 -p 10022:22 --name daily hideyuki/go-daily:0.1 /bin/bash
[martini] listening on :3000 (development)
```

As daemon

```
$ docker run -d -p 13000:3000 -p 10022:22 --name daily hideyuki/go-daily:0.1
```

## Push to Docker Hub

```
$ docker login    # if you need
$ docker push hideyuki/go-daily:0.1
```

