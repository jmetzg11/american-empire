# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Main Commands
- `make run` - Start database, Django admin, and Go web server (complete development environment)
- `make stop` - Stop all services (Django, Go, database)
- `make restart` - Clean and restart Go web server
- `make start-db` - Start PostgreSQL Docker container
- `make clean` - Clean build artifacts and recreate tmp directory
- `make test` - Run Go tests
- `make kill-port` - Kill process on port 8080

### Go Web Server
- `cd web && air` - Start Go web server with hot reload on port 8080 (development)

### Django Admin (Content Management)
- `cd admin && make dev` - Start Django admin on local Docker database (development)
- `cd admin && make prod` - Start Django admin on Supabase database (production)
- `cd admin && make stop` - Stop Django server
- `cd admin && make shell` - Access Django shell
- `cd admin && make migrate` - Run Django migrations (dev database only)

## Architecture Overview

This is a full-stack web application for tracking historical events with a clean separation between public web interface and admin functionality.

### Go Web Server (`web/`) - Public Interface
- **Main server**: `web/main.go` - HTTP server on port 8080
- **Database**: Standard library `database/sql` with PostgreSQL driver (`lib/pq`)
- **Templates**: Go `html/template` for server-side rendering
- **Purpose**: Serves public web pages with HTML templates

**Key Files**:
- `main.go` - HTTP server setup, database connection, template caching
- `routes.go` - Route definitions and static file serving
- `handlers.go` - HTTP request handlers
- `models.go` - Database queries and data models
- `helpers.go` - Template rendering and database connection utilities

**Template Structure** (`web/ui/html/`):
- `base.tmpl` - Base HTML layout
- `pages/` - Page-specific templates

**Static Assets** (`web/ui/static/`):
- CSS, JavaScript, and other static files
- Served at `/static/` path

**Routes**:
- `GET /` - Home page displaying all active events in a table
- `GET /static/*` - Static file serving

**Database Models**:
- `EventSummary` - Displays event ID, title, date, country, and tags

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
  - Go: Uses local Docker PostgreSQL (connection string in `helpers.go:connectDB()`)
  - Django: Uses local Docker PostgreSQL via DB_* environment variables
- **Production**:
  - Go: Uses production PostgreSQL via DATABASE_URL
  - Django: Uses Supabase PostgreSQL via DATABASE_URL

**Key Environment Variables**:
- `DATABASE_URL` - PostgreSQL connection (production)
- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` - Django database config (development)
- `ADMIN_USERNAME/ADMIN_PASSWORD` - Django admin credentials
- `DEBUG` - Django debug mode

### Database Schema
Events are the central entity with:
- One-to-many relationships with Sources and Media
- Many-to-many relationship with Tags via junction table `event_tags`
- Many-to-many relationship with Books via junction table `book_events`
- `Active` timestamp field for approval workflow (NULL = not approved)
- `Flagged` boolean for content moderation

## Development Workflow

### Full Development Setup
1. Quick start: `make run` (starts database, Django admin, and Go web server)
2. Access web app at `http://localhost:8080`
3. Access admin at `http://localhost:8000/admin/`

### Manual Setup
1. Start local PostgreSQL: `cd data && docker compose up -d postgres`
2. Start Django admin: `cd admin && make dev`
3. Start Go web server: `cd web && air`

### Content Management Workflow
1. Django admin is the primary interface for all content management
2. Go web server only displays public content
3. All CRUD operations for events, tags, sources, media happen in Django admin
4. Go server queries the database in read-only mode for public display

### Production Deployment
- Go web server serves public HTML pages
- Django admin runs separately for content management
- Both share the same PostgreSQL database (Supabase in production)

## Important Notes for Development

**Admin Functionality**:
- ALL admin functionality is in Django
- Go web server has NO admin functionality - it's read-only for public display
- No authentication or user management in Go

**Database Connection**:
- Go: Uses flag-based environment selection (`-prod` flag for production)
- Django: Uses DB_* variables from .env when available, falls back to defaults
- Both connect to the same PostgreSQL database

**Template System**:
- Go uses standard library `html/template`
- Templates are cached at startup for performance
- Templates are embedded in the binary using `embed.FS`

# Important Instructions
- Use Django admin for ALL content management tasks
- Go web server is PUBLIC DISPLAY ONLY - read-only database access
- No admin routes or functionality in Go
- Database writes only happen through Django admin
- When adding features, consider whether it belongs in public display (Go) or admin (Django)