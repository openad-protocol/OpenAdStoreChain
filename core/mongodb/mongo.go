package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Config struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DataBase string `json:"data_base"`
	Options  string `json:"options"`
	TimeOut  int64  `json:"time_out"`
}

type MongoDB struct {
	database *mongo.Database
	ctx      context.Context
}

func (mdb *MongoDB) DeleteOne(table string, w interface{}) error {
	if _, err := mdb.database.Collection(table).DeleteOne(mdb.ctx, w); err != nil {
		return err
	}
	return nil
}

func (mdb *MongoDB) DeleteTable(table string) error {
	if err := mdb.database.Collection(table).Drop(mdb.ctx); err != nil {
		return err
	}
	return nil
}

func (mdb *MongoDB) UpdateOne(table string, w, data interface{}, opts ...*options.UpdateOptions) error {
	if _, err := mdb.database.Collection(table).UpdateOne(mdb.ctx, w, data, opts...); err != nil {
		return err
	}
	return nil
}

func NewConn(config Config) (MongoInterface, error) {
	var uri string
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.TimeOut)*time.Millisecond) // ctx

	if len(config.UserName) != 0 && len(config.PassWord) != 0 && len(config.DataBase) != 0 {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			config.UserName, config.PassWord, config.Host, config.Port, config.DataBase)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		database: client.Database(config.DataBase),
		ctx:      context.Background(),
	}, nil
}

func (mdb *MongoDB) TestConn() (err error) {
	return mdb.database.Client().Ping(mdb.ctx, nil)
}

func (mdb *MongoDB) Insert(table string, data []interface{}) (interface{}, error) {
	r, err := mdb.database.Collection(table).InsertOne(mdb.ctx, data)
	if err != nil {
		return nil, err
	}
	return r.InsertedID, nil
}

func (mdb *MongoDB) InsertMany(table string, data []interface{}) ([]interface{}, error) {
	r, err := mdb.database.Collection(table).InsertMany(mdb.ctx, data)
	if err != nil {
		return nil, err
	}
	return r.InsertedIDs, err
}

func (mdb *MongoDB) Select(table string, w bson.M, opts ...*options.FindOptions) (error, []interface{}) {
	var out []interface{}
	cur, err := mdb.database.Collection(table).Find(mdb.ctx, w, opts...)
	if err != nil {
		return err, nil
	}
	defer func() {
		err = cur.Close(mdb.ctx)
		fmt.Printf("%v", err)
	}()
	for cur.Next(mdb.ctx) {
		var elem interface{}
		if err := cur.Decode(&elem); err != nil {
			return err, nil
		}
		out = append(out, elem)
	}
	return nil, out
}

func (mdb *MongoDB) UpdateMany(table string, w, data interface{}, opts ...*options.UpdateOptions) error {
	if _, err := mdb.database.Collection(table).UpdateMany(mdb.ctx, w, data, opts...); err != nil {
		return err
	}
	return nil
}
