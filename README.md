# simple-go-expmnt

```sh
docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang go get
docker run -it --rm -v $PWD:/go/src/myapp -w /go/src/myapp -e "GOPATH=/go/src/myapp/vendor:/go" golang go build
docker run -it --rm -v $PWD:/app -w /app -p 8080:8080 golang ./myapp
```

