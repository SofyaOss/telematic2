build:
	sudo docker compose build

run:
	sudo docker compose up app

go-test:
	go test ./...

lint:
	@go get golang.org/x/lint/golint
	@go install golang.org/x/lint/golint
	golint ./...
