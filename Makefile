DIR_TO_REMOVE := bin
APP = go-weather

run:
	@echo "Running the program..."
	go run .
build:
	@echo "Building the program..."
	SET GOOS=linux
	go build -o $(APP)
	SET GOOS=windows
	@echo "Build complete"
test:
	@echo "Running tests..."
	go test -v ./...
clean:
	@echo "Cleaning up..."
	rmdir /s /q $(DIR_TO_REMOVE)
up_build:
	@echo "Stopping docker images if running"
	docker-compose down
	@echo "Building docker images"
	docker-compose up --build
	@echo "Docker images built and started"
up:
	@echo "Stopping docker images if running"
	docker-compose down
	@echo "Starting docker images"
	docker-compose up
	@echo "Docker images started"
down:
	@echo "Stopping docker images"
	docker-compose down
	@echo "Docker images stopped"

