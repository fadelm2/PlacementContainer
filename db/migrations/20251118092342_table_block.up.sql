CREATE TABLE IF NOT EXISTS blocks (
   id VARCHAR(255) PRIMARY KEY,
    yard_id VARCHAR(255) REFERENCES yards(id),
    name text,
    total_slot int NOT NULL,
    total_row int NOT NULL,
    total_tier int NOT NULL
    );