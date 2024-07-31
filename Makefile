MAIN=./cmd/main.go

run:
	go run $(MAIN)

# Clean target: clean any build artifacts (if needed)
clean:
	go clean

.PHONY: run clean
