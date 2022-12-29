create table if not exists `user`
(
    `id`     bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name`   char(32)        NOT NULL DEFAULT '' COMMENT '用户名',
    `gender` int(10)         NOT NULL DEFAULT 0 COMMENT '性别',
    `age`    int(10)         NOT NULL DEFAULT 10 COMMENT '年龄',
    `create_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
)  ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;