DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_advertisement_created_at;
DROP INDEX IF EXISTS idx_vacant_positions_created_at;
DROP INDEX IF EXISTS idx_methodological_material_created_at;
DROP INDEX IF EXISTS idx_reviews_rating;
DROP INDEX IF EXISTS idx_reviews_created_at;
DROP INDEX IF EXISTS idx_vacant_positions_date;
DROP INDEX IF EXISTS idx_methodological_material_date;

DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS advertisement;
DROP TABLE IF EXISTS vacant_positions;
DROP TABLE IF EXISTS methodological_material;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;

