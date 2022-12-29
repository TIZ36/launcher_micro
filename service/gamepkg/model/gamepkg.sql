create table if not exists `gamepack_version`
(
    `id` bigint    unsigned     NOT NULL AUTO_INCREMENT,
    `park_env_id`  char(32)     NOT NULL COMMENT '游戏包所属某个park_env',
    `pack_type`    tinyint      NOT NULL COMMENT '游戏包类型：强更新/兼容包',
    `download_url` varchar(255) NOT NULL COMMENT '游戏包下载地址',
    `version`      varchar(255) NOT NULL COMMENT '游戏包版本信息',
    `version_code` int unsigned NOT NULL DEFAULT '0' COMMENT '启动器后台统一的计算版本的版本码',
    `status`       tinyint      NOT NULL DEFAULT '0' COMMENT '游戏包状态，待发布/已发布/已下架',
    `publish_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间点',
    `create_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE uni_version(`park_env_id`, `version_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

