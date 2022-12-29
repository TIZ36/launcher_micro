// 个人学习记录，大部分内容来自于 [https://go-zero.dev/cn/docs/advance](https://go-zero.dev/cn/docs/advance)

> 我们写一个基本服务的时候，一般的顺序
***[这里不讨论从需求设计，框架技术设计。假设我们的需求已经很清楚了，已经明确了库表结构]***

1. 根据库表设计，产生 a.sql ddl语句
2. 操作mysql数据库执行 a.sql ddl
3. 新建项目，进行三方件的配置（mysql，redis，cache，auth等）
4. 编写do、repo（orm), gozero 我们可以通过 【goctl + a.sql】帮我们产生基本的model，curd
5. 根据库表及业务拆分，明确 api （req，resp，err-handling）
6. 编写业务层代码（逻辑，数据库操作）
7. 如果是微服务框架，定义proto文件，完成服务期间的rpc调用


---

### **这里用一个环境配置服务为例子展示一下**

- **env.sql表结构设计, 放在 service/env/model/ 文件夹下 （简单贴一个实例）**

    ```mysql
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
    ```

- **model 编写(gozero框架我们这里只用用命令工具生成)**

```go
cd service/env/model
goctl model mysql ddl -src env.sql -dir . -c -style go_zero
```

- **api 编写**
1. 在 service/api/ 目录下创建一个 env.api 文件
2. 编写 api , 增量编写依次写即可，然后重新执行goctl命令，已存在的文件不会生成

```go
type (
	NewGameReq {
		GameId         int64  `json:"game_id"`
		GameName       string `json:"game_name"`
		GameNameAbbrev string `json:"game_name_abbrev"`
	}

	NewParkEnvReq {
		ParkEnvId      string `json:"park_env_id"`
		GameId         string `json:"game_id"`
		Name           string `json:"name"`
		UpdateStrategy int    `json:"update_strategy"`
	}
)

type (
	NewGameReply {
		Code   int64  `json:"code"`
		Result string `json:"result"`
	}

	NewParkEnvResp {
		Code   int64  `json:"code"`
		Result string `json:"result"`
	}
)

service env-api {
	@handler GameHandler
	post /api/env/game/new (NewGameReq) returns (NewGameReply)
	
	@handler EnvHandler
	post /api/env/new (NewParkEnvReq) returns (NewParkEnvResp)
}
```

1. 使用 goctl 命令工具自动生成api部分的其他部分

```go
goctl api go -api env.api -dir . -style go_zero
```

1. 在生成的 service/env/api/etc/env-api.yaml 中配置

```yaml
# mysql
Mysql:
  DataSource: root:0310@tcp(localhost:3306)/launcher_micro?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

# redisCache
CacheRedis:
	- Host: 127.0.0.1:6379
    Type: node
    Pass: '123456'

```

1. servicectx

```go
servcie/api/env/config 文件夹下配置config，主要是添加mysql，redis等

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}

	cache.CacheConf
}

service/api/env/svc 文件夹下配置servic_context, 主要是添加定义的Model对象，方便在业务逻辑层调用数据库

type ServiceContext struct {
	Config config.Config

	model.GameModel
	model.ParkEnvModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		GameModel:    model.NewGameModel(conn, c.CacheConf),
		ParkEnvModel: model.NewParkEnvModel(conn, c.CacheConf),
	}
}
```

