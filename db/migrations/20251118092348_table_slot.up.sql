CREATE TABLE slot (
                      id VARCHAR(255)PRIMARY KEY,
                      block_id varchar(255)  NOT NULL,
                      row INT NOT NULL,
                      tier INT NOT NULL,
                      slot INT NOT NULL,
                      is_occupied BOOLEAN DEFAULT FALSE,
                      CONSTRAINT fk_block FOREIGN KEY(block_id) REFERENCES blocks(id)
);