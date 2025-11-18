CREATE TABLE container (
                           id VARCHAR(255) PRIMARY KEY,
                           container_no VARCHAR(255) NOT NULL UNIQUE,
                           size INT NOT NULL,           -- 20 atau 40
                           type VARCHAR(10) NOT NULL,   -- dry, reefer, dg, dll
                           status VARCHAR(20),          -- import/export/transit
                           current_slot_id varchar(255),
                           CONSTRAINT fk_slot FOREIGN KEY(current_slot_id) REFERENCES slot(id)
);