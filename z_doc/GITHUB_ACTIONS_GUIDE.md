# 🚀 GitHub Actions Deployment Workflow

## Overview

The updated `.github/workflows/deploy.yml` now uses the automated `deploy.sh` script for streamlined deployments.

---

## Deployment Flow

### 1. **Trigger**
Deployment automatically runs when you push a tag:
```bash
git tag v1.0.0
git push origin v1.0.0
```

### 2. **GitHub Actions Steps**

#### Step 1: Checkout Code
```yaml
- name: Checkout code
  uses: actions/checkout@v4
```
Pulls your latest code from the repository.

#### Step 2: Set Up SSH
```yaml
- name: Set up SSH
  run: |
    mkdir -p ~/.ssh
    echo "${{ secrets.SSH_PRIVATE_KEY }}" > ~/.ssh/id_rsa
    chmod 600 ~/.ssh/id_rsa
    ssh-keyscan -H ${{ secrets.SERVER_HOST }} >> ~/.ssh/known_hosts
```
Configures SSH access to your deployment server using GitHub secrets.

#### Step 3: Deploy via SSH
```yaml
- name: Deploy via SSH
  run: |
    ssh ... << 'EOF'
    cd ~/pesa-mind
    git pull origin main
    chmod +x deploy.sh
    ./deploy.sh
    EOF
```
- Pulls latest code
- Makes deploy.sh executable
- Runs the automated deployment script

#### Step 4: Health Check
```yaml
- name: Health check
  run: |
    sleep 10
    curl -f http://${{ secrets.SERVER_HOST }}:8080/health
```
Waits 10 seconds, then verifies the application is responding.

#### Step 5: Success Summary
Displays deployment URLs and helpful commands if successful.

#### Step 6: Failure Info
Shows debug steps if deployment fails.

---

## Required GitHub Secrets

Set these in your GitHub repository settings: `Settings > Secrets and variables > Actions`

| Secret | Description | Example |
|--------|-------------|---------|
| `SSH_PRIVATE_KEY` | Private SSH key for server | `-----BEGIN OPENSSH PRIVATE KEY-----...` |
| `SERVER_HOST` | Server hostname/IP | `vmi3203905.contaboserver.com` |
| `SERVER_USER` | SSH username | `deploy` |

---

## How to Set Up Secrets

### 1. Generate SSH Key (if you don't have one)
```bash
ssh-keygen -t ed25519 -f ~/.ssh/pesa-mind -C "github-actions"
```

### 2. Add Public Key to Server
```bash
cat ~/.ssh/pesa-mind.pub | ssh deploy@your-server "cat >> ~/.ssh/authorized_keys"
```

### 3. Add to GitHub Secrets
1. Go to: `https://github.com/YOUR_USERNAME/pesa-mind/settings/secrets/actions`
2. Click "New repository secret"
3. Add each secret:
   - Name: `SSH_PRIVATE_KEY`, Value: `cat ~/.ssh/pesa-mind`
   - Name: `SERVER_HOST`, Value: `your-server-ip-or-hostname`
   - Name: `SERVER_USER`, Value: `deploy`

---

## What `deploy.sh` Does

The deployment script automatically:

1. ✅ Stops old containers: `docker-compose down`
2. ✅ Rebuilds images: `docker build`
3. ✅ Starts new containers: `docker-compose up -d`
4. ✅ Waits for database: Polls until PostgreSQL is healthy
5. ✅ Waits for application: Polls until app responds to health check
6. ✅ Displays final status: Shows all running containers
7. ✅ Verifies health: Calls health endpoint
8. ✅ Shows access info: Displays deployment URLs

---

## Deployment Example

### Tag and Push
```bash
git tag v1.2.3
git push origin v1.2.3
```

### GitHub Actions Runs
Workflow executes with output like:
```
🚀 PesaMind Deployment Script
✅ Containers stopped
✅ Containers deployed
Waiting for containers to be healthy...
✅ Database is healthy
✅ Application is healthy
✅ DEPLOYMENT SUCCESSFUL!

📊 Access Information:
HTTP:  http://your-server:8080
HTTPS: https://your-server:8443

🏥 Health Check:
{"status":"ok"}

📝 Container Status:
CONTAINER ID   IMAGE            STATUS
17847067c099   pesa-mind-app    Up 57 seconds (healthy)
1f3e50bd1b17   postgres:16      Up About a minute (healthy)
```

---

## Manual Deployment (Without GitHub)

If you want to deploy without pushing a tag:

```bash
# SSH to server
ssh deploy@your-server

# Navigate to project
cd ~/pesa-mind

# Make script executable
chmod +x deploy.sh

# Run deployment
./deploy.sh
```

---

## Monitoring Deployment

### Check Real-Time Logs
```bash
ssh deploy@your-server
docker logs -f pesa-mind-app-1
```

### View Deployment Status
```bash
ssh deploy@your-server
docker ps
```

### Test API After Deployment
```bash
curl http://your-server:8080/health
curl -X POST http://your-server:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456"}'
```

---

## Troubleshooting

### Deployment Failed in GitHub Actions

1. **Check GitHub Actions logs**: `https://github.com/YOUR_USERNAME/pesa-mind/actions`
2. **Click the failed workflow**
3. **Check "Deploy via SSH" step** for SSH/connection errors
4. **Check "Health check" step** for application errors

### SSH Connection Issues
- ✅ Verify SSH keys are correct
- ✅ Verify server host is reachable
- ✅ Verify deploy user exists
- ✅ Test SSH locally: `ssh deploy@your-server`

### Application Won't Start
```bash
ssh deploy@your-server
docker logs pesa-mind-app-1
docker logs pesa-mind-db-1
```

### Port Already in Use
Port 8080/8443 conflict? Check:
```bash
docker ps
lsof -i :8080
```

---

## Best Practices

1. **Tag Releases**: Use semantic versioning
   ```bash
   git tag v1.0.0
   git tag v1.0.1
   git tag v1.1.0
   ```

2. **Test Locally First**:
   ```bash
   docker-compose down
   ./deploy.sh
   curl http://localhost:8080/health
   ```

3. **Monitor After Deploy**:
   ```bash
   ssh deploy@your-server
   docker logs -f pesa-mind-app-1
   ```

4. **Keep Secrets Secure**:
   - Never commit SSH keys
   - Rotate keys periodically
   - Use separate deploy user

---

## Workflow File Location
```
.github/workflows/deploy.yml
```

---

## Summary

✅ **Automated Deployment** via GitHub Actions tags  
✅ **Automated Health Checks** in deploy.sh  
✅ **SSH Key Authentication** with GitHub Secrets  
✅ **Comprehensive Logging** at each step  
✅ **Failure Notifications** with debug info  
✅ **Manual Fallback** with deploy.sh script  

---

**Status:** ✅ Ready for Production  
**Updated:** April 9, 2026

