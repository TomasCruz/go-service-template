clean_build: clean build

clean:
	go clean

build:
	go build -o bin/server

run:
	bin/server

.PHONY: test end2end
test:
	go test -v -count=1 ./...

end2end:
	go test -v -count=1 -tags end2end ./...
