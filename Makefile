lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race ./...

gen-mocks:
	mockgen \
	-destination=test/mocks/provider.go \
	-package=mocks \
	-mock_names=Source=MockSource,SourcesStorage=MockSourcesStorage,\
	ConfigsStorage=MockConfigStorage,Loader=MockLoader,Config=MockConfig \
	github.com/hsqds/conf/internal/provider \
	Source,SourcesStorage,Config,ConfigsStorage,Loader

install-tools:
	go get github.com/golangci/golangci-lint 
	go get github.com/onsi/ginkgo
	go get github.com/golang/mock/mockgen
	go get github.com/golang/mock