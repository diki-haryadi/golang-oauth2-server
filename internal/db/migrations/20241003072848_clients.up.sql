CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE clients (
       "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
       "api_key" VARCHAR(254) NOT NULL UNIQUE,  -- Renamed to avoid reserved keyword
       "secret" VARCHAR(128) NOT NULL,            -- Increased length for security
       "redirect_uri" VARCHAR(200) NOT NULL,      -- Consider adding a check constraint if needed
       "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL
    );

-- Optional: Create an index on redirect_uri if you plan to query by it
CREATE INDEX idx_users_redirect_uri ON clients("redirect_uri");
