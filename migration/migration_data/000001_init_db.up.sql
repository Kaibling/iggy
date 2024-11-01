CREATE TABLE users (
  id text NOT NULL,
  username text NOT NULL UNIQUE,
  password text,
  active int NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by text NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by text NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE tokens (
  id text NOT NULL,
  value text NOT NULL UNIQUE,
  active int NOT NULL,
  expires TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  created_by text NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by text NOT NULL,
  user_id text NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE roles (
  id text NOT NULL,
  name text NOT NULL UNIQUE,
  active int NOT NULL,
  expires TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  created_by text NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by text NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE permissions (
  id text NOT NULL,
  collection text NOT NULL,
  permission int NOT NULL, 
  active int NOT NULL,
  expires TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  created_by text NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by text NOT NULL,
  role_id text NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE user_roles (
  role_id text NOT NULL,
  user_id text NOT NULL,
  active int NOT NULL,
  expires TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  created_by text NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by text NOT NULL,
  PRIMARY KEY (role_id, user_id),
  FOREIGN KEY (role_id) REFERENCES roles(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
