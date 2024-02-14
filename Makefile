build:
	docker-compose up --build

test:
	go test -v ./...