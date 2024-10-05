CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE access_tokens (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "client_id" UUID NOT NULL,  -- Assuming client_id is a UUID
    "client" VARCHAR(200) NOT NULL,  -- Increased length for token
    "user_id" UUID NOT NULL,     -- Assuming user_id is a UUID
    "user" VARCHAR(200) NOT NULL,  -- Increased length for token
    "token" VARCHAR(300) NOT NULL,  -- Increased length for token
    "expires_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL,  -- Expiration timestamp
    "scope" VARCHAR(50) NOT NULL,  -- Consider if this should be unique
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Indexes for performance
CREATE INDEX idx_access_tokens_client_id ON access_tokens("client_id");
CREATE INDEX idx_access_tokens_user_id ON access_tokens("user_id");
CREATE INDEX idx_access_tokens_token ON access_tokens("token");  -- Optional: index for token lookups
