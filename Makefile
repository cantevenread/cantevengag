run:
	go run main.go

build:
	go build  -o bin/cantevengagv2 main.go

clean:
	rm -rf bin

tidy:
	go mod tidy


build-run:
	go build -o bin/learngo main.go && ./bin/cantevengagv2

vendor:
	go mod vendor
