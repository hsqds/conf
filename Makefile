lint:
	golangci-lint run ./...

gen-mocks:
	mockgen \
	-destination=test/mocks/source.go \
	-package=mocks \
	-mock_names=Source=MockSource \
	github.com/hate-squids/config-provider/provider \
	Source

install-tools:
	go get github.com/golangci/golangci-lint 
	go get github.com/onsi/ginkgo
	go get github.com/golang/mock/mockgen
	go get github.com/golang/mock