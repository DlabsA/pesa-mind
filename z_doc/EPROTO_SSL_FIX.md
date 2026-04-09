# 🔧 QUICK FIX - EPROTO SSL Error

## ❌ Error
```
EPROTO: WRONG_VERSION_NUMBER (SSL Error)
https://173.212.219.227:8080  ← WRONG PORT FOR HTTPS
```

## ✅ Fix
Use correct ports:

```
HTTP:  http://173.212.219.227:8080
HTTPS: https://173.212.219.227:8443
```

## 📝 Correct Request

### Before (❌ Wrong)
```bash
curl https://173.212.219.227:8080/api/v1/auth/login
# EPROTO: WRONG_VERSION_NUMBER
```

### After (✅ Correct)
```bash
curl -k https://173.212.219.227:8443/api/v1/auth/login
# Works! ✅
```

## 🔑 Key Points

- **Port 8080** = HTTP only
- **Port 8443** = HTTPS with SSL
- **-k flag** = Ignore self-signed cert (testing only)

## 🚀 Redeploy

```bash
cd ~/pesa-mind
git pull
./deploy.sh
```

## 🧪 Test

```bash
# HTTP test
curl http://173.212.219.227:8080/health

# HTTPS test
curl -k https://173.212.219.227:8443/health
```

---

**Status:** ✅ Fixed - Use port 8443 for HTTPS

