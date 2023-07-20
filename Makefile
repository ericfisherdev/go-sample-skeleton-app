BINARY_NAME=finishlineApp

build:
	@go mod vendor
	@echo "Building FinishLine..."
	@go build -o ./tmp/${BINARY_NAME} .
	@echo "FinishLine built!"

run: build
	@echo "Staring FinishLine..."
	@./tmp/${BINARY_NAME} &
	@echo "FinishLine started!"

clean:
	@echo "Cleaning..."./
	@go clean
	@rm ./tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping the front end..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped FinishLine"

restart: stop start