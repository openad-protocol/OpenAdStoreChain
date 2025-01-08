package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInterface interface {
	Insert(table string, data []interface{}) (interface{}, error)
	InsertMany(table string, data []interface{}) ([]interface{}, error)
	Select(table string, w bson.M, opts ...*options.FindOptions) (error, []interface{})
	UpdateMany(table string, w, date interface{}, opts ...*options.UpdateOptions) error
	UpdateOne(table string, w, data interface{}, opts ...*options.UpdateOptions) error
	DeleteOne(table string, w interface{}) error
	DeleteTable(table string) error
}
