run:
	go mod tidy
	go run ./cmd/app/main.go --release

build:
	go mod tidy
	go build -o app ./cmd/app/main.go
	./app --release

clear:
	rm -f app
