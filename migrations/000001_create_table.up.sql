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



CREATE TYPE role_type AS ENUM ('teacher','developer','admin');

-- USER TABLE
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    date_of_birth DATE,
    role role_type NOT NULL DEFAULT 'developer',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- SETTING TABLE
CREATE TABLE settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    privacy_level VARCHAR(50) NOT NULL DEFAULT 'private',
    notification VARCHAR(30) NOT NULL DEFAULT 'on',
    language VARCHAR(255) NOT NULL DEFAULT 'en',
    theme VARCHAR(255) NOT NULL DEFAULT 'light',
    updated_at TIMESTAMP DEFAULT NOW()
);

-- TOKEN
CREATE TABLE tokens (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);
