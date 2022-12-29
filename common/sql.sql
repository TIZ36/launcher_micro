-- 游戏表
create table if not exists `game`
(
    `game_id`          int     unsigned NOT NULL COMMENT '游戏ID',
    `game_name`        varchar(255)     NULL     DEFAULT ''  COMMENT '游戏全称',
    `game_name_abbrev` varchar(255)     NULL     DEFAULT ''  COMMENT '游戏缩写',
    `create_time`      timestamp        NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`game_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 环境表
create table if not exists `park_env`
(
    `id`              bigint unsigned  NOT NULL AUTO_INCREMENT,
    `park_env_id`     char(32)         NOT NULL COMMENT '环境唯一标识park_env_id',
    `game_id`         int unsigned     NOT NULL COMMENT '游戏ID',
    `name`            varchar(255)     NOT NULL DEFAULT '' COMMENT '环境名称',
    `update_strategy` tinyint          NOT NULL COMMENT '环境对应的更新策略号',
    `create_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE uni_env(`park_env_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 资源表
-- 安装器资源表
create table if not exists `installer_config`
(
    `park_env_id` char(32) NOT NULL,
    `conf_url` varchar(255) NOT NULL,
    `create_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`park_env_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 启动器资源表
create table if not exists `launcher_config`
(
    `park_env_id` char(32) NOT NULL,
    `conf_url` varchar(256) NOT NULL,
    `create_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`park_env_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


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

create table if not exists `user`
(
    `id`     bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name`   char(32)        NOT NULL DEFAULT '' COMMENT '用户名',
    `gender` tinyint         NOT NULL DEFAULT 0 COMMENT '性别',
    `age`    int(10)         NOT NULL DEFAULT 10 COMMENT '年龄',
    `create_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp       NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    )  ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;