# American Empire

Explore techniques used by the United States to assert its global dominace. 

### Prerequisites 
- Go 1.23+
- Node.js 18+ 
- Python 3.12+ with uv
- Air (for Go hot reload): `go install github.com/cosmtrek/air@latest`

### Development Setup

#### Public Application
```bash
# Backend with hot reload
go run cmd/migrate/main.go 
go run cmd/seed/main.go
air

# Frontend (new terminal)
cd frontend
npm install
npm run dev
```

#### Content Management (Local Only)
```bash
# Django admin interface
cd admin
make dev    # Development (SQLite database)
# or
make prod   # Production (PostgreSQL database)
```

Access the Django admin at `http://localhost:8000/admin/` with credentials from `.env` file.
