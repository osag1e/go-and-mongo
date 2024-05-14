build-api:
	@go build -o bin/api ./cmd/api/

run: build-api
	@./bin/api


clean: 
	@rm -rf bin

