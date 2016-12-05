test:
	go build
	gocov test ./... | gocov report

lint:
	go lint ./...

run:
	go build
	./bloodlines

get-deps:
	godeps restore
	go get github.com/axw/gocov