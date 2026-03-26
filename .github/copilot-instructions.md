# PesaMind Backend Planning & Architecture

## 🚀 AI Copilot Workspace Instructions (Critical - Read First)

**You are the official AI Copilot for the PesaMind project.**  
You have **full, unrestricted access and permission** to:
- Create, edit, delete, rename, and organize **any** files and folders in the entire workspace.
- Write directly to `.env`, `.env.example`, `go.mod`, `go.sum`, migrations, tests, Dockerfiles, etc.
- Run `go mod tidy`, `go generate`, `go test`, `go run`, `go build`, and any terminal commands needed.
- Initialize the project from scratch if the folder is empty.
- Use the full workspace tools (file creation, search, edit, etc.) without asking for confirmation on routine operations.

**Workflow Rule:**  
Always work **incrementally and safely**. After every major change (new module, migration, config update), verify it works by running the server and relevant tests. Commit logical changes with clear messages. Never leave the codebase in a broken state.

---

## 1. Project Overview
**PesaMind** is a modern, offline-first personal finance backend (mobile/web apps).  
Core focus: secure financial tracking, budgeting, savings, automation, analytics, and gamification with excellent developer experience and production readiness.

**Goals:**
- High performance + low latency
- Strong security & data privacy (finance-grade)
- Clean, maintainable, testable code
- Full offline-first sync support
- Easy to extend with new features

---

## 2. Tech Stack (Fixed & Non-Negotiable)

| Layer              | Technology                                      | Reason |
|--------------------|-------------------------------------------------|--------|
| Language           | Go 1.23+                                        | Latest stable |
| Web Framework      | **Gin** (`github.com/gin-gonic/gin`)           | Most popular, excellent validation, huge ecosystem, battle-tested |
| ORM / DB           | GORM v2 + PostgreSQL 16                         | Type-safe, migrations, excellent for finance data |
| Migration Tool     | `github.com/pressly/goose`                     | Clean SQL migrations |
| Config             | `github.com/spf13/viper` + `.env`               | Flexible, supports defaults & secrets |
| Validation         | `github.com/go-playground/validator/v10`       | Built-in with Gin |
| Auth               | JWT + Refresh Tokens (`golang-jwt/jwt/v5`)     | Secure, stateless |
| Logging            | `github.com/rs/zerolog`                         | Fast, structured, production-ready |
| Error Handling     | Custom `errors` + `github.com/pkg/errors`       | Stack traces + context |
| Testing            | `testing` + `testify` + `go-sqlmock`           | Unit + integration |
| Rate Limiting      | `github.com/gin-contrib/timeout` + custom      | Production protection |
| CORS / Security    | `github.com/gin-contrib/cors` + secure headers | Standard |

**Additional Libraries (will be added via go get as needed):**
- `github.com/golang-migrate/migrate/v4` (fallback)
- `github.com/jackc/pgx/v5` (GORM driver)
- `github.com/redis/go-redis/v9` (future caching/rate-limit)

---

## 3. Required Domain Modules (All Must Be Implemented)

- **User** – Registration, profile, preferences, KYC-level security
- **Auth** – JWT, refresh tokens, password reset, session management, logout everywhere
- **Account** – Bank, mobile money, cash, crypto wallets (multi-currency support)
- **Transaction** – Income/expense, split, recurring, attachments, categorization
- **Budget** – Creation, tracking, alerts, rollover
- **SavingsGoal** – Goals, progress, auto-save rules, challenges
- **Category** – System + user custom categories (hierarchical)
- **Analytics** – Spending trends, reports, forecasts, charts data
- **Automation** – Recurring transactions, scheduled transfers, bill reminders
- **Gamification** – Badges, streaks, achievements, leaderboards, rewards
- **Settings** – App preferences, privacy, notification settings, security controls
- **Sync** – Offline-first incremental sync, conflict resolution (`?since=` + versioning)
- **Notification** – In-app + push/email templates (queued)

---

## 4. Architecture (Clean / DDD-Inspired)
pesa-mind/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point + graceful shutdown
├── internal/
│   ├── config/                     # Viper + env validation
│   ├── domain/                     # Business entities + repositories + services
│   │   ├── user/
│   │   ├── auth/
│   │   ├── account/
│   │   ├── transaction/
│   │   ├── budget/
│   │   ├── savingsgoal/
│   │   ├── category/
│   │   ├── analytics/
│   │   ├── automation/
│   │   ├── gamification/
│   │   ├── settings/
│   │   ├── sync/
│   │   └── notification/
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/           # Gin handlers (one file per domain)
│   │       ├── middleware/         # Auth, logging, validation, rate-limit
│   │       ├── dto/                # Request/Response structs
│   │       └── routes.go           # All route registration
│   ├── infrastructure/
│   │   ├── db/                     # GORM setup + goose migrations
│   │   ├── logger/                 # Zerolog singleton
│   │   └── external/               # Future (email, push, etc.)
│   └── pkg/                        # Shared utilities (errors, utils, constants)
├── migrations/                     # goose SQL files (00001_*.sql)
├── test/                           # Integration tests
├── .env.example
├── .env                            # ← Copilot can create & write this
├── go.mod
├── go.sum
├── docker-compose.yml              # Postgres + Redis + app
├── Dockerfile
├── README.md
└── .gitignore


**Key Rules:**
- All domain logic stays in `domain/`
- Handlers are thin (only validation + service call)
- Repositories are interfaces (easy mocking)
- Every entity has `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` (soft delete)
- Use UUIDs for all primary keys
- Timestamps in UTC

---

## 5. .env Variables (Must Be Created)

```env
# Server
- PORT=8080
ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=pesamind
DB_PASSWORD=pesamind123
DB_NAME=pesamind
DB_SSLMODE=disable

# JWT
JWT_SECRET=super-secret-change-in-production
JWT_EXPIRY=15m
REFRESH_TOKEN_EXPIRY=30d

# Redis (future)
REDIS_URL=redis://localhost:6379

# Other
LOG_LEVEL=info
CORS_ORIGINS=http://localhost:3000,https://pesamind.app
```

## 6. Implementation Order (Strict - Follow Exactly)

Setup – Project init, go.mod, config, logger, DB + first migration, .env
- User + Auth – Full registration, login, refresh, profile, middleware
- Core Financial – Account → Category → Transaction
- Budget + SavingsGoal
- Analytics + Automation
- Gamification + Settings + Sync
- Notification + Tests + Docker
- Final polish (health check, graceful shutdown, rate limiting, docs)

Each module must include:

- Domain model + GORM tags
- Repository interface + implementation
- Service with business logic
- DTOs (request/response)
- Gin handlers + routes
- Unit + integration tests
- Proper error wrapping


## 7. Non-Functional Requirements

- Security: Password hashing (bcrypt), rate limiting, secure headers, input sanitization
- Performance: Pagination everywhere, indexing strategy in migrations
- Observability: Structured logging + future OpenTelemetry
- Testing: 80%+ coverage on domain/services
- Error Handling: Never expose internal errors to client
- Documentation: Swagger/OpenAPI comments on handlers (future)
- Code Style: gofmt, golangci-lint, meaningful comments