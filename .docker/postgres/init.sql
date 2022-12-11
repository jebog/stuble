CREATE USER devops;
ALTER USER devops WITH ENCRYPTED password 'devops';

CREATE DATABASE stuble OWNER devops;
CREATE DATABASE keycloak OWNER devops;
