# 🚀 DEPLOYMENT FIX - QUICK START

## ❌ Problem
```
Error: Bind for 0.0.0.0:80 failed: port is already allocated
```

Port 80 and 443 are in use by your existing `traefik` container.

---

## ✅ Solution

### Option 1: Quick Manual Fix (SSH to Server)

```bash
cd ~/pesa-mind

# Stop containers
docker-compose down

# Redeploy
docker-compose up -d

# Verify
docker ps
curl http://localhost:8080/health
```

### Option 2: Use Deployment Script

```bash
cd ~/pesa-mind

# Make executable
chmod +x deploy.sh

# Run deployment
./deploy.sh
```

---

## ✅ What Changed

**File:** `docker-compose.yml`

**Change:** Caddy now uses ports 8080 and 8443 instead of 80 and 443

```yaml
# BEFORE (Caused conflict)
ports:
  - "80:80"
  - "443:443"

# AFTER (No conflict)
ports:
  - "8080:80"
  - "8443:443"
```

---

## 📊 Access After Deployment

| Service | Port | URL |
|---------|------|-----|
| HTTP | 8080 | `http://server:8080` |
| HTTPS | 8443 | `https://server:8443` |
| Traefik | 80 | `http://server:80` (existing) |
| Traefik | 443 | `https://server:443` (existing) |

---

## 🔍 Verify Deployment

```bash
# See all running containers
docker ps

# Check if pesa-mind is healthy
docker logs pesa-mind-app-1

# Test API
curl http://localhost:8080/health

# Expected response
{"status": "ok"}
```

---

## 🆘 If Still Having Issues

```bash
# Show what's using port 80
lsof -i :80

# Show what's using port 443
lsof -i :443

# Remove pesa-mind volumes and start fresh
docker-compose down -v
docker-compose up -d
```

---

## 📁 Files Created/Updated

1. ✏️ `docker-compose.yml` - Updated port mappings
2. 📄 `DEPLOYMENT_FIX.md` - Detailed documentation
3. 📄 `deploy.sh` - Automated deployment script

---

## 🎯 Summary

✅ Port conflict identified and fixed  
✅ Pesa-mind now uses ports 8080 (HTTP) and 8443 (HTTPS)  
✅ Can coexist with existing traefik on ports 80/443  
✅ Ready for immediate redeployment  

---

## 🚀 Next Command

On your deployment server, run:

```bash
cd ~/pesa-mind
docker-compose down
docker-compose up -d
docker ps  # Verify all running
```

**Your app will be live at:** `http://your-server:8080`

