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

# Detect docker compose command (v2 or v1)
if command -v docker &> /dev/null; then
    if docker compose version &> /dev/null; then
        DOCKER_COMPOSE="docker compose"
    elif command -v docker-compose &> /dev/null; then
        DOCKER_COMPOSE="docker-compose"
    else
        echo -e "${RED}❌ Docker Compose not found!${NC}"
        exit 1
    fi
else
    echo -e "${RED}❌ Docker not found!${NC}"
    exit 1
fi

echo -e "${YELLOW}Using: $DOCKER_COMPOSE${NC}"
echo ""

# Step 1: Stop current deployment
echo -e "${YELLOW}Step 1: Stopping current containers...${NC}"
$DOCKER_COMPOSE down || true
echo -e "${GREEN}✅ Containers stopped${NC}"
echo ""

# Step 2: Remove old volumes if needed (non-interactive for CI/CD)
echo -e "${YELLOW}Step 2: Cleaning up...${NC}"
if [ -z "$CI" ]; then
    # Interactive mode (local)
    read -p "Remove old volumes? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        $DOCKER_COMPOSE down -v
        echo -e "${GREEN}✅ Volumes removed${NC}"
    else
        echo -e "${GREEN}✅ Volumes retained${NC}"
    fi
else
    # CI/CD mode - always keep volumes
    echo -e "${GREEN}✅ Volumes retained (CI/CD mode)${NC}"
fi
echo ""

# Step 3: Rebuild and deploy
echo -e "${YELLOW}Step 3: Building and deploying new containers...${NC}"
$DOCKER_COMPOSE up -d
echo -e "${GREEN}✅ Containers deployed${NC}"
echo ""

# Step 4: Wait for containers to be healthy
echo -e "${YELLOW}Step 4: Waiting for containers to be healthy...${NC}"
sleep 10

# Check database
echo "Checking database..."
for i in {1..30}; do
    if $DOCKER_COMPOSE exec -T db pg_isready -U pesamind > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Database is healthy${NC}"
        break
    fi
    if [ $i -eq 30 ]; then
        echo -e "${RED}❌ Database failed to start${NC}"
        $DOCKER_COMPOSE logs db
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
        $DOCKER_COMPOSE logs pesa-mind-app-1
        exit 1
    fi
    echo "Waiting... ($i/30)"
    sleep 1
done
echo ""

# Step 5: Verify all containers
echo -e "${YELLOW}Step 5: Verifying all containers...${NC}"
docker ps --filter "name=pesa-mind"
echo ""

# Step 6: Display access information
echo -e "${GREEN}✅ DEPLOYMENT SUCCESSFUL!${NC}"
echo ""
echo "📊 Access Information:"
echo "====================="
echo -e "HTTP:  ${GREEN}http://localhost:8080${NC}"
echo -e "HTTPS: ${GREEN}https://localhost:8443${NC}"
echo ""
echo "🏥 Health Check:"
echo "================"
curl -s http://localhost:8080/health | jq '.' 2>/dev/null || curl -s http://localhost:8080/health
echo ""
echo "📝 Container Status:"
echo "==================="
$DOCKER_COMPOSE ps
echo ""
echo "🎉 Ready to use!"



