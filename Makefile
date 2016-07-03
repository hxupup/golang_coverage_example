GOFILES = $(shell find . -type f -name '*.go')

.PHONY: buildtest.test
buildtest.test: $(GOFILES)
	go test -c -covermode=count -coverpkg ./...

.PHONY: system-tests
system-tests: buildtest.test
	./test.test -systemTest -test.coverprofile ./coverage.cov

.PHONY: testsuite
testsuite: system-tests
	go tool cover -html=./coverage.cov -o ./coverage.html

