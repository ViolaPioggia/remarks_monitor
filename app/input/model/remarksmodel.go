package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ RemarksModel = (*customRemarksModel)(nil)

type (
	// RemarksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRemarksModel.
	RemarksModel interface {
		remarksModel
	}

	customRemarksModel struct {
		*defaultRemarksModel
	}
)

// NewRemarksModel returns a model for the mongo.
func NewRemarksModel(url, db, collection string, c cache.CacheConf) RemarksModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customRemarksModel{
		defaultRemarksModel: newDefaultRemarksModel(conn),
	}
}
