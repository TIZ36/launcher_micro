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
