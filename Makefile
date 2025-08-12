sourcefiles = $(wildcard **/*.go)

build: $(sourcefiles)
	go build -o HarborMaster ./cmd/HarborMaster

run: build
	NEEDED_HEADER= ./HarborMaster

test:
	go test ./...

clean:
	-rm HarborMaster
