package mongodb

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type PlayerPay struct {
	ChannelNumber string `json:"channel_number"`
	Status        int32  `json:"status"`
	Account       string `json:"Account"`
	CreateTime    string `json:"CreateTime"`
	OrderId       string `json:"OrderID"`
	PayCode       int32  `json:"PayCode"`
	PayType       string `json:"PayType"`
	PlayerId      int32  `json:"PlayerId"`
	RMB           int32  `json:"RMB"`
}

func TestMongo(t *testing.T) {
	config := Config{
		Host:     "127.0.0.1",
		Port:     27017,
		DataBase: "PaymentDB",
	}
	client, err := NewConn(config)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	var playerPay []interface{}
	w := bson.M{}
	//field := bson.D{
	//	{"channel_number", 1},
	//	{"status", 1},
	//}
	//opt := options.Find().SetProjection(field)
	err, playerPay = client.Select("player_pay", w)
	if err == nil {
		for _, v := range playerPay {
			var p PlayerPay
			row, err := json.Marshal(v.(primitive.D).Map())
			if err != nil {
				t.Fatal("error")
			}
			err = json.Unmarshal(row, &p)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(p.CreateTime, p.PayCode)
		}
	}
}
