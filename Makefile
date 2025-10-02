.PHONY: clean clean-db test run kill-port start-db stop


clean:
	rm -rf ./tmp
	mkdir -p ./tmp
	rm -f app

test:
	go test ./...

start-db:
	cd data && docker-compose up -d postgres && echo "Waiting for PostgreSQL to be ready..." && until docker-compose exec postgres pg_isready -U admin -d american_empire; do sleep 1; done && echo "PostgreSQL is ready!"

run: start-db
	@echo "Starting Django admin..."
	cd admin && DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py migrate && \
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py shell -c "from django.contrib.auth.models import User; import os; User.objects.create_superuser(os.environ.get('ADMIN_USERNAME', 'admin'), '', os.environ.get('ADMIN_PASSWORD', 'admin')) if not User.objects.filter(username=os.environ.get('ADMIN_USERNAME', 'admin')).exists() else print('Superuser already exists')" && \
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py seed_db && \
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py runserver &
	@echo "Starting Go backend..."
	cd web && air


kill-port:
	@if lsof -ti:8080 > /dev/null 2>&1; then \
		lsof -ti:8080 | xargs kill -9 && echo "Killed process on port 8080"; \
	else \
		echo "No process using port 8080"; \
	fi

restart: clean
	@pkill -f "air" || true
	cd web && air

stop:
	@echo "Stopping Django admin..."
	-@pkill -f "manage.py runserver" 2>/dev/null || true
	@echo "Stopping Go backend..."
	-@$(MAKE) kill-port
	@echo "Stopping database..."
	-@cd data && docker-compose down -v
