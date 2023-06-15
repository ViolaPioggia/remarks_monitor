package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ RecordModel = (*customRecordModel)(nil)

type (
	// RecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordModel.
	RecordModel interface {
		FindRecordsFromBtoS(ctx context.Context, kind int64, offset int64, num int64) (*[]Record, error)
		InsertMany(ctx context.Context, data []Record) error
		recordModel
	}

	customRecordModel struct {
		*defaultRecordModel
	}
)

// NewRecordModel returns a model for the mongo.
func NewRecordModel(url, db, collection string, c cache.CacheConf) RecordModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customRecordModel{
		defaultRecordModel: newDefaultRecordModel(conn),
	}
}
func (m *defaultRecordModel) FindRecordsFromBtoS(ctx context.Context, kind int64, offset int64, num int64) (*[]Record, error) {
	var t string
	if kind == 0 {
		t = "username"
	} else if kind == 1 {
		t = "domain"
	} else if kind == 2 {
		t = "content"
	}

	var data []Record
	sort := bson.D{{"nums", -1},
		{"_id", -1}}
	findOptions := options.Find().SetSort(sort).SetSkip(offset * num).SetLimit(num)
	err := m.conn.Find(ctx, &data, bson.M{"type": t}, findOptions)
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultRecordModel) InsertMany(ctx context.Context, data []Record) error {
	var documents []any
	for _, v := range data {
		if v.ID.IsZero() {
			v.ID = primitive.NewObjectID()
			v.CreateAt = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
			v.UpdateAt = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
			documents = append(documents, v)
		}
	}
	_, err := m.conn.InsertMany(ctx, documents)
	return err
}
