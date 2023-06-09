// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var prefixRemarksCacheKey = "cache:remarks:"

type remarksModel interface {
	Insert(ctx context.Context, data *Remarks) error
	FindOne(ctx context.Context, id string) (*Remarks, error)
	Find(ctx context.Context, id string) ([]*Remarks, error)
	FindByUsername(ctx context.Context, username string) ([]*Remarks, error)
	Update(ctx context.Context, data *Remarks) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultRemarksModel struct {
	conn *monc.Model
}

func newDefaultRemarksModel(conn *monc.Model) *defaultRemarksModel {
	return &defaultRemarksModel{conn: conn}
}

func (m *defaultRemarksModel) Insert(ctx context.Context, data *Remarks) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixRemarksCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultRemarksModel) FindOne(ctx context.Context, id string) (*Remarks, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Remarks
	key := prefixRemarksCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultRemarksModel) Find(ctx context.Context, id string) ([]*Remarks, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data []*Remarks

	err = m.conn.Find(ctx, &data, bson.M{"_id": oid})
	return data, nil
}
func (m *defaultRemarksModel) FindByUsername(ctx context.Context, username string) ([]*Remarks, error) {

	var data []*Remarks

	err := m.conn.Find(ctx, &data, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *defaultRemarksModel) Update(ctx context.Context, data *Remarks) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixRemarksCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultRemarksModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}
	key := prefixRemarksCacheKey + id
	res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return res, err
}
