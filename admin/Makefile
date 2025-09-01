.PHONY: dev prod stop shell migrate

dev:
	DATABASE_PATH="../data/american-empire.db" DEBUG=True uv run python manage.py migrate
	DATABASE_PATH="../data/american-empire.db" DEBUG=True uv run python manage.py shell -c "from django.contrib.auth.models import User; import os; User.objects.create_superuser(os.environ.get('ADMIN_USERNAME', 'admin'), '', os.environ.get('ADMIN_PASSWORD', 'admin')) if not User.objects.filter(username=os.environ.get('ADMIN_USERNAME', 'admin')).exists() else print('Superuser already exists')"
	DATABASE_PATH="../data/american-empire.db" DEBUG=True uv run python manage.py runserver

prod:
	ENVIRONMENT=production uv run python manage.py migrate
	ENVIRONMENT=production uv run python manage.py shell -c "from django.contrib.auth.models import User; import os; User.objects.create_superuser(os.environ.get('ADMIN_USERNAME', 'admin'), '', os.environ.get('ADMIN_PASSWORD', 'admin')) if not User.objects.filter(username=os.environ.get('ADMIN_USERNAME', 'admin')).exists() else print('Superuser already exists')"
	ENVIRONMENT=production uv run python manage.py runserver

migrate:
	DATABASE_PATH="../data/american-empire.db" DEBUG=True uv run python manage.py migrate

stop:
	pkill -f "python manage.py runserver" || true

shell:
	uv run python manage.py shell
