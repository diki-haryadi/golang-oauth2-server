CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
   "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
   "username" VARCHAR(254) UNIQUE NOT NULL,
   "password" VARCHAR(60) NOT NULL,
   "role" VARCHAR(50) NOT NULL,  -- Consider ENUM if roles are fixed
   "role_id" VARCHAR(50) NOT NULL UNIQUE,  -- If role_id should be unique
   "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL
--  CONSTRAINT unique_username_role UNIQUE ("username", "role")  -- Optional: if a user can have only one role
);
