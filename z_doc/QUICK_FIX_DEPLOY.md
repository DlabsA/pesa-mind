# ⚡ QUICK FIX - Deploy Script Docker Compose Error

## 🔴 Error
```
docker-compose: command not found
Error: Process completed with exit code 127.
```

## ✅ Fixed!

The `deploy.sh` script now automatically detects your Docker Compose version:
- ✅ Uses `docker compose` for Docker v2 (your server)
- ✅ Falls back to `docker-compose` for Docker v1
- ✅ Works with GitHub Actions and manual deployment

## 🚀 Deploy Now (Choose One)

### Option A: Manual SSH (Fastest)
```bash
ssh deploy@your-server
cd ~/pesa-mind
git pull
chmod +x deploy.sh
./deploy.sh
```

### Option B: GitHub Actions (Automated)
```bash
git tag v1.0.1
git push origin v1.0.1
```

## ✨ What Changed

| File | Change |
|------|--------|
| `deploy.sh` | Auto-detects Docker Compose version |
| `deploy.yml` | Sets CI=true for non-interactive mode |

## 🧪 Verify

```bash
curl http://your-server:8080/health
# Should return: {"status":"ok"}
```

---

**Status:** ✅ READY TO DEPLOY

