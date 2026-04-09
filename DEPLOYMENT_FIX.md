# 🚀 PesaMind Deployment Fix - Port Conflict Resolution

## 🔴 Problem
Docker deployment was failing with:
```
Error: Bind for 0.0.0.0:80 failed: port is already allocated
```

## ✅ Root Cause
Port 80 and 443 are already in use by the existing `traefik` reverse proxy on your deployment server.

Your running containers:
- `uganda-gis-traefik` → Uses ports 80, 443
- `pesa-mind-caddy` → Was trying to use ports 80, 443 ❌ CONFLICT

## ✅ Solution Implemented

### Changed Ports in `docker-compose.yml`

**Before:**
```yaml
caddy:
  ports:
    - "80:80"      # HTTP
    - "443:443"    # HTTPS
```

**After:**
```yaml
caddy:
  ports:
    - "8080:80"    # HTTP accessible on 8080
    - "8443:443"   # HTTPS accessible on 8443
```

This allows pesa-mind to run alongside your existing traefik setup without port conflicts.

---

## 🔧 Deployment Instructions

### Step 1: Stop Current Containers
```bash
docker-compose down
```

### Step 2: Redeploy with Fixed Ports
```bash
docker-compose up -d
```

### Step 3: Verify Deployment
```bash
# Check all containers are running
docker ps

# Verify pesa-mind is healthy
docker logs pesa-mind-app-1

# Test the API
curl http://localhost:8080/health
```

---

## 📊 Port Mapping Reference

### After Deployment

| Service | Internal | External | Status |
|---------|----------|----------|--------|
| **traefik** (existing) | 80 | 80 | ✅ Running |
| **traefik** (existing) | 443 | 443 | ✅ Running |
| **pesa-mind-caddy** (new) | 80 | **8080** | ✅ Available |
| **pesa-mind-caddy** (new) | 443 | **8443** | ✅ Available |
| **pesa-mind-app** (new) | 8080 | Internal | ✅ Running |
| **pesa-mind-db** (new) | 5432 | Internal | ✅ Running |

---

## 📝 API Access URLs

### HTTP
```
http://your-server:8080/api/v1/users/register
http://your-server:8080/health
```

### HTTPS
```
https://your-server:8443/api/v1/users/register
https://your-server:8443/health
```

---

## 🔐 Option: Use Traefik Instead of Caddy

If you want to use the existing traefik instead of adding Caddy, you can:

1. **Remove Caddy** from docker-compose.yml
2. **Configure traefik** to route requests to pesa-mind
3. **Access via** http://your-domain/pesa (if configured)

---

## ✅ Deployment Checklist

- [ ] Stopped old containers: `docker-compose down`
- [ ] Updated `docker-compose.yml` with new port mapping (8080, 8443)
- [ ] Redeployed: `docker-compose up -d`
- [ ] All containers running: `docker ps`
- [ ] App is healthy: `docker logs pesa-mind-app-1`
- [ ] Health check passes: `curl http://localhost:8080/health`
- [ ] Database is healthy: `docker logs pesa-mind-db-1`

---

## 🐛 Troubleshooting

### Still Getting Port Already Allocated Error?

```bash
# Find what's using the port
lsof -i :80
lsof -i :443

# Or check Docker
docker ps --all | grep 80
docker ps --all | grep 443

# Stop conflicting container if needed
docker stop <container-id>
```

### Caddy Container Not Starting?

```bash
# Check caddy logs
docker logs pesa-mind-caddy-1

# Verify Caddyfile exists
ls -la Caddyfile

# Check file syntax
docker run --rm -v $(pwd)/Caddyfile:/etc/caddy/Caddyfile:ro caddy:alpine caddy validate
```

### App Not Responding?

```bash
# Check app logs
docker logs pesa-mind-app-1

# Check database connection
docker logs pesa-mind-db-1

# Test database from app container
docker exec pesa-mind-app-1 psql -U pesamind -d pesamind -c "SELECT 1"
```

---

## 📚 Summary

✅ **Problem:** Port 80/443 already in use by traefik  
✅ **Solution:** Changed pesa-mind-caddy to use ports 8080/8443  
✅ **File Changed:** `docker-compose.yml`  
✅ **No Code Changes:** Only configuration updated  
✅ **Ready:** For immediate redeployment  

---

## 🚀 Next Steps

1. Run: `docker-compose down`
2. Run: `docker-compose up -d`
3. Verify: `docker ps`
4. Test: `curl http://localhost:8080/health`

**Deployment should succeed!** 🎉

---

**Last Updated:** April 9, 2026

