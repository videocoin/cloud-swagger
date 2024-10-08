GOOS?=linux
GOARCH?=amd64

GCP_PROJECT=videocoin-network

NAME=swagger
VERSION=$$(git rev-parse --abbrev-ref HEAD)-$$(git rev-parse --short HEAD)

ENV?=snb

.PHONY: deploy

default: build

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./main.go

deps:
	GO111MODULE=on go mod vendor

docker-build:
	docker build -t gcr.io/${GCP_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	docker push gcr.io/${GCP_PROJECT}/${NAME}:${VERSION}

release: docker-build docker-push

deploy:
	ENV=${ENV} deploy/deploy.sh