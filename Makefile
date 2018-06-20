dep:
	@go get -u github.com/golang/dep/cmd/dep

dep-ensure: dep
	@dep ensure -v

build:
	@go build

start:
	@./learn-go