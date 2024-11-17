CREATE TABLE
  "users" (
    id TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT,
    active int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by TEXT NOT NULL,
    PRIMARY KEY (id)
  );

CREATE TABLE
  "tokens" (
    id TEXT NOT NULL,
    value TEXT NOT NULL,
    active int NOT NULL,
    expires TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by TEXT NOT NULL,
    user_id TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
  );

create table
  "workflows" (
    id TEXT,
    name TEXT NOT NULL,
    code TEXT,
    object_type TEXT NOT NULL,
    fail_on_error BOOLEAN NOT NULL,
    build_in BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    modified_by TEXT NOT NULL,
    deleted_at TIMESTAMP null,
    PRIMARY KEY (id),
    UNIQUE (name, deleted_at)
  );

create table
  "runs" (
    id TEXT,
    request_id TEXT,
    workflow_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    error TEXT,
    start_time TIMESTAMP NOT NULL,
    finish_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    modified_by TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (workflow_id) REFERENCES workflows (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
  );

CREATE TABLE
  "workflows_children" (
    workflow_id TEXT NOT NULL,
    children_id TEXT NOT NULL,
    PRIMARY KEY (children_id, workflow_id),
    FOREIGN KEY (children_id) REFERENCES workflows (id) ON DELETE CASCADE,
    FOREIGN KEY (workflow_id) REFERENCES workflows (id) ON DELETE CASCADE
  );

CREATE TABLE
  "run_logs" (
    id TEXT,
    message TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    run_id TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (run_id) REFERENCES runs (id) ON DELETE CASCADE
  );

CREATE TABLE
    "dynamic_tables" (
        id TEXT,
        table_name TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        created_by TEXT NOT NULL,
        modified_by TEXT NOT NULL,
        PRIMARY KEY (id)
);

CREATE TABLE
    "dynamic_table_variables" (
        id TEXT,
        name TEXT NOT NULL,
        variable_type TEXT NOT NULL,
        dynamic_table_id TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        created_by TEXT NOT NULL,
        modified_by TEXT NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (dynamic_table_id) REFERENCES dynamic_tables (id) ON DELETE CASCADE
);
