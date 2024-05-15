
.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -C cmd/app/. -o ../../bin/app.exe


.PHONY: run
run: build
	@bin/app

.PHONY: test
test:
	@go test -v ./... -count=1

