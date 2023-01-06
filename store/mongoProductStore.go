package store

import (
	"context"

	"github.com/brainzcode/nextjs_go_ecommerce/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db: db,
	}
}

func (s *MongoProductStore) Insert(p *types.Product) error {
	_, err := s.db.Collection(s.coll).InsertOne(context.TODO(), p)
	return err
}

func (s *MongoProductStore) GetByID(id string) (*types.Product, error) {
	s.db.Collection(s.coll)

	return nil, nil
}
