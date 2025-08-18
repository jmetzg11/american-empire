# American Empire

Explore techniques used by the United States to assert its global dominace. 

### Prerequisites 
- Go 1.23+
- Node.js 18+ 
- Air (for Go hot reload): `go install github.com/cosmtrek/air@latest`

### Development Setup
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
