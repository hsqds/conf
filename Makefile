lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race -cover -coverprofile=coverage.out -outputdir=./test ./...

gen-mocks:
	mockgen \
	-destination=test/mocks/conf.go \
	-package=mocks \
	-mock_names=Source=MockSource,SourcesStorage=MockSourcesStorage,\
	ConfigsStorage=MockConfigStorage,Loader=MockLoader,Config=MockConfig \
	github.com/hsqds/conf \
	Source,SourcesStorage,Config,ConfigsStorage,Loader

coverage:
	go tool cover -html=./test/coverage.out

install-tools:
	go get github.com/golangci/golangci-lint 
	go get github.com/onsi/ginkgo
	go get github.com/golang/mock/mockgen
	go get github.com/golang/mock