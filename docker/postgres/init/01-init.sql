CREATE DATABASE auth_db;
CREATE DATABASE users_db;
CREATE DATABASE videos_db;
CREATE DATABASE notifications_db;

CREATE USER auth_user WITH PASSWORD 'auth_password';
CREATE USER users_user WITH PASSWORD 'users_password';
CREATE USER videos_user WITH PASSWORD 'videos_password';
CREATE USER notifications_user WITH PASSWORD 'notifications_password';


GRANT ALL PRIVILEGES ON DATABASE auth_db TO auth_user;
GRANT ALL PRIVILEGES ON DATABASE users_db TO users_user;
GRANT ALL PRIVILEGES ON DATABASE videos_db TO videos_user;
GRANT ALL PRIVILEGES ON DATABASE notifications_db TO notifications_user;