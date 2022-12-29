package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LauncherConfigModel = (*customLauncherConfigModel)(nil)

type (
	// LauncherConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLauncherConfigModel.
	LauncherConfigModel interface {
		launcherConfigModel
	}

	customLauncherConfigModel struct {
		*defaultLauncherConfigModel
	}
)

// NewLauncherConfigModel returns a model for the database table.
func NewLauncherConfigModel(conn sqlx.SqlConn, c cache.CacheConf) LauncherConfigModel {
	return &customLauncherConfigModel{
		defaultLauncherConfigModel: newLauncherConfigModel(conn, c),
	}
}
