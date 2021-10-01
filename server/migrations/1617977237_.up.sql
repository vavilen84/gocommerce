CREATE TABLE migration
(
    id         INT UNSIGNED NOT NULL PRIMARY KEY,
    version    BIGINT UNSIGNED NOT NULL,
    filename   varchar(255) NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL
) ENGINE=InnoDB CHARSET=utf8;

CREATE TABLE product
(
    id         INT UNSIGNED NOT NULL PRIMARY KEY,
    title      varchar(255) NOT NULL,
    sku        varchar(255) NOT NULL,
    price      BIGINT UNSIGNED NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL,
    updated_at BIGINT UNSIGNED NOT NULL,
    deleted_at BIGINT UNSIGNED NOT NULL
) ENGINE=InnoDB CHARSET=utf8;