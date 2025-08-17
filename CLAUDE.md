# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Backend (Go)
- `make run` - Start the Go backend server on port 8080
- `make test` - Run Go tests
- `make clean` - Clean build artifacts and recreate tmp directory
- `make clean-db` - Remove local SQLite database file
- `go run cmd/migrate/main.go` - Run database migrations
- `go run cmd/seed/main.go` - Seed database with initial data

### Frontend (SvelteKit)
- `cd frontend && npm run dev` - Start development server with Vite
- `cd frontend && npm run build` - Build for production
- `cd frontend && npm run preview` - Preview production build
- `cd frontend && npm run lint` - Run Prettier code formatting check
- `cd frontend && npm run format` - Format code with Prettier

## Architecture Overview

This is a full-stack web application for tracking historical events with an admin interface for content management.

### Backend (Go + Gin)
- **Main server**: `main.go` - Gin web server on port 8080
- **Database**: GORM with SQLite (development) / PostgreSQL (production)
- **Authentication**: JWT-based admin authentication with bcrypt password hashing
- **File storage**: Local filesystem (development) / Supabase storage (production)
- **Rate limiting**: Login attempts and API calls are rate-limited

**Key Models** (`backend/models/models.go`):
- `Event` - Main content entity with title, date, country, description
- `Source` - Links/references for events
- `Media` - Photos/videos attached to events
- `Tag` - Categorization system with many-to-many relationship

**API Structure** (`backend/routes/routes.go`):
- Public endpoints: `/api/` (events), `/api/event`, `/api/contribute`, `/api/tags`
- Admin endpoints: All prefixed with `/api/admin-` and require authentication
- Authentication: `/api/login`, `/api/auth-me`

### Frontend (SvelteKit + Tailwind)
- **Framework**: SvelteKit with Vite build system
- **Styling**: Tailwind CSS v4
- **API Communication**: Fetch-based API client in `src/lib/api.js`
- **Routing**: File-based routing with SvelteKit conventions

**Key Routes**:
- `/` - Public event listing
- `/event/[id]` - Individual event details
- `/contribute` - Public event submission form
- `/admin` - Admin dashboard requiring authentication
- `/admin/event/[id]` - Admin event editing interface

**Development Proxy**: Vite proxies `/api` requests to `localhost:8080` for the Go backend

### Environment Configuration
- **Development**: SQLite database, local file storage, GIN_MODE not set to "release"
- **Production**: PostgreSQL via DATABASE_URL, Supabase storage via SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY
- **Authentication**: ADMIN_PASSWORD environment variable (gets hashed at startup)

### Database Schema
Events are the central entity with:
- One-to-many relationships with Sources and Media
- Many-to-many relationship with Tags via junction table
- `Active` timestamp field for approval workflow
- `Flagged` boolean for content moderation

## Development Workflow

1. Start backend: `make run`
2. Start frontend: `cd frontend && npm run dev`
3. Access app at `http://localhost:5173` (frontend dev server)
4. API requests proxy through to `http://localhost:8080`

The application supports both local development and cloud deployment on Fly.io with separate Dockerfiles for frontend and backend.