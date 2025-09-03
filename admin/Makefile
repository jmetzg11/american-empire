.PHONY: dev prod stop shell migrate start-db stop-db

start-db:
	cd ../data && docker-compose up -d postgres
	@echo "Waiting for PostgreSQL to be ready..."
	@until cd ../data && docker-compose exec postgres pg_isready -U admin -d american_empire; do sleep 1; done
	@echo "PostgreSQL is ready!"

stop-db:
	cd ../data && docker-compose down

dev: start-db
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py migrate
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py shell -c "from django.contrib.auth.models import User; import os; User.objects.create_superuser(os.environ.get('ADMIN_USERNAME', 'admin'), '', os.environ.get('ADMIN_PASSWORD', 'admin')) if not User.objects.filter(username=os.environ.get('ADMIN_USERNAME', 'admin')).exists() else print('Superuser already exists')"
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py seed_db
	DB_HOST=localhost DB_PORT=5432 DB_NAME=american_empire DB_USER=admin DB_PASSWORD=admin DEBUG=True uv run python manage.py runserver

prod:
	ENVIRONMENT=production uv run python manage.py migrate
	ENVIRONMENT=production uv run python manage.py shell -c "from django.contrib.auth.models import User; import os; User.objects.create_superuser(os.environ.get('ADMIN_USERNAME', 'admin'), '', os.environ.get('ADMIN_PASSWORD', 'admin')) if not User.objects.filter(username=os.environ.get('ADMIN_USERNAME', 'admin')).exists() else print('Superuser already exists')"
	ENVIRONMENT=production uv run python manage.py runserver

migrate:
	DATABASE_PATH="../data/american-empire.db" DEBUG=True uv run python manage.py migrate

stop: stop-db
	pkill -f "python manage.py runserver" || true

shell:
	uv run python manage.py shell
