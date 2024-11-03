CREATE TABLE
  "users" (
    id text NOT NULL,
    username text NOT NULL UNIQUE,
    password text,
    active int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by text NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by text NOT NULL,
    PRIMARY KEY (id)
  );

CREATE TABLE
  "tokens" (
    id text NOT NULL,
    value text NOT NULL UNIQUE,
    active int NOT NULL,
    expires TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    created_by text NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by text NOT NULL,
    user_id text NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
  );

create table
  "workflows" (
    id text,
    name text not null,
    code text,
    object_type text not null,
    fail_on_error BOOLEAN NOT NULL,
    build_in BOOLEAN NOT NULL,
    created_at TIMESTAMP not null,
    modified_at TIMESTAMP not null,
    created_by text not null,
    modified_by text not null,
    deleted_at TIMESTAMP null,
    PRIMARY KEY (id),
    UNIQUE (name, deleted_at)
  );

create table
  "runs" (
    id text,
    workflow_id text not null,
    error text,
    start_time TIMESTAMP not null,
    finish_time TIMESTAMP not null,
    created_at TIMESTAMP not null,
    modified_at TIMESTAMP not null,
    created_by text not null,
    modified_by text not null,
    PRIMARY KEY (id),
    FOREIGN KEY (workflow_id) REFERENCES workflows (id)
  );

CREATE TABLE
  "workflows_children" (
    workflow_id text NOT NULL,
    children_id text NOT NULL,
    PRIMARY KEY (children_id, workflow_id),
    FOREIGN KEY (children_id) REFERENCES workflows (id) ON DELETE CASCADE,
    FOREIGN KEY (workflow_id) REFERENCES workflows (id) ON DELETE CASCADE
  );

CREATE TABLE
  "run_logs" (
    id text,
    message text NOT NULL,
    timestamp TIMESTAMP not null,
    run_id text NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (run_id) REFERENCES runs (id) ON DELETE CASCADE
  );

CREATE UNIQUE INDEX workflows_name_deleted_at_unique ON workflows (name)
WHERE
  deleted_at IS NULL;