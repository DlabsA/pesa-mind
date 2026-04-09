# ⚡ QUICK FIX - Caddy Let's Encrypt ACME Errors

## ❌ Problem
```
error: "tls: unrecognized name"
error: "Invalid response from https://api.dlabs.cc/.well-known/acme-challenge/...: 404"
```
Caddy can't get Let's Encrypt cert for staging IP.

## ✅ Solution
Use **self-signed certificates** for staging (no validation needed):

Updated `Caddyfile`:
```caddyfile
https://:443 {
    tls internal
    reverse_proxy app:8080
}
```

## 🚀 Deploy Now
```bash
cd ~/pesa-mind
git pull
./deploy.sh
```

## 🧪 Test
```bash
# HTTP
curl http://173.212.219.227:8080/health

# HTTPS (ignore self-signed warning)
curl -k https://173.212.219.227:8443/health
```

## 🔑 Key Changes
- ✅ Disabled ACME for staging
- ✅ Enabled self-signed certificates
- ✅ Works instantly without validation
- ✅ Production: Use ACME when domain ready

---

**Status:** ✅ Fixed - Staging uses self-signed certs

