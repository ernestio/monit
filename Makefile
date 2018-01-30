install:
	go install -v

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

dev-deps: deps
	go get github.com/smartystreets/goconvey/convey
	go get github.com/alecthomas/gometalinter
	gometalinter --install

cover:
	go test -v $(go list ./... | grep -v /vendor/) --cover

test:
	go test --cover -v $(go list ./... | grep -v /vendor/)

build:
	go build -v ./...

lint:
	gometalinter --config .linter.conf
