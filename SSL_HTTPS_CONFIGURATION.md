# 🔒 SSL/HTTPS Configuration - Fix for EPROTO Error

## ❌ Problem
```
EPROTO: error:100000f7:SSL routines:OPENSSL_internal:WRONG_VERSION_NUMBER
```

You were trying to use HTTPS on port 8080:
```
https://173.212.219.227:8080/api/v1/auth/login  ❌ WRONG
```

**Issue:** Port 8080 is HTTP only, not HTTPS.

---

## ✅ Solution

### Correct Port Configuration

| Protocol | Port | URL | Use Case |
|----------|------|-----|----------|
| **HTTP** | 8080 | `http://173.212.219.227:8080` | ✅ Development |
| **HTTPS** | 8443 | `https://173.212.219.227:8443` | ✅ Staging/Testing |
| **HTTP** | 80 | `http://your-domain:80` | ✅ Production (domain) |
| **HTTPS** | 443 | `https://your-domain:443` | ✅ Production (domain) |

---

## 🚀 Correct URLs for Your Setup

### Development/Staging (Local)
```
HTTP:  http://localhost:8080/api/v1/auth/login
HTTPS: https://localhost:8443/api/v1/auth/login
```

### Staging (IP Address)
```
HTTP:  http://173.212.219.227:8080/api/v1/auth/login
HTTPS: https://173.212.219.227:8443/api/v1/auth/login
```

### Production (Domain)
```
HTTP:  http://api.dlabs.cc/api/v1/auth/login
HTTPS: https://api.dlabs.cc/api/v1/auth/login
```

---

## 📝 cURL Examples

### HTTP Request
```bash
curl -X POST http://173.212.219.227:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456"}'
```

### HTTPS Request (with self-signed cert)
```bash
curl -k -X POST https://173.212.219.227:8443/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456"}'
```

**Note:** `-k` flag ignores self-signed certificate verification (OK for testing)

---

## 🧪 Test Your API

### 1. Health Check (HTTP)
```bash
curl http://173.212.219.227:8080/health
# Response: {"status":"ok"}
```

### 2. Health Check (HTTPS)
```bash
curl -k https://173.212.219.227:8443/health
# Response: {"status":"ok"}
```

### 3. Login (HTTPS - Correct)
```bash
curl -k -X POST https://173.212.219.227:8443/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePassword123"
  }'
```

---

## 🔒 SSL Certificate Configuration

### For Development (Self-Signed)
Caddy automatically generates self-signed certificates:
```
tls internal
```

### For Production (Let's Encrypt)
Caddy automatically gets certificates for domains:
```
api.dlabs.cc {
    reverse_proxy app:8080
}
```

---

## 💻 Frontend/Client Configuration

### JavaScript/Fetch
```javascript
// HTTP
fetch('http://173.212.219.227:8080/api/v1/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({...})
})

// HTTPS (with self-signed cert)
fetch('https://173.212.219.227:8443/api/v1/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({...})
})
```

### Postman
1. Select **GET/POST** method
2. Enter URL: `https://173.212.219.227:8443/api/v1/auth/login`
3. Go to **Settings** (gear icon)
4. Disable "SSL certificate verification"
5. Send request

### HTTPie
```bash
# Disable SSL verification
http --verify=no POST https://173.212.219.227:8443/api/v1/auth/login \
  email=test@example.com password=Test123456
```

---

## 🔄 Environment Variables

Update your `.env` for different environments:

### Development
```env
API_BASE_URL=http://localhost:8080
API_PROTOCOL=http
API_HOST=localhost
API_PORT=8080
```

### Staging
```env
API_BASE_URL=https://173.212.219.227:8443
API_PROTOCOL=https
API_HOST=173.212.219.227
API_PORT=8443
```

### Production
```env
API_BASE_URL=https://api.dlabs.cc
API_PROTOCOL=https
API_HOST=api.dlabs.cc
API_PORT=443
```

---

## ✅ Docker Port Mapping

Your docker-compose.yml currently:
```yaml
caddy:
  ports:
    - "8080:80"    # HTTP
    - "8443:443"   # HTTPS
```

This means:
- External port 8080 → Internal port 80 (HTTP)
- External port 8443 → Internal port 443 (HTTPS)

**Access via:**
- `http://server:8080` ← HTTP traffic
- `https://server:8443` ← HTTPS traffic

---

## 🎯 Summary

| Before (❌) | After (✅) |
|------------|----------|
| `https://ip:8080` | `https://ip:8443` |
| EPROTO error | ✅ Works |
| SSL mismatch | ✅ Correct port |

---

## 📚 Related Files

- `Caddyfile` - Updated with SSL configuration
- `docker-compose.yml` - Port mapping (8080→80, 8443→443)
- `.env` - API base URLs

---

## 🚀 Redeploy

After these changes, redeploy:
```bash
cd ~/pesa-mind
git pull
./deploy.sh
```

Then test with correct HTTPS port (8443):
```bash
curl -k https://173.212.219.227:8443/health
```

---

**Fix: Use port 8443 for HTTPS, port 8080 for HTTP** ✅

