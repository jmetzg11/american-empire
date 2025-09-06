# American Empire

Explore techniques used by the United States to assert its global dominace. 

### Prerequisites 
- Go 1.23+
- Node.js 18+ 
- Python 3.12+ with uv
- Air (for Go hot reload): `go install github.com/cosmtrek/air@latest`

### Development Setup
```bash
# Start local PostgreSQL
cd data && docker compose up -d postgres

# Set up Django admin (migrations & superuser)
cd admin && make dev

# Backend with hot reload (new terminal)
air

# Frontend (new terminal)
cd frontend
npm install
npm run dev
```

### Access Points
- **Public App**: http://localhost:5173
- **Django Admin**: http://localhost:8000/admin/
- **Go API**: http://localhost:8080/api/
