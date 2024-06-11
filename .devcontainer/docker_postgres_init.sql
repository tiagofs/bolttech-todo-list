\connect bolttech

CREATE TABLE IF NOT EXTISTS users (
  id SERIAL PRIMARY KEY,
  email UNIQUE NOT NULL
);