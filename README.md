# PesaMind Backend

A modern, offline-first personal finance backend for mobile/web apps.

## Features
- Secure financial tracking, budgeting, savings, automation, analytics, gamification
- Clean, maintainable, testable Go code
- Gin, GORM, PostgreSQL, JWT, Goose, Viper, Zerolog
- Full offline-first sync support

## Setup

1. Copy `.env.example` to `.env` and adjust as needed.
2. Start PostgreSQL (see `docker-compose.yml` for example).
3. Run DB migrations (using Goose):
   ```sh
   goose -dir ./migrations postgres "host=localhost user=pesamind password=pesamind123 dbname=pesamind sslmode=disable" up
   ```
4. Run the server:
   ```sh
   go run ./cmd/api
   ```

## Testing

Run all tests:
```sh
go test ./...
```

---

See `.github/copilot-instructions.md` for architecture and contribution guidelines.

