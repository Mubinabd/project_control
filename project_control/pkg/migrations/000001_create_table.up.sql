CREATE TABLE groups(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    swagger_url VARCHAR(255) NOT NULL,
    name VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE developers(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    group_id uuid NOT NULL REFERENCES groups(id),
    name VARCHAR(20) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    telegram_username VARCHAR(20) NOT NULL
);

CREATE TABLE private(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    swagger_url VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    telegram_username VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE documentation(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    group_id uuid  REFERENCES groups(id),
    private_id uuid REFERENCES private(id),
    title TEXT NOT NULL,
    description text not NULL,
    url VARCHAR(255) NOT NULL
);

