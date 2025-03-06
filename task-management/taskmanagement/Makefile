run:
	go run ./src/main.go

build:
	go build -o ./bin/projectmanagement ./src/main.go

test:
	go test ./src/...

coverage:
	go test -coverprofile=coverage.out ./src/...
	go tool cover -html=coverage.out

lint:
	golangci-lint run

docker-run:
	docker run -p 8080:8080 go-api

docker-build:
	docker build -t go-api .

docker-push:
	docker tag go-api:latest
	docker push go-api:latest

docker-compose-up:
	docker-compose up

docker-compose-down:
	docker-compose down

