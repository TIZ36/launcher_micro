create table if not exists `launcher_version`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `park_env_id`  char(32)     NOT NULL COMMENT '启动器所属某个park_env',
    `download_url` varchar(255) NOT NULL COMMENT '启动器下载地址',
    `version` varchar(255) NOT NULL COMMENT '启动器版本信息',
    `version_code` int unsigned NOT NULL DEFAULT '0' COMMENT '启动器版本码',
    `status`       tinyint      NOT NULL DEFAULT '0' COMMENT '游戏包状态，待发布/已发布/已下架',
    `publish_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间点',
    `create_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE uni_version(`park_env_id`, `version_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;