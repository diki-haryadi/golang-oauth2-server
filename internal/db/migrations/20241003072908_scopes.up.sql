CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE scopes (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "scope" VARCHAR(200) NOT NULL UNIQUE,
    "description" VARCHAR(300) NOT NULL,
    "is_default" VARCHAR(200) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Optional: Create an index on 'description' if you plan to search by it frequently
CREATE INDEX idx_scopes_description ON scopes("description");
