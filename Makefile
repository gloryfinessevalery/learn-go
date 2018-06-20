dep:
	go get -u github.com/golang/dep/cmd/dep

dep-ensure: dep
	dep ensure -v

install:
	go install

start:
	go run main.go