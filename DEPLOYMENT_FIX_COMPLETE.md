# PesaMind Docker Deployment - Complete Troubleshooting & Fix Guide

## Current Status
**Issue**: App container is unhealthy and preventing Caddy from starting.

## Root Causes & Solutions Applied

### 1. ✅ Missing `curl` in App Container
**Symptom**: Healthcheck command fails because curl doesn't exist
**Fix Applied**: Added `RUN apk add --no-cache curl` to Dockerfile

### 2. ✅ Insufficient Start Period
**Symptom**: Healthcheck runs before app is fully initialized
**Fix Applied**: Increased `start_period` from 15s → 30s

### 3. ✅ Dockerfile Casing Issue
**Symptom**: Build warning "FromAsCasing: 'as' and 'FROM' keywords' casing do not match"
**Fix Applied**: Changed `as` → `AS` on line 1

### 4. ✅ Missing Caddy SSL/TLS Tools
**Symptom**: Caddy warnings about missing certutil
**Fix Applied**: Created custom Dockerfile.caddy with nss-tools, curl, wget

## Complete Deployment Instructions

### For Your Server (deploy@vmi3203905:~/pesa-mind)

**Option A - Automatic (Recommended)**:
```bash
cd ~/pesa-mind
bash DEPLOY_NOW.sh
```

**Option B - Manual**:
```bash
cd ~/pesa-mind

# 1. Pull latest changes
git pull origin main

# 2. Stop and clean up
docker compose down -v

# 3. Rebuild
docker compose build --no-cache app caddy

# 4. Start
docker compose up -d

# 5. Monitor
docker compose logs -f
```

## Verification Checklist

After deployment, verify:

```bash
# Check container health
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# Should show all three as "healthy" or "running"
# pesa-mind-db-1    Up ... (healthy)
# pesa-mind-app-1   Up ... (healthy)
# pesa-mind-caddy-1 Up ... (healthy)
```

### Detailed Health Tests

```bash
# 1. Test database
docker exec pesa-mind-db-1 pg_isready -U pesamind
# Expected: accepting connections

# 2. Test app endpoint directly
docker exec pesa-mind-app-1 curl -v http://localhost:8080/health
# Expected: {"status":"ok"}

# 3. Test app via Caddy reverse proxy
curl -v http://localhost:8080/health
# Expected: {"status":"ok"}

# 4. Test Caddy health metrics
docker exec pesa-mind-caddy-1 wget --spider http://localhost:2019/metrics
# Expected: HTTP/1.1 200 OK
```

## If Issues Persist

### Check Logs

```bash
# App logs (where the main business logic runs)
docker logs pesa-mind-app-1 -f --tail=50

# Look for:
# ✅ "Listening and serving HTTP on :8080"
# ✅ "GET /health" requests being logged
# ❌ Any connection errors to database
# ❌ Any panic or fatal errors

# Caddy logs (reverse proxy)
docker logs pesa-mind-caddy-1 -f --tail=50

# Look for:
# ✅ "serving initial configuration"
# ✅ "http.log" entries showing requests proxied to app
# ❌ Any connection refused errors
# ❌ Any TLS errors

# Database logs
docker logs pesa-mind-db-1 -f --tail=50

# Look for:
# ✅ "database system is ready to accept connections"
# ❌ Any connection pool errors
```

### Common Issues & Solutions

**Issue: "Container pesa-mind-app-1 is unhealthy"**
```bash
# Cause: App not responding to health check within 30s
# Solution:
docker exec pesa-mind-app-1 sh
curl http://localhost:8080/health
# If this fails, check app logs: docker logs pesa-mind-app-1
```

**Issue: "connection refused" errors in Caddy**
```bash
# Cause: App not fully started when Caddy tries to connect
# Solution: Already handled - start_period: 30s ensures this
# If still happening, increase start_period to 45s in docker-compose.yml
```

**Issue: "curl: command not found"**
```bash
# Cause: curl not installed in container
# Solution: Already handled - added to Dockerfile
# Rebuild: docker compose build --no-cache app
```

**Issue: Database connection errors**
```bash
# Verify database is running and accepting connections
docker exec pesa-mind-db-1 pg_isready -U pesamind

# If fails, check .env file:
cat .env | grep DB_

# Should show:
# DB_HOST=db
# DB_PORT=5432
# DB_USER=pesamind
# DB_PASSWORD=pesamind123
# DB_NAME=pesamind
```

## Timeline of Container Startup

This is what **should** happen:

```
Time 0s:     Containers start
Time 0-5s:   Database initializing
Time 5-10s:  App starting, migrations running, compiled
Time 10-30s: App listening on :8080, NOT passing healthcheck yet (start_period)
Time 30s:    First healthcheck attempt
Time 30-35s: App passes healthcheck ✅
Time 35-60s: Caddy's start_period (waiting before its first check)
Time 60s:    Caddy's first healthcheck
Time 60-65s: Caddy becomes healthy ✅

TOTAL TIME: ~1 minute until all containers are healthy
```

## Files Modified in This Fix

| File | Change |
|------|--------|
| `Dockerfile` | Added curl; Fixed casing (as → AS) |
| `Dockerfile.caddy` | NEW: Custom image with SSL tools |
| `docker-compose.yml` | Fixed healthchecks, dependencies, timing |
| `Caddyfile` | Fixed formatting |
| `DEPLOY_NOW.sh` | NEW: Automated deployment script |

## Quick Reference Commands

```bash
# Deploy everything
cd ~/pesa-mind && bash DEPLOY_NOW.sh

# View real-time logs (all containers)
docker compose logs -f

# View app logs only
docker logs pesa-mind-app-1 -f

# Check specific container health
docker inspect pesa-mind-app-1 --format='{{json .State.Health}}'

# Force healthcheck to run now
docker exec pesa-mind-app-1 curl http://localhost:8080/health

# Rebuild just the app (faster than full build)
docker compose build --no-cache app && docker compose up -d app

# Clean everything and start fresh
docker compose down -v && docker compose up --build
```

## Support

If containers still don't become healthy after 2 minutes:

1. Check logs: `docker logs pesa-mind-app-1`
2. Test endpoint: `docker exec pesa-mind-app-1 curl http://localhost:8080/health`
3. Check database: `docker exec pesa-mind-db-1 pg_isready -U pesamind`
4. Check .env: `cat .env`

The issue will be in one of these areas.

