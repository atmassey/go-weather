DIR_TO_REMOVE := bin

run:
	@echo "Running the program..."
	go run .
build:
	@echo "Building the program..."
	go build -o bin/ .
test:
	@echo "Running tests..."
	go test -v ./...
clean:
	@echo "Cleaning up..."
	rmdir /s /q $(DIR_TO_REMOVE)

