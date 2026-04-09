# Docker Healthcheck Fix - Deployment Instructions

## Problem
Your deployment has two unhealthy containers:
- `pesa-mind-app-1` (status: unhealthy)
- `pesa-mind-caddy-1` (status: unhealthy)

## Root Causes Fixed
1. **App Healthcheck Bug**: Was checking `http://localhost/health` but should use `http://localhost:8080/health` (missing port)
2. **Timing Issues**: Healthcheck intervals were too long (30s) with insufficient start_period
3. **Caddy Start Dependency**: Caddy wasn't properly waiting for app to be healthy
4. **Caddyfile Formatting**: Caddy logs warned about formatting issues

## Changes Made
✅ `docker-compose.yml`:
- Fixed app healthcheck to use correct port: `http://localhost:8080/health`
- Reduced intervals: 10s check with 5s timeout, 5 retries, 15s start_period
- Added explicit `service_healthy` dependency for caddy on app
- Improved caddy healthcheck timing: 10s interval, 60s start_period

✅ `Caddyfile`:
- Reformatted with proper indentation (tabs instead of spaces)
- Removed formatting warnings

✅ `Dockerfile.caddy` (NEW):
- Created custom Caddy image with required SSL/TLS tools
- Includes `nss-tools` (for certutil), curl, wget
- Eliminates warning: "certutil is not available"
- Allows proper certificate installation in system trust store

✅ `docker-compose.yml`:
- Updated caddy service to build from custom Dockerfile instead of using `caddy:alpine`

## Deployment Steps

### On your server (deploy@vmi3203905):

1. **Pull the latest changes**:
   ```bash
   cd ~/pesa-mind
   git pull origin main
   ```

2. **Rebuild and restart containers**:
   ```bash
   docker-compose down
   docker-compose build --no-cache app caddy
   docker-compose up -d
   ```

3. **Monitor the restart** (takes ~30-60s):
   ```bash
   # Watch status
   watch -n 2 'docker ps --format "table {{.Names}}\t{{.Status}}"'
   
   # Or check logs
   docker logs -f pesa-mind-app-1
   docker logs -f pesa-mind-caddy-1
   ```

4. **Verify health**:
   ```bash
   # App should show (healthy)
   docker exec pesa-mind-app-1 curl -f http://localhost:8080/health
   
   # Caddy should show (healthy) 
   docker exec pesa-mind-caddy-1 wget --spider http://localhost:2019/metrics
   ```

## Expected Timeline
- **0-15s**: App container starting, pre-warming
- **15-25s**: App passes healthcheck, marked healthy
- **25-60s**: Caddy waiting in start_period
- **60s+**: Caddy checks health and becomes healthy

## Troubleshooting

If containers still fail healthchecks:

1. **Check app logs for errors**:
   ```bash
   docker logs pesa-mind-app-1 | tail -50
   ```

2. **Check Caddy logs**:
   ```bash
   docker logs pesa-mind-caddy-1 | tail -50
   ```

3. **Manual health test**:
   ```bash
   docker exec pesa-mind-app-1 curl -v http://localhost:8080/health
   ```

4. **If database connection issues**:
   ```bash
   docker exec pesa-mind-db-1 pg_isready -U pesamind
   ```

## Verification Checklist
- [ ] All three containers show `healthy` in `docker ps`
- [ ] `curl -f http://localhost:8080/health` returns `{"status":"ok"}`
- [ ] `curl -f https://localhost/health` works (via Caddy)
- [ ] Application logs show no errors
- [ ] Caddy logs show no format warnings

## Support
If issues persist after deployment, check:
1. `.env` file is properly loaded
2. Database migrations completed successfully
3. No port conflicts on 8080, 8443, 5432

