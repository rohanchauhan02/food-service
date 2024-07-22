BINARY=engine

# run golang service
start:
	go run app/main.go

# run dev docker-compose
run:
	docker-compose up -d --build

# stop docker compose
stop:
	docker-compose down

# build docker image
docker:
	docker build -t food-service:latest .

# run golang test
test:
	go test ./...

# clean binary build of this go service
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

vendor:
	go mod vendor

engine:
	go build -o engine app/main.go

.PHONY: engine start run stop docker test vendor  clean

