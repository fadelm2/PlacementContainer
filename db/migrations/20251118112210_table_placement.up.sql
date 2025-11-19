CREATE TABLE placements (
                            id VARCHAR(255) PRIMARY KEY,
                            yard_id VARCHAR(255) REFERENCES yards(id),
                            block_id VARCHAR(255) REFERENCES blocks(id),
                            container_number VARCHAR(255) UNIQUE,
                            slot INT,
                            row INT,
                            tier INT,
                            width INT,
                            size INT,
                            type VARCHAR(100),
                            height DOUBLE PRECISION,
                            created_at TIMESTAMP DEFAULT now(),
                            updated_at TIMESTAMP DEFAULT now()
);
