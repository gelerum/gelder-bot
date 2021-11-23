package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID  `bson:"_id" json:"_id"`
	ChatID   int                 `bson:"chatID" json:"chatID"`
	Expenses map[int]interface{} `bson:"expenses" json:"expenses"`
	Income   []interface{}       `bson:"income" json:"income,"`
}

func (c client) Get(chatID int) {
	fmt.Print(chatID)
	filter := bson.M{"chatID": chatID}
	var a interface{}
	var sec User
	var k string
	var u map[string]interface{}
	c.Coll.FindOne(context.TODO(), filter).Decode(&a)
	c.Coll.FindOne(context.TODO(), filter).Decode(&sec)
	c.Coll.FindOne(context.TODO(), filter).Decode(&k)
	c.Coll.FindOne(context.TODO(), filter).Decode(&u)
	var d User
	var t User
	bson.Unmarshal([]byte(a.(string)), &d)
	bson.Unmarshal([]byte(k), &t)
	fmt.Println(d)
	fmt.Println(t)
	fmt.Println()
	fmt.Println()
	fmt.Println(u["Expenses"])
}

func (c client) CreateUserDocument(chatID int) error {
	count, err := c.Coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	if err != nil {
		return err
	}
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		_, err = c.Coll.InsertOne(context.TODO(), document)
	}
	return err
}

func (c client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}}}}
	_, err := c.Coll.UpdateOne(context.TODO(), filter, update)
	return err
}
