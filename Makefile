install:
	go install -v

deps:
	go get github.com/r3labs/sse
	go get gopkg.in/redis.v3
	go get github.com/nats-io/nats
	go get github.com/dgrijalva/jwt-go
	go get github.com/r3labs/pattern
	go get github.com/ernestio/ernest-config-client

dev-deps: deps
	go get github.com/smartystreets/goconvey/convey
	go get github.com/alecthomas/gometalinter
	gometalinter --install

cover:
	go test -v ./... --cover

test:
	go test -v -race ./...

build:
	go build -v ./...

lint:
	gometalinter --config .linter.conf

