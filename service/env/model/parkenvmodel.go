package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ParkEnvModel = (*customParkEnvModel)(nil)

type (
	// ParkEnvModel is an interface to be customized, add more methods here,
	// and implement the added methods in customParkEnvModel.
	ParkEnvModel interface {
		parkEnvModel
	}

	customParkEnvModel struct {
		*defaultParkEnvModel
	}
)

// NewParkEnvModel returns a model for the database table.
func NewParkEnvModel(conn sqlx.SqlConn, c cache.CacheConf) ParkEnvModel {
	return &customParkEnvModel{
		defaultParkEnvModel: newParkEnvModel(conn, c),
	}
}
