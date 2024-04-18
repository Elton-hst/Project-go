GO_VERSION=1.22.1
CMD=wsl

# generates swagger api docs
docs:
	swag init

run-tests:
	go clean -cache
	go test -v ./...

build:
	go build -o bin/main main.go

run:
	go run main.go

# build file Dockerfile
build: 
	${CMD} docker build

# uploading all docker images
up-docker:
	${CMD} docker start postgres
	${CMD} docker start redis
	${CMD} docker start prometheus
	${CMD} docker start grafana

# pause all docker iamges
pause-docker:
	${CMD} docker pause postgres
	${CMD} docker pause redis
	${CMD} docker pause prometheus
	${CMD} docker pause grafana

# stop all docker iamges
stop-docker:
	${CMD} docker stop postgres
	${CMD} docker stop redis
	${CMD} docker stop prometheus
	${CMD} docker stop grafana


