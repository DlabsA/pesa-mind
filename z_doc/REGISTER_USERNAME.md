# Register with Optional Username

## 📋 Updated Endpoint

### POST /api/v1/users/register

Registers a new user and automatically creates their profile with an optional custom username.

---

## 📝 Request

### Headers
```
Content-Type: application/json
```

### Body Schema

| Field | Type | Required | Validation | Description |
|-------|------|----------|-----------|-------------|
| `email` | string | Yes | valid email | User's email address (must be unique) |
| `password` | string | Yes | min 8 chars | Account password |
| `username` | string | No | 3-50 chars | Custom display username. **If omitted or empty, defaults to email** |

### Examples

#### With Custom Username
```json
{
  "email": "john@example.com",
  "password": "SecurePassword123",
  "username": "johnsmith"
}
```

#### With Default Username (empty string)
```json
{
  "email": "jane@example.com",
  "password": "SecurePassword123",
  "username": ""
}
```

#### Without Username Field (defaults to email)
```json
{
  "email": "alice@example.com",
  "password": "SecurePassword123"
}
```

---

## ✅ Response (201 Created)

### Success Response

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "john@example.com",
  "profile": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "johnsmith",
    "type": "Free",
    "balance": 0.0
  }
}
```

### Error Responses

#### 400 - Invalid Request
```json
{
  "error": "Key: 'RegisterRequest.Username' Error:Field validation for 'Username' failed on the 'min' tag"
}
```

#### 400 - Invalid Email
```json
{
  "error": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

#### 400 - Short Password
```json
{
  "error": "Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag"
}
```

#### 500 - Server Error
```json
{
  "error": "failed to hash password"
}
```

---

## 🔍 Validation Rules

| Field | Rule | Example |
|-------|------|---------|
| **Email** | Must be valid email | ✅ `user@example.com` ❌ `invalid-email` |
| **Password** | Minimum 8 characters | ✅ `SecurePass123` ❌ `pass` |
| **Username** | 3-50 characters, optional | ✅ `johnsmith` ❌ `ab` (too short) |
| **Username** | If empty/omitted = email | ✅ Defaults to email automatically |

---

## 🔄 Username Logic

The username field follows this priority:

1. **If provided and valid** → Use provided username
2. **If empty string** → Defaults to email
3. **If omitted** → Defaults to email
4. **If invalid** → Returns 400 error

### Examples

```javascript
// Input: { email: "john@ex.com", username: "johnsmith" }
// Result: username = "johnsmith" ✅

// Input: { email: "john@ex.com", username: "" }
// Result: username = "john@ex.com" ✅

// Input: { email: "john@ex.com" }
// Result: username = "john@ex.com" ✅

// Input: { email: "john@ex.com", username: "ab" }
// Result: Error - username too short ❌
```

---

## 🧪 Testing Examples

### cURL - With Custom Username
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "MySecurePass123",
    "username": "johndoe"
  }'
```

### cURL - Default to Email
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane.smith@example.com",
    "password": "MySecurePass456"
  }'
```

### JavaScript/Fetch
```javascript
const response = await fetch('http://localhost:8080/api/v1/users/register', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    email: 'user@example.com',
    password: 'SecurePassword123',
    username: 'customusername' // optional
  })
});

const data = await response.json();
console.log(data.profile.username);
```

### Python/Requests
```python
import requests

response = requests.post(
  'http://localhost:8080/api/v1/users/register',
  json={
    'email': 'user@example.com',
    'password': 'SecurePassword123',
    'username': 'customusername'  # optional
  }
)

print(response.json()['profile']['username'])
```

---

## 📊 Database Impact

### Users Table
```sql
INSERT INTO users (id, email, password_hash)
VALUES ('uuid-1', 'john@example.com', '$2a$10$...');
```

### Profiles Table
```sql
INSERT INTO profiles (id, user_id, username, type, balance)
VALUES (
  'uuid-2',
  'uuid-1',
  'johnsmith',        -- From request or defaults to email
  'Free',             -- Default
  0.0                 -- Default
);
```

---

## 🔐 Security Notes

- ✅ Passwords are hashed with bcrypt before storage
- ✅ Email is unique (UNIQUE constraint in DB)
- ✅ Username is unique (UNIQUE constraint in DB)
- ✅ User + Profile created atomically (transaction)
- ✅ No sensitive data in response

---

## 📈 Profile Defaults

Every registered user automatically gets a profile with:
| Field | Default Value |
|-------|---------------|
| `type` | `"Free"` |
| `balance` | `0.0` |
| `username` | Email (if not provided) |

---

## 🧩 Code Implementation

### Updated DTO
```go
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
}
```

### Updated Service
```go
func (s *Service) Register(email, passwordHash string, username string) (*User, error) {
	// Default username to email if not provided
	if username == "" {
		username = email
	}
	// ... rest of registration logic
}
```

### Updated Handler
```go
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	// ... validation ...
	usr, err := h.Service.Register(req.Email, string(hashed), req.Username)
	// ... response ...
}
```

---

## ✨ Summary

- ✅ Username is now **optional** in registration request
- ✅ Defaults to **email** if not provided or empty
- ✅ Validates length: **3-50 characters**
- ✅ Automatically creates **Profile** with defaults
- ✅ Profile returned in **response**
- ✅ Atomic **transaction** for consistency
- ✅ Fully **backward compatible**

---

**Last Updated:** April 9, 2026

