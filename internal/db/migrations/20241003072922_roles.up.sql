CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
   "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
   "name" VARCHAR(200) NOT NULL UNIQUE,  -- Unique role name
   "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  -- Timestamp of creation
   "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  -- Timestamp of last update
   "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL  -- Timestamp of deletion
);

-- Optional: Create an index on 'name' if you plan to search by it frequently
CREATE INDEX idx_roles_name ON roles("name");
