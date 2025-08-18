.PHONY: clean clean-db test run


clean:
	rm -rf ./tmp
	mkdir -p ./tmp
	rm -f app

clean-db:
	rm -f ./data/american-empire.db

migrate:
	go run cmd/migrate/main.go 

seed: 
	go run cmd/seed/main.go

test:
	go test ./...

run:
	go run main.go
