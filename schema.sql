CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS posts
(
  id uuid NOT NULL DEFAULT uuid_generate_v1(),
  title varchar(255) NOT NULL,
  description varchar(255) DEFAULT NULL,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NULL DEFAULT NULL

);
