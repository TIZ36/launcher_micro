package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ InstallerConfigModel = (*customInstallerConfigModel)(nil)

type (
	// InstallerConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInstallerConfigModel.
	InstallerConfigModel interface {
		installerConfigModel
	}

	customInstallerConfigModel struct {
		*defaultInstallerConfigModel
	}
)

// NewInstallerConfigModel returns a model for the database table.
func NewInstallerConfigModel(conn sqlx.SqlConn, c cache.CacheConf) InstallerConfigModel {
	return &customInstallerConfigModel{
		defaultInstallerConfigModel: newInstallerConfigModel(conn, c),
	}
}
