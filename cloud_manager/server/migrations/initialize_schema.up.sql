CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS credentials (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    display_name text NOT NULL,
    provider text NOT NULL CHECK ((provider == "aws") AND (aws_access_key_id IS NOT NULL) AND ()),
    aws_access_key_id text,
    aws_access_key text
);

ALTER TABLE credentials
ADD CONSTRAINT check_aws_credentials_coherence CHECK ((provider == "aws") == ((aws_access_key_id IS NOT NULL) AND (aws_access_key IS NOT NULL)))

CREATE TABLE IF NOT EXISTS cloud_resource_factories (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    token text NOT NULL UNIQUE,
    display_name text NOT NULL,
    provider text NOT NULL,
    agent_docker_version text NOT NULL,
    type text NOT NULL,
    cpu_inst_limit_cores numeric NOT NULL,
    memory_inst_limit_bytes bigint NOT NULL,
    disk_inst_limit_bytes bigint NOT NULL,
    cpu_fact_limit_cores numeric NOT NULL,
    memory_fact_limit_bytes bigint NOT NULL,
    disk_fact_limit_bytes bigint NOT NULL,
    cpu_used_cores numeric NOT NULL,
    memory_used_bytes bigint NOT NULL,
    disk_used_bytes bigint NOT NULL,
    inst_count_limit integer NOT NULL,
    inst_count_used integer NOT NULL,
    aws_cluster text,
    aws_vpc text,
    aws_tag text,
    aws_nets text[],
    aws_sec_groups text[]
);

CREATE TABLE IF NOT EXISTS cloud_resources (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    factory_id uuid REFERENCES cloud_resource_factories(id) ON DELETE CASCADE,
    token text NOT NULL UNIQUE,
    display_name text NOT NULL,
    status text NOT NULL,
    cpu_limit_cores numeric NOT NULL,
    memory_limit_bytes bigint NOT NULL,
    disk_limit_bytes bigint NOT NULL,
    cpu_guarantee_cores numeric NOT NULL,
    memory_guarantee_bytes bigint NOT NULL,
    disk_guarantee_bytes bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id text UNIQUE,
    expire_at timestamp NOT NULL,
    status text NOT NULL
);
