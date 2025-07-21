NAME ?= gopu
TAG ?= latest
REPO ?= etzba/${NAME}

all: test up exec down

# docker
up:
	docker-compose down
	docker-compose up -d gopu

down:
	docker-compose down 

cleanup:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor=etzba/gopu:latest --format="{{.ID}}"))

# tests and run
test:
	go test -v ./...

run:
	go run main.go

upload:
	goploader --dir=files/ --url=http://localhost:8080/pics --method=post
	goploader --dir=files/ --url=http://localhost:8080/docs --method=put

# test with etzba 
exec:
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=2
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=4 -r=12 --output=etzba/results/$$(date +%Y%m%d_%H%M%S)_result.json
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=6 -r=24 --output=etzba/results/$$(date +%Y%m%d_%H%M%S)_result.json

# build \ push to dockerhub
.PHONY: docker-build
docker-build:
	docker build -t ${REPO}:${TAG} .

.PHONY: docker-push
docker-push:
	docker push ${REPO}:${TAG}

# install or upgrade helm in kubernetes
install:
	helm install ${NAME} chart/ -n ${NAME} --create-namespace

upgrade:
	helm upgrade --install ${NAME} chart/ -n ${NAME}

