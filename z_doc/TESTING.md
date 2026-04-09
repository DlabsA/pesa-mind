# Testing Guide - PesaMind Profile Integration

## Overview
This guide explains how to test the refactored Profile model integration with User registration and authentication.

## Test Files Created

### 1. **Unit Tests - Repository Layer**
**File:** `internal/domain/user/gorm_repository_test.go`

Tests the GORM repository implementation:
- `TestGormUserRepositoryCreate` - Verify user and profile creation in a transaction
- `TestGormUserRepositoryFindByID` - Retrieve user with profile by ID
- `TestGormUserRepositoryFindByEmail` - Retrieve user with profile by email
- `TestGormUserRepositoryUpdate` - Update user information
- `TestGormUserRepositoryDelete` - Soft delete user

**Run tests:**
```bash
cd internal/domain/user
go test -v -run TestGormUserRepository
```

### 2. **Integration Tests - Service Layer**
**File:** `internal/domain/user/service_test.go`

Tests the business logic and service interactions:
- `TestServiceRegisterWithProfile` - Register user and verify profile is created
- `TestServiceGetByIDWithProfile` - Fetch user with profile by ID
- `TestServiceGetByEmailWithProfile` - Fetch user with profile by email
- `TestServiceRegisterMultipleUsers` - Verify multiple users can be registered independently

**Run tests:**
```bash
cd internal/domain/user
go test -v -run TestService
```

### 3. **API Integration Tests**
**File:** `test/api_register_test.go`

Tests the HTTP endpoints end-to-end:
- `TestAPIRegisterWithProfile` - Complete registration flow via API
- `TestAPIRegisterInvalidEmail` - Validation for invalid email
- `TestAPIRegisterShortPassword` - Validation for short password
- `TestAPIRegisterMultipleUsers` - Multiple users via API

**Run tests:**
```bash
cd test
go test -v -run TestAPI
```

---

## Running All Tests

### Run all user domain tests:
```bash
go test -v ./internal/domain/user/...
```

### Run all integration tests:
```bash
go test -v ./test/...
```

### Run entire test suite:
```bash
go test -v ./...
```

### Run with coverage:
```bash
go test -v -cover ./internal/domain/user/
go test -v -cover ./test/
```

### Generate coverage report:
```bash
go test -coverprofile=coverage.out ./internal/domain/user/ ./test/
go tool cover -html=coverage.out -o coverage.html
# Open coverage.html in browser
```

---

## Manual Testing via cURL

### 1. Register a User with Profile
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "SecurePassword123"
  }'
```

**Expected Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "john@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "john@example.com",
    "type": "Free",
    "balance": 0.0
  }
}
```

### 2. Login to Get Tokens (with Profile)
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "SecurePassword123"
  }'
```

**Expected Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "john@example.com",
    "type": "Free",
    "balance": 0.0
  }
}
```

### 3. Health Check
```bash
curl http://localhost:8080/health
```

---

## Starting the Server for Manual Testing

### 1. Start PostgreSQL (if using Docker):
```bash
docker-compose up -d postgres
```

### 2. Run migrations:
```bash
goose -dir migrations postgres "postgres://pesamind:pesamind123@localhost:5432/pesamind?sslmode=disable" up
```

### 3. Start the server:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

---

## Test Scenarios

### Scenario 1: Happy Path - Register and Login
1. Register user with valid email and password
2. Verify user and profile are created
3. Login with the same credentials
4. Verify profile is returned in login response

### Scenario 2: Profile Creation Guarantees
1. Register a user
2. Query database for both User and Profile records
3. Verify Profile.UserID matches User.ID
4. Verify Profile has default values (Type="Free", Balance=0.0)

### Scenario 3: Email Uniqueness
1. Register user with email A
2. Try to register another user with the same email A
3. Verify second registration fails with appropriate error

### Scenario 4: Multiple Users Independence
1. Register 3 different users
2. Verify each has unique ID, Profile.ID, and Username
3. Verify they don't interfere with each other

---

## Troubleshooting

### Test Fails: "failed to connect to db"
- Ensure PostgreSQL is running
- Check `.env` file has correct DB credentials

### Test Fails: "profile should exist"
- Verify Profile model is registered in AutoMigrate
- Check Create transaction is not failing silently

### API returns 500 error
- Check server logs for detailed error message
- Verify JWT middleware is configured correctly

### Coverage below expected
- Run tests with `-v` flag for detailed output
- Check if all code paths are being tested

---

## CI/CD Integration

For GitHub Actions or similar, add to your workflow:

```yaml
- name: Run tests
  run: go test -v -race -coverprofile=coverage.out ./...

- name: Upload coverage
  uses: codecov/codecov-action@v3
  with:
    files: ./coverage.out
```

---

## Performance Testing

To stress test registration endpoint:

```bash
# Using Apache Bench (ab)
ab -n 100 -c 10 -p data.json -T application/json http://localhost:8080/api/v1/users/register

# Using hey
hey -n 100 -c 10 -m POST -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456"}' \
  http://localhost:8080/api/v1/users/register
```

---

## Next Steps

1. Run unit tests: `go test ./internal/domain/user/...`
2. Run integration tests: `go test ./test/...`
3. Start server and test manually with cURL
4. Monitor logs for any issues
5. Verify profile data in database

Good luck! 🚀

