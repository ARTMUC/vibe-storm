-- Down migration for VibeStorm Event Storming to UML Generator application

-- Drop tables in reverse order of creation to respect foreign key constraints

DROP TABLE IF EXISTS user_favorites;
DROP TABLE IF EXISTS project_invitations;
DROP TABLE IF EXISTS ai_processing_logs;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS uml_diagrams;
DROP TABLE IF EXISTS note_relationships;
DROP TABLE IF EXISTS sticky_notes;
DROP TABLE IF EXISTS project_users;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS users;
