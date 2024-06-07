
run:
	@echo "Running the program..."
	go run .
build:
	@echo "Building the program..."
	SET GOOS=windows
	go build -o go-weather
	@echo "Build complete"
test:
	@echo "Running tests..."
	go test -v ./...
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

