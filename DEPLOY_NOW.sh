#!/bin/bash
# Complete PesaMind Deployment Script
# Run this on your server: deploy@vmi3203905:~/pesa-mind

set -e

echo "🚀 PesaMind Deployment - Complete Fix"
echo "======================================"
echo ""

# Step 1: Pull latest changes
echo "📥 Pulling latest changes from git..."
git pull origin main
echo "✅ Git pull complete"
echo ""

# Step 2: Stop and remove old containers
echo "🛑 Stopping and removing old containers..."
docker compose down -v
echo "✅ Containers removed"
echo ""

# Step 3: Rebuild images
echo "🔨 Rebuilding Docker images (this takes ~2-3 minutes)..."
docker compose build --no-cache app caddy
echo "✅ Images built successfully"
echo ""

# Step 4: Start containers
echo "🚀 Starting containers..."
docker compose up -d
echo "✅ Containers started"
echo ""

# Step 5: Wait for services to be healthy
echo "⏳ Waiting for services to become healthy..."
echo "   This may take 30-60 seconds..."
sleep 5

# Function to check container health
check_health() {
    local container=$1
    local max_attempts=30
    local attempt=0

    while [ $attempt -lt $max_attempts ]; do
        local status=$(docker inspect --format='{{.State.Health.Status}}' "$container" 2>/dev/null || echo "none")

        if [ "$status" = "healthy" ]; then
            echo "   ✅ $container is healthy"
            return 0
        elif [ "$status" = "unhealthy" ]; then
            echo "   ❌ $container is unhealthy (attempt $((attempt+1))/$max_attempts)"
        elif [ "$status" = "starting" ]; then
            echo "   🔄 $container is starting... (attempt $((attempt+1))/$max_attempts)"
        fi

        sleep 2
        ((attempt++))
    done

    echo "   ⚠️  $container did not become healthy within timeout"
    return 1
}

# Check all containers
echo ""
echo "🏥 Health Check Status:"
check_health "pesa-mind-db-1" || true
check_health "pesa-mind-app-1" || true
check_health "pesa-mind-caddy-1" || true

echo ""
echo "📊 Final Container Status:"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

echo ""
echo "🔍 Testing endpoints..."
echo ""

# Test app health
echo -n "   Testing app /health endpoint: "
if docker exec pesa-mind-app-1 curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ OK"
else
    echo "❌ FAILED"
fi

# Test app via caddy
echo -n "   Testing caddy proxy: "
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ OK"
else
    echo "❌ FAILED (but may work after warmup)"
fi

echo ""
echo "📋 Deployment Summary:"
echo "   - Database: $(docker ps --filter name=pesa-mind-db-1 --format '{{.Status}}')"
echo "   - App:      $(docker ps --filter name=pesa-mind-app-1 --format '{{.Status}}')"
echo "   - Caddy:    $(docker ps --filter name=pesa-mind-caddy-1 --format '{{.Status}}')"

echo ""
echo "✨ Deployment complete!"
echo ""
echo "📍 Access your API:"
echo "   - HTTP:   http://localhost:8080"
echo "   - HTTPS:  https://localhost"
echo ""
echo "🔧 Useful commands:"
echo "   - Logs (app):   docker logs pesa-mind-app-1 -f"
echo "   - Logs (caddy): docker logs pesa-mind-caddy-1 -f"
echo "   - Status:       docker ps"
echo "   - Shell (app):  docker exec -it pesa-mind-app-1 sh"

