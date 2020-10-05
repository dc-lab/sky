CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE files (
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

CREATE TABLE hash_counts (
    hash text PRIMARY KEY,
    ref_count integer NOT NULL DEFAULT 0
);

CREATE TABLE nodes (
    location text PRIMARY KEY,
    free_space integer NOT NULL DEFAULT 0,
    report_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE hash_targets (
    hash text NOT NULL REFERENCES hash_counts ON DELETE CASCADE,
    location text NOT NULL REFERENCES nodes ON DELETE CASCADE
);

CREATE TABLE hash_locations (
    hash text NOT NULL REFERENCES hash_counts ON DELETE CASCADE,
    location text NOT NULL REFERENCES nodes ON DELETE CASCADE,
    report_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
