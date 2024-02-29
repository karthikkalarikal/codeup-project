FRONT_END_BINARY=frontApp
BROKER_BINARY=gatewayApp
AUTH_BINARY=authApp
PROBLEM_BINARY=testApp
## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_gateway build_auth build_problem_service
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_gateway:
	@echo "Building broker binary..."
	cd api-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth : builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_problem_service:
	@echo "Building service binary..."
	cd problem-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PROBLEM_BINARY} ./cmd/api
	@echo "Done!"



## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd front-end && ./${FRONT_END_BINARY} &																																																																																																																

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"