sourcefiles = $(wildcard **/*.go)

build: $(sourcefiles)
	go build -o HarborMaster ./cmd/HarborMaster

run: build
	./HarborMaster

test:
	go test ./...

clean:
	-rm HarborMaster