sourcefiles = $(wildcard **/*.go)

build: $(sourcefiles)
	go build -o HarborMaster ./cmd/HarborMaster

run: build
	./HarborMaster -url https://registry.max-heidinger.de -username_file ./registry_username -password_file ./registry_password

test:
	go test ./...

clean:
	-rm HarborMaster