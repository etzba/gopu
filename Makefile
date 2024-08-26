TAG ?= latest
REPO ?= etzba/gopu

all: test up exec down

# unit tests
test:
	go test -v ./...

run:
	go run main.go

up:
	docker-compose down
	docker-compose up -d gopu

exec:
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=2
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=4 -r=12 --output=etzba/results/$$(date +%Y%m%d_%H%M%S)_result.json
	etz api --auth=etzba/secret.yaml --exec=etzba/executions.yaml -d=3s -w=6 -r=24 --output=etzba/results/$$(date +%Y%m%d_%H%M%S)_result.json

down:
	docker-compose down 

cleanup:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor=etzba/gopu:latest --format="{{.ID}}"))

# build image and push to dockerhub
.PHONY: docker-build
docker-build:
	docker build -t ${REPO}:${TAG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${REPO}:${TAG}

helm:
	helm install gopu chart/ -n gopu --create-namespace
