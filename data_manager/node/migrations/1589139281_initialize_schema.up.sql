CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS files(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner text NOT NULL DEFAULT '',
    name text NOT NULL DEFAULT '',
    tags json,
    task_id text NOT NULL DEFAULT '',
    executable boolean NOT NULL DEFAULT false,
    hash text NOT NULL DEFAULT '',
    content_type text NOT NULL DEFAULT '',
    upload_token uuid DEFAULT uuid_generate_v4()
);

CREATE TABLE IF NOT EXISTS hash_ref_counts(
    hash text PRIMARY KEY,
    ref_count integer NOT NULL
);
