hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	clear && go run main.go