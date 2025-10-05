# American Empire

Explore techniques used by the United States to assert its global dominance.

### Prerequisites
- Go 1.23+
- Python 3.12+ with uv
- Air (for Go hot reload): `go install github.com/cosmtrek/air@latest`
- Docker & Docker Compose

### Development Setup
```bash
# Quick start - runs database, Django admin, and Go web server
make run

# Or manually:
# 1. Start local PostgreSQL
cd data && docker compose up -d postgres

# 2. Start Django admin with migrations & seeding
cd admin && make dev

# 3. Start Go web server with hot reload
cd web && air
```

### Access Points
- **Web App**: http://localhost:8080
- **Django Admin**: http://localhost:8000/admin/

### Architecture
- **Go Web Server** (`web/`): HTML templates with server-side rendering
- **Django Admin** (`admin/`): Content management interface
- **PostgreSQL**: Shared database for both Go and Django
- **Templates**: Go html/template in `web/ui/html/`
- **Static Assets**: CSS/JS in `web/ui/static/`
