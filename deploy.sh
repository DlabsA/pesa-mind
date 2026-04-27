#!/bin/bash

# PesaMind Deployment Script
# Handles port conflicts and redeploys the application

set -e

echo "🚀 PesaMind Deployment Script"
echo "=============================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Step 1: Stop current deployment
echo -e "${YELLOW}Step 1: Stopping current containers...${NC}"
docker compose down -v && docker compose up --build

# Step 4: Wait for containers to be healthy
echo -e "${YELLOW}Step 4: Waiting for containers to be healthy...${NC}"
sleep 10

# Check database
echo "Checking database..."
for i in {1..30}; do
    if docker compose exec -T db pg_isready -U pesamind > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Database is healthy${NC}"
        break
    fi
    if [ $i -eq 30 ]; then
        echo -e "${RED}❌ Database failed to start${NC}"
        exit 1
    fi
    echo "Waiting... ($i/30)"
    sleep 1
done
echo ""

# Check app
echo "Checking application..."
for i in {1..30}; do
    if curl -s http://localhost:8080/health > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Application is healthy${NC}"
        break
    fi
    if [ $i -eq 30 ]; then
        echo -e "${RED}❌ Application failed to start${NC}"
        docker logs pesa-mind-app-1
        exit 1
    fi
    echo "Waiting... ($i/30)"
    sleep 1
done
echo ""

# Step 5: Verify all containers
echo -e "${YELLOW}Step 5: Verifying all containers...${NC}"
docker ps
echo ""

# Step 6: Display access information
echo -e "${GREEN}✅ DEPLOYMENT SUCCESSFUL!${NC}"
echo ""
echo "📊 Access Information:"
echo "====================="
echo -e "HTTP:  ${GREEN}http://your-server:8080${NC}"
echo -e "HTTPS: ${GREEN}https://your-server:8443${NC}"
echo ""
echo "🏥 Health Check:"
echo "================"
curl -s http://localhost:8080/health | jq '.' 2>/dev/null || echo "Health endpoint: http://localhost:8080/health"
echo ""
echo "📝 Container Status:"
echo "==================="
docker ps
echo ""
echo "🎉 Ready to use!"

