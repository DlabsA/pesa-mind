# Quick Reference - Username Field Update

## Endpoint Updated
```
POST /api/v1/users/register
```

## New Request Format

### With Custom Username
```json
{
  "email": "user@example.com",
  "password": "SecurePass123",
  "username": "customusername"
}
```

### Without Username (uses email as default)
```json
{
  "email": "user@example.com",
  "password": "SecurePass123"
}
```

## Response (Both cases)
```json
{
  "id": "uuid-1",
  "email": "user@example.com",
  "profile": {
    "id": "uuid-2",
    "user_id": "uuid-1",
    "username": "customusername",  // or "user@example.com" if not provided
    "type": "Free",
    "balance": 0.0
  }
}
```

## Validation
- ✅ Email: Required, valid email format
- ✅ Password: Required, minimum 8 characters
- ✅ Username: Optional, 3-50 characters if provided

## Files Changed
1. `internal/interfaces/http/dto/user_dto.go` - Added username field
2. `internal/domain/user/service.go` - Updated Register method
3. `internal/interfaces/http/handlers/user_handler.go` - Pass username to service

## Build Status
```bash
go build ./cmd/api
✅ Build successful
```

## Tests Status
```bash
go test -v ./internal/domain/user/models_test.go ./internal/domain/user/model.go
✅ All 5 tests PASS
```

## Default Behavior
- If `username` is **omitted** → Uses **email**
- If `username` is **empty string** → Uses **email**
- If `username` is **provided** → Uses **provided value**

## Example Commands

### cURL with Custom Username
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"SecurePass123","username":"johndoe"}'
```

### cURL with Default Username
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"SecurePass123"}'
```

---

For detailed information, see:
- **USERNAME_UPDATE.md** - Complete update summary
- **REGISTER_USERNAME.md** - Full endpoint documentation

