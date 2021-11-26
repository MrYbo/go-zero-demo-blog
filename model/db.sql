CREATE TABLE `user`
(
    `id`       int          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username` varchar(255) NOT NULL UNIQUE COMMENT 'username',
    `password` varchar(255) NOT NULL COMMENT 'password',
    `avatar`   varchar(255) DEFAULT NULL COMMENT 'avatar',
    `phone`    varchar(255) DEFAULT NULL COMMENT 'phone',
    `name`     varchar(255) DEFAULT NULL COMMENT 'name',
    `address`  varchar(255) DEFAULT NULL COMMENT 'address',
    `birthday` varchar(255) DEFAULT NULL COMMENT 'birthday',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;