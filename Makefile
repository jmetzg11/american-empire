.PHONY: clean clean-db test run


clean:
	rm -rf ./tmp
	mkdir -p ./tmp
	rm -f app

clean-db:
	rm -f ./data/good-guys.db


test:
	go test ./...

run:
	go run main.go
