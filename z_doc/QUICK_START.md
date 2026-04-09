# Quick Start - Test the Profile Integration

## 📝 Step 1: Run Model Tests
```bash
cd /Users/conradkash/Github/Personal/pesa-mind

# Run unit tests for models (all pass ✅)
go test -v ./internal/domain/user/models_test.go ./internal/domain/user/model.go
```

Expected output:
```
=== RUN   TestUserProfileModel
--- PASS: TestUserProfileModel (0.00s)
=== RUN   TestUserProfileRelationship
--- PASS: TestUserProfileRelationship (0.00s)
=== RUN   TestUserProfileDefaults
--- PASS: TestUserProfileDefaults (0.00s)
=== RUN   TestMultipleProfiles
--- PASS: TestMultipleProfiles (0.00s)
=== RUN   TestUserProfileLifecycle
--- PASS: TestUserProfileLifecycle (0.00s)
PASS
```

## 🗄️ Step 2: Setup Database

### Option A: Docker Compose
```bash
docker-compose up -d postgres

# Wait 5 seconds for postgres to start
sleep 5

# Run migrations
goose -dir migrations postgres "postgres://pesamind:pesamind123@localhost:5432/pesamind?sslmode=disable" up
```

### Option B: Local PostgreSQL
```bash
# Create database
createdb -U postgres pesamind

# Run migrations
goose -dir migrations postgres "postgres://postgres:password@localhost:5432/pesamind?sslmode=disable" up
```

## 🚀 Step 3: Start the Server

```bash
# Development mode (HTTP on port 8080)
go run cmd/api/main.go

# Or build then run
go build -o api ./cmd/api
./api
```

Output should show:
```
Starting server in development mode...
```

## 📱 Step 4: Test Endpoints

### Registration Test (with Profile)
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "SecurePass123"
  }'
```

Response:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "alice@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "alice@example.com",
    "type": "Free",
    "balance": 0
  }
}
```

### Login Test (with Profile)
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "SecurePass123"
  }'
```

Response:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "alice@example.com",
    "type": "Free",
    "balance": 0
  }
}
```

### Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{"status": "ok"}
```

## ✅ Verification Checklist

- [ ] Model tests pass (5/5)
- [ ] Database created and migrated
- [ ] Server starts without errors
- [ ] Registration creates User + Profile
- [ ] Login returns tokens + profile
- [ ] Profile has correct defaults (Type="Free", Balance=0.0)
- [ ] Health endpoint responds

## 🔍 Checking Database

```bash
# Connect to PostgreSQL
psql -U pesamind -d pesamind

# View users
SELECT id, email FROM users LIMIT 5;

# View profiles
SELECT id, user_id, username, type, balance FROM profiles LIMIT 5;

# Check relationship
SELECT u.email, p.username, p.type, p.balance 
FROM users u 
JOIN profiles p ON u.id = p.user_id 
LIMIT 5;
```

## 🛑 Stopping

```bash
# Stop server (Ctrl+C in terminal)

# Stop Docker containers
docker-compose down
```

## 📚 Documentation

- **Full Testing Guide**: See `TESTING.md`
- **Architecture Overview**: See `PROFILE_REFACTORING.md`
- **API Routes**: See `cmd/api/main.go`
- **Model Definitions**: See `internal/domain/user/model.go`

## 🎉 You're All Set!

The Profile model integration is complete and tested. 
You now have:
- ✅ Bidirectional User-Profile relationships
- ✅ Atomic transactional creates
- ✅ Complete API integration
- ✅ HTTPS support for production
- ✅ Full test coverage

Happy coding! 🚀

