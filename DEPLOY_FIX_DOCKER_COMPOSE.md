# 🔧 FIX: Deploy.sh Docker Compose Command Error

## 🔴 Problem
```
./***.sh: line 20: docker-compose: command not found
Error: Process completed with exit code 127.
```

## ✅ Root Cause
Your server has **Docker v2** installed, which uses `docker compose` (new syntax) instead of `docker-compose` (old syntax).

Modern Docker installations don't have the `docker-compose` command - they use `docker compose` instead.

---

## ✅ Solution Applied

### Updated Files

#### 1. **deploy.sh** 
- ✅ Now detects Docker Compose version automatically
- ✅ Uses `docker compose` if available (v2)
- ✅ Falls back to `docker-compose` if available (v1)
- ✅ Works with CI/CD environments (non-interactive)

**Key Changes:**
```bash
# BEFORE (Fails on Docker v2)
docker-compose down

# AFTER (Detects and uses correct version)
if docker compose version &> /dev/null; then
    DOCKER_COMPOSE="docker compose"
elif command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
fi
$DOCKER_COMPOSE down
```

#### 2. **.github/workflows/deploy.yml**
- ✅ Sets `CI=true` environment variable
- ✅ Makes deploy.sh non-interactive for GitHub Actions
- ✅ Proper error handling and logging

**Key Addition:**
```yaml
export CI=true
./deploy.sh
```

---

## 🚀 How to Redeploy

### Option 1: Manual SSH Deployment
```bash
# SSH to your server
ssh deploy@your-server

# Navigate to project
cd ~/pesa-mind

# Pull latest fixed deploy.sh
git pull origin main

# Make executable
chmod +x deploy.sh

# Run deployment
./deploy.sh
```

### Option 2: GitHub Actions Tag (Automatic)
```bash
# On your local machine
git add -A
git commit -m "Fix: docker compose command detection"
git tag v1.0.1
git push origin v1.0.1
```

GitHub Actions will automatically run the deployment with the fixed script.

---

## 🧪 Test the Fix

### On Your Server
```bash
cd ~/pesa-mind

# Test docker compose detection
if docker compose version &> /dev/null; then
    echo "✅ docker compose v2 detected"
else
    echo "❌ docker compose not found"
fi

# Run deployment
./deploy.sh
```

### Expected Output
```
🚀 PesaMind Deployment Script
Using: docker compose
Step 1: Stopping current containers...
✅ Containers stopped
Step 2: Cleaning up...
✅ Volumes retained (CI/CD mode)
Step 3: Building and deploying new containers...
✅ Containers deployed
...
✅ DEPLOYMENT SUCCESSFUL!
```

---

## 📊 What Changed

| Item | Before | After |
|------|--------|-------|
| Command | `docker-compose` (hardcoded) | `docker compose` (detected) |
| Compatibility | ❌ Docker v2 fails | ✅ Docker v2 works |
| Fallback | ❌ None | ✅ Falls back to v1 if available |
| CI/CD Support | ❌ Interactive mode | ✅ Non-interactive mode |
| Error Handling | ❌ Fails silently | ✅ Clear error messages |

---

## ✅ Verification Checklist

After deployment, verify:

- [ ] Script runs without errors
- [ ] Containers deploy successfully: `docker ps`
- [ ] App is healthy: `docker logs pesa-mind-app-1`
- [ ] Database is healthy: `docker logs pesa-mind-db-1`
- [ ] Health endpoint works: `curl http://localhost:8080/health`

---

## 🎯 Next Steps

1. **Pull the fix**: `git pull origin main`
2. **Test locally**: `./deploy.sh` (if you have Docker)
3. **Deploy via SSH**: Manual SSH deployment (recommended for immediate fix)
4. **Or use GitHub Actions**: Push a new tag to trigger automatic deployment

---

## 📚 Reference

**Docker Compose Versions:**
- **v1** (old): `docker-compose` command
- **v2** (new): `docker compose` command

**Your Server Status:**
- ✅ Has Docker v2 (uses `docker compose`)
- ✅ No `docker-compose` binary available

---

## 🚀 Deploy Now

### Quick Deploy (30 seconds)
```bash
# SSH to server
ssh deploy@your-server

# Update and deploy
cd ~/pesa-mind && git pull && chmod +x deploy.sh && ./deploy.sh
```

---

**Status:** ✅ Fixed and Ready  
**Files Updated:** deploy.sh, .github/workflows/deploy.yml  
**Backward Compatible:** Yes (supports both v1 and v2)

