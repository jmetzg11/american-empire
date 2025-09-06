.PHONY: clean clean-db test run


clean:
	rm -rf ./tmp
	mkdir -p ./tmp
	rm -f app

test:
	go test ./...

run:
	go run main.go
