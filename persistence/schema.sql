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
