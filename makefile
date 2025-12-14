PROJECT?=filewatch
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
APP?=app
PORT?=11002
ImageName?=filewatchImageName
ContainerName?=filewatchContainer
DBSERVER?=MySQLx
MKFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURDIR := $(dir $(MKFILE))

cleanDocker:
	sh clean.sh

clean:
	rm -f ${APP}

build:
	GOOS=linux GOARCH=amd64 go build -tags netgo \
	-ldflags "-s -w -X version.Release=${RELEASE} \
	-X version.Commit=${COMMIT} \
	-X version.BuildTime=${BUILD_TIME}" \
	-o ${APP}

docker: build
	docker build -t ${ImageName} .
	rm -f ${APP}
	docker images

runx: docker cleanDocker
	docker run -d --name ${ContainerName} \
	-v /etc/localtime:/etc/localtime:ro \
	-v ${CURDIR}www:/app/www  \
	--restart=always \
	-p ${PORT}:80 \
	--link ${DBSERVER} \
	--env-file ${CURDIR}envfile \
	${ImageName}

run:
	go run fwatcher.go

stop:
	docker stop ${ContainerName}

log:
	 docker logs -f -t --tail 20 ${ContainerName}
rm:
	docker rm ${ContainerName}
s:
	git push -u origin master
