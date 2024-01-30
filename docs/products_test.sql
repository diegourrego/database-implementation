USE storage_test;

DROP TABLE IF EXISTS products;

CREATE TABLE products
(
    `id`    SERIAL PRIMARY KEY,
    `name`  VARCHAR(35),
    `type`  VARCHAR(20),
    `count` INTEGER,
    `price` FLOAT8,
    `product_code` VARCHAR(15) UNIQUE
);
