-- Schema:
-- CREATE TABLE "bottle-song" (
--         start_bottles INTEGER NOT NULL,
--         take_down     INTEGER NOT NULL,
--         result        TEXT
-- );
-- Task: update bottle-song table and set the result based on the
-- start_bottles and take_down.
UPDATE "bottle-song"
    SET result =(
    WITH RECURSIVE verses(n) AS(
        SELECT start_bottles
        UNION ALL
        SELECT n-1 FROM verses WHERE n>start_bottles-take_down+1
    )
    SELECT GROUP_CONCAT(
        CASE
            WHEN n = 10 THEN 'Ten green bottles' 
            WHEN n = 9 THEN 'Nine green bottles'
            WHEN n = 8 THEN 'Eight green bottles' 
            WHEN n = 7 THEN 'Seven green bottles'
            WHEN n = 6 THEN 'Six green bottles' 
            WHEN n = 5 THEN 'Five green bottles'
            WHEN n = 4 THEN 'Four green bottles' 
            WHEN n = 3 THEN 'Three green bottles'
            WHEN n = 2 THEN 'Two green bottles' 
            WHEN n = 1 THEN 'One green bottle'
            ELSE 'Zero'
            END
            ||' hanging on the wall,'||char(10)||
            CASE
            WHEN n = 10 THEN 'Ten green bottles' 
            WHEN n = 9 THEN 'Nine green bottles'
            WHEN n = 8 THEN 'Eight green bottles' 
            WHEN n = 7 THEN 'Seven green bottles'
            WHEN n = 6 THEN 'Six green bottles' 
            WHEN n = 5 THEN 'Five green bottles'
            WHEN n = 4 THEN 'Four green bottles' 
            WHEN n = 3 THEN 'Three green bottles'
            WHEN n = 2 THEN 'Two green bottles' 
            WHEN n = 1 THEN 'One green bottle'
            ELSE 'Zero'
            END
            ||' hanging on the wall,'||char(10)||
            'And if one green bottle should accidentally fall,'||char(10)||
            'There''ll be '||
            CASE
            WHEN n-1= 10 THEN 'ten green bottles'
            WHEN n-1= 9 THEN 'nine green bottles'
            WHEN n-1= 8 THEN 'eight green bottles'
            WHEN n-1= 7 THEN 'seven green bottles'
            WHEN n-1= 6 THEN 'six green bottles'
            WHEN n-1= 5 THEN 'five green bottles'
            WHEN n-1= 4 THEN 'four green bottles'
            WHEN n-1= 3 THEN 'three green bottles'
            WHEN n-1= 2 THEN 'two green bottles'
            WHEN n-1= 1 THEN 'one green bottle'
            ELSE 'no green bottles'
            END
            ||' hanging on the wall.',char(10)||char(10))
            FROM verses
        );
            
        
