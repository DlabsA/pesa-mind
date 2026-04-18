# PesaMind API Quick Reference

## 🚀 Quick Start

### 1. Register
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePassword123",
    "username": "johndoe"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePassword123"
  }'
```

Response includes:
- `access_token` - Use for authenticated requests (expires in 15 minutes)
- `refresh_token` - Use to get new access token (expires in 30 days)

### 3. Create Transaction
```bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "profile_id": "550e8400-e29b-41d4-a716-446655440001",
    "category_id": "550e8400-e29b-41d4-a716-446655440002",
    "amount": 50.00,
    "type": "expense",
    "note": "Groceries",
    "date": 1713456000
  }'
```

## 📌 Core Concepts

### Authentication Flow
```
1. Register → Get User ID & Profile ID
2. Login → Get Access Token + Refresh Token
3. Use Access Token for all protected endpoints
4. When expired, use Refresh Token to get new Access Token
```

### Data Types
- **IDs**: Always UUID format (e.g., "550e8400-e29b-41d4-a716-446655440000")
- **Timestamps**: Unix seconds (e.g., 1713456000)
- **Money**: Decimal numbers (e.g., 50.00, 1000.50)
- **Types**: Fixed values (e.g., "income", "expense", "Free", "Premium")

## 🔑 Essential IDs for Transactions

To create a transaction, you need:
1. **profile_id** - Get from `/users/me` response → `profile.id`
2. **category_id** - Create with `POST /categories` or get from `GET /categories`

## 🌐 Base URL
```
http://localhost:8080/api/v1
```

## 📚 Full Endpoint List

| Method | Endpoint | Auth Required | Purpose |
|--------|----------|---|---------|
| GET | `/health` | ❌ | Check server status |
| POST | `/users/register` | ❌ | Register new user |
| POST | `/auth/login` | ❌ | Login |
| POST | `/auth/refresh` | ❌ | Refresh token |
| GET | `/users/me` | ✅ | Get current user |
| PATCH | `/users/me` | ✅ | Update user |
| POST | `/users/me/change_password` | ✅ | Change password |
| POST | `/categories` | ✅ | Create category |
| GET | `/categories` | ✅ | List categories |
| POST | `/transactions` | ✅ | Create transaction |
| GET | `/transactions` | ✅ | List transactions |
| POST | `/budgets` | ✅ | Create budget |
| GET | `/budgets` | ✅ | List budgets |
| POST | `/savingsgoals` | ✅ | Create savings goal |
| GET | `/savingsgoals` | ✅ | List savings goals |
| POST | `/analytics` | ✅ | Create snapshot |
| GET | `/analytics` | ✅ | List snapshots |
| GET | `/analytics/income` | ✅ | Total income |
| GET | `/analytics/expenses` | ✅ | Total expenses |
| GET | `/analytics/budget-utilization` | ✅ | Budget status |
| GET | `/analytics/savings-progress` | ✅ | Savings status |
| POST | `/automation/sms` | ✅ | Create SMS rule |
| GET | `/automation/sms` | ✅ | List SMS rules |
| POST | `/automation/sms/transaction` | ✅ | SMS transaction |
| GET | `/notifications` | ✅ | List notifications |
| POST | `/notifications/{id}/read` | ✅ | Mark as read |
| GET | `/notifications/preferences` | ✅ | Get preferences |
| POST | `/notifications/preferences` | ✅ | Set preferences |
| GET | `/gamification/badges` | ✅ | List badges |
| GET | `/gamification/streaks` | ✅ | List streaks |
| GET | `/gamification/achievements` | ✅ | List achievements |
| GET | `/gamification/leaderboard` | ✅ | Get leaderboard |
| GET | `/gamification/rewards` | ✅ | List rewards |
| POST | `/gamification/rewards/{id}/claim` | ✅ | Claim reward |

## 🔒 Error Responses

All errors return JSON with `error` field:
```json
{
  "error": "Invalid email or password"
}
```

### Common Errors
| Status | Error | Solution |
|--------|-------|----------|
| 400 | "Invalid request" | Check required fields |
| 401 | "Unauthorized" | Add Authorization header |
| 404 | "Not found" | Check resource ID |
| 500 | "Internal error" | Contact support |

## 💡 Tips & Tricks

### Get Your Profile ID
```bash
curl -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/users/me \
  | jq '.profile.id'
```

### List Categories to Get Category ID
```bash
curl -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/categories
```

### Convert Date to Unix Timestamp
```bash
# macOS/Linux
date -d "2024-04-18" +%s

# JavaScript
new Date("2024-04-18").getTime() / 1000

# Python
import time
time.mktime(time.strptime("2024-04-18", "%Y-%m-%d"))
```

### Refresh Token When Expired
```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "<refresh_token>"}'
```

## 📖 Documentation Files

- **Full API Reference**: See `docs/API.md`
- **OpenAPI Spec**: See `docs/openapi.yaml`
- **Architecture Guide**: See `REFACTORING_ACCOUNT_TO_PROFILE.md`

## 🧪 Testing Tools

- **Postman**: Import `docs/openapi.yaml`
- **Swagger UI**: Visit https://editor.swagger.io/ and import YAML
- **cURL**: Examples above
- **Insomnia**: Import OpenAPI file

## 🆘 Support

For detailed endpoint documentation with all request/response examples, see `docs/API.md`

---

**Last Updated**: April 18, 2026

