# PesaMind API Documentation

## Overview

PesaMind is a modern, offline-first personal finance backend API. This document provides comprehensive information about all available endpoints, request/response formats, and authentication.

## Base URL

- **Development**: `http://localhost:8080/api/v1`
- **Production**: `https://api.pesamind.app/api/v1`

## Authentication

The API uses **JWT (JSON Web Tokens)** for authentication. 

### Getting Started

1. **Register** a new account using `/users/register`
2. **Login** using `/auth/login` to get access and refresh tokens
3. **Include the token** in the `Authorization` header for protected endpoints:
   ```
   Authorization: Bearer <your_access_token>
   ```

### Token Refresh

When your access token expires, use the refresh token with `/auth/refresh` to get a new one:
```json
{
  "refresh_token": "your_refresh_token"
}
```

## API Endpoints

### Health Check

#### GET `/health`
Health check endpoint to verify server status.

**Response (200 OK):**
```json
{
  "status": "ok"
}
```

---

### User Management

#### POST `/users/register`
Register a new user account.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123",
  "username": "johndoe"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "johndoe",
    "type": "Free",
    "balance": 0.0
  }
}
```

**Validation:**
- Email must be valid format
- Password must be at least 8 characters
- Username (optional) must be 3-50 characters, defaults to email if not provided

---

#### GET `/users/me`
Get current user profile (requires authentication).

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "johndoe",
    "type": "Free",
    "balance": 0.0
  }
}
```

---

#### PATCH `/users/me`
Update user profile (requires authentication).

**Request:**
```json
{
  "email": "newemail@example.com",
  "username": "newusername"
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "newemail@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "newusername",
    "type": "Free",
    "balance": 0.0
  }
}
```

---

#### POST `/users/me/change_password`
Change user password (requires authentication).

**Request:**
```json
{
  "current_password": "OldPassword123",
  "new_password": "NewPassword456"
}
```

**Response (200 OK):**
```json
{
  "message": "password changed successfully"
}
```

---

### Authentication

#### POST `/auth/login`
Login with email and password.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123"
}
```

**Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "johndoe",
    "type": "Free",
    "balance": 0.0
  }
}
```

---

#### POST `/auth/refresh`
Refresh access token using refresh token.

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Categories

#### POST `/categories`
Create a new category (requires authentication).

**Request:**
```json
{
  "name": "Groceries",
  "type": "expense",
  "parent_id": null
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Groceries",
  "type": "expense",
  "parent_id": null
}
```

---

#### GET `/categories`
List all user categories (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Groceries",
    "type": "expense",
    "parent_id": null
  },
  {
    "id": "550e8400-e29b-41d4-a716-446655440003",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Salary",
    "type": "income",
    "parent_id": null
  }
]
```

---

### Transactions

#### POST `/transactions`
Create a new transaction (requires authentication).

**Request:**
```json
{
  "profile_id": "550e8400-e29b-41d4-a716-446655440001",
  "category_id": "550e8400-e29b-41d4-a716-446655440002",
  "amount": 50.00,
  "type": "expense",
  "note": "Weekly groceries",
  "date": 1713456000
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440004",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "profile_id": "550e8400-e29b-41d4-a716-446655440001",
  "category_id": "550e8400-e29b-41d4-a716-446655440002",
  "amount": 50.00,
  "type": "expense",
  "note": "Weekly groceries",
  "date": 1713456000
}
```

**Note:** `date` should be a Unix timestamp (seconds since epoch)

---

#### GET `/transactions`
List all user transactions (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440004",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "profile_id": "550e8400-e29b-41d4-a716-446655440001",
    "category_id": "550e8400-e29b-41d4-a716-446655440002",
    "amount": 50.00,
    "type": "expense",
    "note": "Weekly groceries",
    "date": 1713456000
  }
]
```

---

### Budgets

#### POST `/budgets`
Create a new budget (requires authentication).

**Request:**
```json
{
  "name": "Monthly Groceries",
  "amount": 500.00,
  "period": "monthly",
  "start_date": 1713456000,
  "end_date": 1716048000
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440005",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Monthly Groceries",
  "amount": 500.00,
  "period": "monthly",
  "start_date": 1713456000,
  "end_date": 1716048000
}
```

---

#### GET `/budgets`
List all user budgets (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440005",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Monthly Groceries",
    "amount": 500.00,
    "period": "monthly",
    "start_date": 1713456000,
    "end_date": 1716048000
  }
]
```

---

### Savings Goals

#### POST `/savingsgoals`
Create a new savings goal (requires authentication).

**Request:**
```json
{
  "title": "Vacation Fund",
  "target": 5000.00,
  "deadline": 1724000000,
  "auto_save": true
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440006",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Vacation Fund",
  "target": 5000.00,
  "current": 0.0,
  "deadline": 1724000000,
  "auto_save": true
}
```

---

#### GET `/savingsgoals`
List all user savings goals (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440006",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Vacation Fund",
    "target": 5000.00,
    "current": 1200.00,
    "deadline": 1724000000,
    "auto_save": true
  }
]
```

---

### Analytics

#### POST `/analytics`
Create analytics snapshot (requires authentication).

**Request:**
```json
{
  "type": "spending_trend",
  "data": "{\"monthly_trend\": [1000, 1200, 1100]}",
  "period": "monthly"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440007",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "spending_trend",
  "data": "{\"monthly_trend\": [1000, 1200, 1100]}",
  "period": "monthly",
  "created_at": 1713456000,
  "updated_at": 1713456000
}
```

---

#### GET `/analytics`
List analytics snapshots (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440007",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "type": "spending_trend",
    "data": "{\"monthly_trend\": [1000, 1200, 1100]}",
    "period": "monthly",
    "created_at": 1713456000,
    "updated_at": 1713456000
  }
]
```

---

#### GET `/analytics/income`
Get total income (requires authentication).

**Response (200 OK):**
```json
{
  "total": 15000.00
}
```

---

#### GET `/analytics/expenses`
Get total expenses (requires authentication).

**Response (200 OK):**
```json
{
  "total": 3500.00
}
```

---

#### GET `/analytics/budget-utilization`
Get budget utilization data (requires authentication).

**Response (200 OK):**
```json
{
  "total_budget": 10000.00,
  "spent": 3500.00,
  "remaining": 6500.00,
  "utilization_percentage": 35.0
}
```

---

#### GET `/analytics/savings-progress`
Get savings progress (requires authentication).

**Response (200 OK):**
```json
{
  "total_saved": 5200.00,
  "total_goal": 10000.00,
  "progress_percentage": 52.0,
  "active_goals": 3
}
```

---

### Automation

#### POST `/automation/sms`
Create SMS automation rule (requires authentication).

**Request:**
```json
{
  "trigger": "daily",
  "action": "send_reminder"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440008",
  "trigger": "daily",
  "action": "send_reminder"
}
```

---

#### GET `/automation/sms`
List SMS automation rules (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440008",
    "trigger": "daily",
    "action": "send_reminder"
  }
]
```

---

#### POST `/automation/sms/transaction`
Create transaction from SMS (requires authentication).

**Request:**
```json
{
  "message": "Spent 50 on groceries"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440009",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "profile_id": "550e8400-e29b-41d4-a716-446655440001",
  "category_id": "550e8400-e29b-41d4-a716-446655440002",
  "amount": 50.00,
  "type": "expense",
  "note": "SMS parsed: Spent 50 on groceries",
  "date": 1713456000
}
```

---

### Notifications

#### GET `/notifications`
List user notifications (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440010",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Budget Alert",
    "message": "You've reached 80% of your monthly budget",
    "read": false,
    "created_at": 1713456000
  }
]
```

---

#### POST `/notifications/{id}/read`
Mark notification as read (requires authentication).

**Response (200 OK):**
```json
{
  "message": "Notification marked as read"
}
```

---

#### GET `/notifications/preferences`
Get notification preferences (requires authentication).

**Response (200 OK):**
```json
{
  "email_notifications": true,
  "push_notifications": true,
  "sms_notifications": false,
  "budget_alerts": true,
  "transaction_alerts": true,
  "goal_reminders": true
}
```

---

#### POST `/notifications/preferences`
Set notification preferences (requires authentication).

**Request:**
```json
{
  "email_notifications": true,
  "push_notifications": true,
  "sms_notifications": false,
  "budget_alerts": true,
  "transaction_alerts": true,
  "goal_reminders": true
}
```

**Response (200 OK):**
```json
{
  "message": "Preferences updated successfully"
}
```

---

### Gamification

#### GET `/gamification/badges`
List user badges (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440011",
    "name": "First Transaction",
    "description": "Complete your first transaction",
    "earned": true,
    "earned_at": 1713456000
  }
]
```

---

#### GET `/gamification/streaks`
List user streaks (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440012",
    "name": "Daily Tracker",
    "current_count": 15,
    "best_count": 30,
    "last_updated": 1713456000
  }
]
```

---

#### GET `/gamification/achievements`
List user achievements (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440013",
    "name": "Budget Master",
    "description": "Stay within budget for 3 months",
    "completed": true,
    "completed_at": 1713456000,
    "points": 100
  }
]
```

---

#### GET `/gamification/leaderboard`
Get leaderboard (requires authentication).

**Response (200 OK):**
```json
[
  {
    "rank": 1,
    "user_id": "550e8400-e29b-41d4-a716-446655440014",
    "username": "budgetking",
    "points": 5000,
    "badges_count": 25
  },
  {
    "rank": 2,
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "johndoe",
    "points": 4800,
    "badges_count": 23
  }
]
```

---

#### GET `/gamification/rewards`
List available rewards (requires authentication).

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440015",
    "name": "Premium Upgrade",
    "description": "Upgrade to Premium membership for 1 month",
    "points_required": 1000,
    "available": true
  }
]
```

---

#### POST `/gamification/rewards/{reward_id}/claim`
Claim a reward (requires authentication).

**Response (200 OK):**
```json
{
  "message": "Reward claimed successfully",
  "reward": {
    "id": "550e8400-e29b-41d4-a716-446655440015",
    "name": "Premium Upgrade",
    "claimed_at": 1713456000
  }
}
```

---

## Error Handling

All error responses follow this format:

```json
{
  "error": "Error message describing what went wrong"
}
```

### Common HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Missing or invalid token |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

---

## Rate Limiting

The API implements rate limiting to prevent abuse. Current limits:
- **Public endpoints**: 100 requests per minute per IP
- **Protected endpoints**: 1000 requests per minute per authenticated user

---

## Additional Notes

- All timestamps are in **Unix timestamp format** (seconds since epoch)
- All IDs are **UUIDs** (Universally Unique Identifiers)
- Monetary values are in **double precision floating-point** format
- The API uses **UTC timezone** for all dates

---

## OpenAPI Specification

For detailed API specification, see the OpenAPI 3.0 file: `docs/openapi.yaml`

To view it interactively, use tools like:
- **Swagger UI**: https://editor.swagger.io/
- **ReDoc**: https://redoc.ly/
- **Postman**: Import the OpenAPI file directly

