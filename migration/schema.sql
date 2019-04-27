CREATE DATABASE golangw2;

\c golangw2

CREATE TABLE IF NOT EXISTS secret(
    hash character varying(255) NOT NULL,
    secret_text character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    remaining_views int NOT NULL
);

-- psql -h localhost -U postgres -a -f migration/schema.sql