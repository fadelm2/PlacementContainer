CREATE TABLE blocks (
                        id VARCHAR(255) PRIMARY KEY,
                        yard_id VARCHAR(255) REFERENCES yards(id),
                        code VARCHAR(255) NOT NULL,
                        total_slot INT NOT NULL,
                        total_row INT NOT NULL,
                        total_tier INT NOT NULL
);