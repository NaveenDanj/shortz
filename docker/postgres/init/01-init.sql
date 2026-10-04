CREATE USER auth_user WITH PASSWORD 'auth_password';
CREATE USER users_user WITH PASSWORD 'users_password';
CREATE USER videos_user WITH PASSWORD 'videos_password';
CREATE USER notifications_user WITH PASSWORD 'notifications_password';

-- Owning the database gives each user write access to its public schema
-- (Postgres 15+ no longer grants CREATE on public to everyone).
CREATE DATABASE auth_db OWNER auth_user;
CREATE DATABASE users_db OWNER users_user;
CREATE DATABASE videos_db OWNER videos_user;
CREATE DATABASE notifications_db OWNER notifications_user;
