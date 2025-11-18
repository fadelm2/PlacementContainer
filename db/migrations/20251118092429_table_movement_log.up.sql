CREATE TABLE movement_log (
                              id VARCHAR(255) PRIMARY KEY,
                              container_id VARCHAR(255) NOT NULL,
                              from_slot_id INT,
                              to_slot_id INT,
                              moved_at TIMESTAMP DEFAULT NOW(),
                              CONSTRAINT fk_container FOREIGN KEY(container_id) REFERENCES container(id)
);