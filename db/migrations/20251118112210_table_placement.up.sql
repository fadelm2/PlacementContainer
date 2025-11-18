CREATE TABLE placement (
                           id VARCHAR(36) PRIMARY KEY,
                           block_id VARCHAR(36) NOT NULL,
                           container_id VARCHAR(36) NOT NULL,
                           slot_start INT NOT NULL,
                           slot_end INT NOT NULL,
                           row_num INT NOT NULL,
                           tier INT NOT NULL,
                           status VARCHAR(50) NOT NULL,
                           created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           CONSTRAINT fk_container FOREIGN KEY (container_id) REFERENCES container(id),
                           CONSTRAINT fk_block FOREIGN KEY (block_id) REFERENCES blocks(id)
);
