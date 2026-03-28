-- Seed badges and achievements with known UUIDs for gamification triggers

-- Badges
INSERT INTO badges (id, name, description, icon, created_at, updated_at)
VALUES
  ('11111111-1111-1111-1111-111111111111', 'First Transaction', 'Awarded for creating your first transaction', 'first_tx.png', NOW(), NOW()),
  ('33333333-3333-3333-3333-333333333333', '7-Day Streak', 'Awarded for logging transactions 7 days in a row', 'streak7.png', NOW(), NOW()),
  ('44444444-4444-4444-4444-444444444444', 'Savings Goal Completed', 'Awarded for completing a savings goal', 'goal.png', NOW(), NOW());

-- Achievements
INSERT INTO achievements (id, name, description, icon, created_at, updated_at)
VALUES
  ('22222222-2222-2222-2222-222222222222', 'Goal Achiever', 'Awarded for achieving a savings goal', 'goal_achiever.png', NOW(), NOW()),
  ('55555555-5555-5555-5555-555555555555', 'Super Saver', 'Awarded for saving a significant amount', 'super_saver.png', NOW(), NOW());

