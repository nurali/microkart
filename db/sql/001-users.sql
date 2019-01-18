CREATE TABLE IF NOT EXISTS users (
    id uuid primary key DEFAULT uuid_generate_v4() NOT NULL,
    name text,
    passwd text
);
