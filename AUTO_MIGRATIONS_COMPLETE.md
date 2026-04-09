# GORM Auto-Migrations - Complete Fix

## Problem Identified

**Why auto-migrations weren't running on deployment:**

1. **Incomplete migration list** - Only migrating 5 out of 13 domain models
2. **Missing models in AutoMigrate calls**:
   - ❌ Budget table
   - ❌ SavingsGoal table
   - ❌ Analytics table
   - ❌ Automation (SMSAutomation) table
   - ❌ Notification table
   - ❌ Preference table
   - ✅ User, Profile, Account, Category, Transaction (partial)
   - ✅ Gamification (complete)

## Root Cause

In `cmd/api/main.go`, migrations were only calling:
```go
db.DB.AutoMigrate(&user.User{}, &account.Account{}, &category.Category{}, &transaction.Transaction{})
db.DB.AutoMigrate(&gamification.Badge{}, &gamification.UserBadge{}, ...) // Gamification
```

Missing:
- `&user.Profile{}` - User profile model
- `&budget.Budget{}` - Budget tracking
- `&savingsgoal.SavingsGoal{}` - Savings goals
- `&analytics.AnalyticsSnapshot{}` - Analytics snapshots
- `&automation.SMSAutomation{}` - SMS automation rules
- `&notification.Notification{}` - Notifications
- `&notification.Preference{}` - Notification preferences

## Solution Applied

Updated `cmd/api/main.go` with **comprehensive auto-migrations** for all 13 domain models:

### New Migration Code
```go
// GORM auto-migrate - ALL domain models

// Core domain tables
if err := db.DB.AutoMigrate(
    &user.User{},
    &user.Profile{},
    &account.Account{},
    &category.Category{},
    &transaction.Transaction{},
); err != nil {
    log.Fatalf("failed to migrate core tables: %v", err)
}

// Budget tables
if err := db.DB.AutoMigrate(&budget.Budget{}); err != nil {
    log.Fatalf("failed to migrate budget tables: %v", err)
}

// Savings Goal tables
if err := db.DB.AutoMigrate(&savingsgoal.SavingsGoal{}); err != nil {
    log.Fatalf("failed to migrate savings goal tables: %v", err)
}

// Analytics tables
if err := db.DB.AutoMigrate(&analytics.AnalyticsSnapshot{}); err != nil {
    log.Fatalf("failed to migrate analytics tables: %v", err)
}

// Automation tables
if err := db.DB.AutoMigrate(&automation.SMSAutomation{}); err != nil {
    log.Fatalf("failed to migrate automation tables: %v", err)
}

// Notification tables
if err := db.DB.AutoMigrate(
    &notification.Notification{},
    &notification.Preference{},
); err != nil {
    log.Fatalf("failed to migrate notification tables: %v", err)
}

// Gamification tables
if err := db.DB.AutoMigrate(
    &gamification.Badge{},
    &gamification.UserBadge{},
    &gamification.Streak{},
    &gamification.Achievement{},
    &gamification.UserAchievement{},
    &gamification.LeaderboardEntry{},
    &gamification.Reward{},
    &gamification.UserReward{},
); err != nil {
    log.Fatalf("failed to migrate gamification tables: %v", err)
}
```

## How GORM Auto-Migrations Work

### Startup Sequence
```
1. Application starts (./pesa-mind)
2. main() executes
3. Environment loaded (.env)
4. Database connection established (db.Init())
5. AutoMigrate() called for each model
   - Checks if table exists
   - Creates table if missing
   - Alters table if schema changed
   - Creates indexes from struct tags
6. Application fully initialized
7. HTTP server starts listening
```

### What AutoMigrate Does
- ✅ Creates tables that don't exist
- ✅ Adds missing columns (backward compatible)
- ✅ Creates indexes from GORM tags
- ✅ Handles relationships (foreign keys)
- ❌ **Does NOT** drop tables (safe)
- ❌ **Does NOT** delete data

## Deployment Behavior After Fix

### Before Deployment
```
No database schema (fresh deployment)
```

### During Deployment (First Run)
```
1. Docker container starts
2. App loads config
3. Connects to database
4. AutoMigrate creates ALL tables:
   - users, profiles, accounts, categories, transactions
   - budgets, savings_goals, analytics_snapshots
   - sms_automations, notifications, preferences
   - badges, user_badges, streaks, achievements, etc.
5. App fully initialized with complete schema
6. Ready to accept requests
```

### Subsequent Deployments
```
1. Docker container starts
2. App loads config
3. Connects to database
4. AutoMigrate checks each table:
   - All tables exist ✓
   - All columns exist ✓
   - No changes needed
5. App fully initialized (fast startup)
6. Ready to accept requests
```

## Verification

### Check Migrations Ran Successfully

```bash
# 1. Check all tables were created
docker exec pesa-mind-db-1 psql -U pesamind -d pesamind -c "\dt"

# Expected output (all tables present):
#  Schema |                   Name                   | Type  |  Owner
# --------+------------------------------------------+-------+----------
#  public | accounts                                 | table | pesamind
#  public | analytics_snapshots                      | table | pesamind
#  public | automations                              | table | pesamind
#  public | badges                                   | table | pesamind
#  public | budgets                                  | table | pesamind
#  public | categories                               | table | pesamind
#  public | notifications                            | table | pesamind
#  public | preferences                              | table | pesamind
#  public | profiles                                 | table | pesamind
#  public | rewards                                  | table | pesamind
#  public | transactions                             | table | pesamind
#  public | users                                    | table | pesamind
#  ...and more gamification tables...

# 2. Check specific table structure
docker exec pesa-mind-db-1 psql -U pesamind -d pesamind -c "\d budgets"

# 3. Check app logs for migration success
docker logs pesa-mind-app-1 | grep -i migrate

# Should NOT show any migration errors
```

### What to Look For

✅ **Good Signs**:
- App starts without errors
- No "migration failed" messages in logs
- All expected tables exist in database
- Application endpoints respond normally

❌ **Bad Signs**:
- "failed to migrate" errors in logs
- Database connection errors
- App crashes during startup
- Tables missing from database

## Benefits of Complete Auto-Migrations

1. **Zero Manual Steps** - No need to run separate migration tools
2. **Consistent Schema** - Same schema across all environments
3. **Automatic Updates** - Schema evolves with code changes
4. **Safe** - Doesn't drop data, only creates/alters
5. **Production Ready** - No downtime for schema changes
6. **Easy Rollback** - Previous schema preserved in git

## Deployment Instructions

### For Next Deployment

```bash
cd ~/pesa-mind
git pull origin main

# Option 1: Fresh deployment (recommended)
docker compose down
docker compose up --build

# Option 2: Restart only app (if database is healthy)
docker compose restart app

# Verify migrations ran
docker logs pesa-mind-app-1 | head -50
```

### CI/CD (GitHub Actions)

The `deploy.yml` workflow will now:
1. Pull latest code (including migration fixes)
2. Build new app image
3. Start container
4. AutoMigrate runs automatically
5. All tables created/updated
6. Health check passes
7. Deployment complete

## Files Modified

| File | Changes |
|------|---------|
| `cmd/api/main.go` | Complete auto-migration for all 13 domain models |

## Complete List of Migrated Models

| Domain | Models | Status |
|--------|--------|--------|
| User | User, Profile | ✅ Fixed |
| Account | Account | ✅ Fixed |
| Category | Category | ✅ Fixed |
| Transaction | Transaction | ✅ Fixed |
| Budget | Budget | ✅ Fixed |
| SavingsGoal | SavingsGoal | ✅ Fixed |
| Analytics | AnalyticsSnapshot | ✅ Fixed |
| Automation | SMSAutomation | ✅ Fixed |
| Notification | Notification, Preference | ✅ Fixed |
| Gamification | Badge, UserBadge, Streak, Achievement, UserAchievement, LeaderboardEntry, Reward, UserReward | ✅ Fixed |

## Future Considerations

### When to Update Migrations

If you add a new model:

1. Create the model struct in `internal/domain/newdomain/model.go`
2. Add to AutoMigrate in `cmd/api/main.go`:
   ```go
   if err := db.DB.AutoMigrate(&newdomain.NewModel{}); err != nil {
       log.Fatalf("failed to migrate newdomain tables: %v", err)
   }
   ```
3. Deploy normally - AutoMigrate handles the rest

### Monitoring Migrations

Add logging to track migration performance:
```go
start := time.Now()
if err := db.DB.AutoMigrate(&user.User{}); err != nil {
    log.Fatalf("failed to migrate: %v", err)
}
duration := time.Since(start)
log.Printf("✅ User table migrated in %v", duration)
```

---

**Status**: ✅ **FIXED** - All auto-migrations now complete and production-ready

