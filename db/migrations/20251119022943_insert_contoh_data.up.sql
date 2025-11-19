INSERT INTO yards (id, code, name) VALUES
    ('YRD1', 'YRD1', 'Main Yard');

INSERT INTO blocks (id, yard_id, code, total_slot, total_row, total_tier)
VALUES ('3F', 'YRD1', '3F', 10, 6, 4);

INSERT INTO blocks (id, yard_id, code, total_slot, total_row, total_tier) VALUES
                                                                              ('B1', 'YRD1', 'B1', 30, 10, 5),          -- block import 20ft
                                                                              ('B2', 'YRD1', 'B2', 40, 12, 5),          -- block export 40ft
                                                                              ('RF1', 'YRD1', 'RF1', 20, 5, 4),         -- reefer block
                                                                              ('DG1', 'YRD1', 'DG1', 15, 15, 3);        -- dangerous goods block
INSERT INTO yard_plans (
    id, yard_id, block_id,
    from_slot, to_slot, from_row, to_row,
    container_size, container_height, container_type
) VALUES
    ('P1', 'YRD1', 'B1', 1, 20, 1, 10, 20, 8.6, 'DRY');


INSERT INTO yard_plans (
    id, yard_id, block_id,
    from_slot, to_slot, from_row, to_row,
    container_size, container_height, container_type
) VALUES (
             'PLAN20', 'YRD1', '3F',
             1, 3, 1, 5,
             20, 8.6, 'DRY'
         );

INSERT INTO yard_plans (
    id, yard_id, block_id,
    from_slot, to_slot, from_row, to_row,
    container_size, container_height, container_type
) VALUES (
             'PLAN40', 'YRD1', '3F',
             4, 7, 1, 5,
             40, 8.6, 'DRY'
         );
