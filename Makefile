.PHONY: clean clean-db test run kill-port


clean:
	rm -rf ./tmp
	mkdir -p ./tmp
	rm -f app

test:
	go test ./...

run:
	cd web && air

kill-port:
	@if lsof -ti:8080 > /dev/null 2>&1; then \
		lsof -ti:8080 | xargs kill -9 && echo "Killed process on port 8080"; \
	else \
		echo "No process using port 8080"; \
	fi
