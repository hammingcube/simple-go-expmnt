# simple-go-expmnt

```sh
docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang go get
docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang go build
docker run -it --rm -v $PWD:/app -w /app -p 8080:8080 golang ./myapp
```

```sh
GOPATH=/Users/madhavjha/src/github.com/maddyonline/simple-go-expmnt/vendor go get -d .
GOPATH=/Users/madhavjha/src/github.com/maddyonline/simple-go-expmnt/vendor go build  -o bin/runit .
```

```sh

docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang:1.4.2-cross go get -d .

docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang:1.4.2-cross sh -c '
for GOOS in darwin linux; do
  for GOARCH in 386 amd64; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -o bin/myapp-$GOOS-$GOARCH
  done
done
'

docker run -it --rm -v $PWD:/app -w /app -p 8080:8080 golang bin/myapp-linux-386
docker run -it --rm -v $PWD:/app -w /app -p 8080:8080 golang bin/myapp-darwin-386
```