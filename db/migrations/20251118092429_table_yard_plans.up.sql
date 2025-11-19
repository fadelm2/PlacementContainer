CREATE TABLE yard_plans (
                            id VARCHAR(255) PRIMARY KEY,
                            yard_id VARCHAR(255) REFERENCES yards(id),
                            block_id VARCHAR(255) REFERENCES blocks(id),
                            from_slot INT,
                            to_slot INT,
                            from_row INT,
                            to_row INT,
                            container_size INT,
                            container_height DOUBLE PRECISION,
                            container_type VARCHAR(100)
);