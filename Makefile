.PHONY: bin

bin:
	@go build -o bin/interpreter ./cmd/interpreter
	@echo "Binary generated successfully!"
