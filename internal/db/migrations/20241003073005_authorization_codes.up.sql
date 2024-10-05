CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE authorization_codes (
   "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
   "client_id" UUID NOT NULL,  -- Assuming client_id is a UUID
   "client" VARCHAR(200) NOT NULL,  -- Increased length for token
   "user_id" UUID NOT NULL,     -- Assuming user_id is a UUID
   "user" VARCHAR(200) NOT NULL,  -- Increased length for token
   "code" VARCHAR(300) NOT NULL,  -- Increased length for token
   "redirect_uri" VARCHAR(200) NOT NULL,
   "expires_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL,  -- Expiration timestamp
   "scope" VARCHAR(50) NOT NULL,  -- Consider if this should be unique
   "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Indexes for performance
CREATE INDEX idx_authorization_codes_client_id ON authorization_codes("client_id");
CREATE INDEX idx_authorization_codes_user_id ON authorization_codes("user_id");
CREATE INDEX idx_authorization_codes_token ON authorization_codes("token");  -- Optional: index for token lookups
