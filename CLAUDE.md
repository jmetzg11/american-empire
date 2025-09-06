# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Backend (Go)
- `air` - Start the Go backend server with hot reload on port 8080 (development)
- `make run` - Start the Go backend server on port 8080 (production)
- `make test` - Run Go tests
- `make clean` - Clean build artifacts and recreate tmp directory

### Frontend (SvelteKit)
- `cd frontend && npm run dev` - Start development server with Vite
- `cd frontend && npm run build` - Build for production
- `cd frontend && npm run preview` - Preview production build
- `cd frontend && npm run lint` - Run Prettier code formatting check
- `cd frontend && npm run format` - Format code with Prettier

### Django Admin (Content Management)
- `cd admin && make dev` - Start Django admin on local Docker database (development)
- `cd admin && make prod` - Start Django admin on Supabase database (production)
- `cd admin && make stop` - Stop Django server
- `cd admin && make shell` - Access Django shell
- `cd admin && make migrate` - Run Django migrations (dev database only)

## Architecture Overview

This is a full-stack web application for tracking historical events with a clean separation between public API and admin functionality.

### Backend (Go + Gin) - Public API Only
- **Main server**: `main.go` - Gin web server on port 8080
- **Database**: GORM with PostgreSQL (always uses DATABASE_URL)
- **File storage**: Local filesystem (development) / Supabase storage (production)
- **Purpose**: Serves public API and handles public contributions

**Key Models** (`backend/models/models.go`):
- `Event` - Main content entity with title, date, country, description
- `Source` - Links/references for events
- `Media` - Photos/videos attached to events
- `Tag` - Categorization system with many-to-many relationship
- `Book` - Book recommendations with many-to-many relationship to events

**API Structure** (`backend/routes/routes.go`):
- `GET /api/` - List all active events
- `POST /api/event` - Get single event details
- `POST /api/contribute` - Submit new event (public contributions)
- `GET /api/tags` - List all tags
- `GET /api/book/:id` - Get book details

**Photo Storage**:
- Development: Local files in `data/photos/`, served at `/photos`
- Production: Supabase storage, served directly from CDN
- Uses `InitSupabase()` and `saveUploadedPhoto()` for public contributions

### Frontend (SvelteKit + Tailwind) - Public Interface Only
- **Framework**: SvelteKit with Vite build system
- **Styling**: Tailwind CSS v4
- **API Communication**: Direct fetch calls to Go backend
- **Routing**: File-based routing with SvelteKit conventions

**Key Routes**:
- `/` - Public event listing with filtering
- `/event/[id]` - Individual event details
- `/contribute` - Public event submission form

**Development Proxy**: Vite proxies `/api` requests to `localhost:8080` for the Go backend

### Django Admin (Content Management)
- **Framework**: Django with uv package management
- **Purpose**: Complete admin interface for managing all content
- **Database**: Shares the same PostgreSQL database as Go backend
- **Authentication**: Uses ADMIN_USERNAME/ADMIN_PASSWORD from .env file

**Admin Interface Features**:
- Event management (create, edit, approve, flag)
- Tag management with event assignments
- Source management (links/references)
- Media management (photos, YouTube videos)
- Book management with event associations

**Models** (`admin/core/models.py`):
- Mirrors Go models exactly: Event, Source, Media, Tag, Book
- Uses same database tables as Go application
- Provides Django admin interface with inlines and filtering

### Environment Configuration
- **Development**: 
  - Go: Uses local Docker PostgreSQL (`postgresql://admin:admin@localhost:5432/american_empire`)
  - Django: Uses local Docker PostgreSQL
  - Photos: Stored locally in `data/photos/`
- **Production**: 
  - Both Go and Django: Use Supabase PostgreSQL via DATABASE_URL
  - Photos: Stored in Supabase storage, served from CDN

**Key Environment Variables**:
- `DATABASE_URL` - PostgreSQL connection (required for both Go and Django)
- `ADMIN_USERNAME/ADMIN_PASSWORD` - Django admin credentials
- `SUPABASE_URL/SUPABASE_SERVICE_ROLE_KEY` - For photo storage (production only)
- `GMAIL_USER/GMAIL_PASS` - For contribution notifications

### Database Schema
Events are the central entity with:
- One-to-many relationships with Sources and Media
- Many-to-many relationship with Tags via junction table `event_tags`
- Many-to-many relationship with Books via junction table `book_events`
- `Active` timestamp field for approval workflow (NULL = not approved)
- `Flagged` boolean for content moderation

## Development Workflow

### Full Development Setup
1. Start local PostgreSQL: `cd data && docker compose up -d postgres`
2. Start Django admin: `cd admin && make dev`
3. Start Go backend: `make run` or `air`
4. Start frontend: `cd frontend && npm run dev`
5. Access public app at `http://localhost:5173`
6. Access admin at `http://localhost:8000/admin/`

### Content Management Workflow
1. Django admin is the primary interface for all content management
2. Go backend only handles public API and contributions
3. Public contributions are submitted via Go API and managed via Django admin
4. All CRUD operations for events, tags, sources, media happen in Django admin

### Production Deployment
- Go backend serves public API only
- Frontend is static SvelteKit build
- Django admin runs separately for content management
- Both share the same Supabase PostgreSQL database

## Important Notes for Development

**Admin Functionality**:
- ALL admin functionality has been moved to Django
- Go backend no longer has JWT authentication, admin routes, or admin handlers
- No admin-related code remains in the frontend

**Database Connection**:
- Go backend requires `DATABASE_URL` environment variable
- Django uses DB_* variables from .env when available, falls back to defaults
- Both connect to the same PostgreSQL database

**Photo Management**:
- Public contributions can upload photos via Go API
- Photo management (delete, edit) happens in Django admin
- Production photos served directly from Supabase CDN

# Important Instructions
- Use Django admin for ALL content management tasks
- Go backend is PUBLIC API ONLY - no admin functionality
- Frontend is PUBLIC INTERFACE ONLY - no admin routes
- Database operations should prefer Django admin over direct SQL
- When adding features, consider whether it belongs in public API (Go) or admin (Django)