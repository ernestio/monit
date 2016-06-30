install:
	go install -v

deps: dev-deps
	go get github.com/r3labs/sse
	go get gopkg.in/redis.v3
	go get github.com/nats-io/nats
	go get github.com/dgrijalva/jwt-go

dev-deps:
	go get github.com/smartystreets/goconvey/convey
	go get github.com/golang/lint/golint

cover:
	go test -v ./... --cover

test:
	go test -v -race ./...

build:
	go build -v ./...

lint:
	golint ./...
	go vet ./...
