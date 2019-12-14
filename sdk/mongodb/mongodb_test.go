package test

import (
	"context"
	"go-demo/utils/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type User struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Age      int    `bson:"age"`
}

func TestMongoDB(t *testing.T) {
	if env.IsCI() {
		return
	}
	var (
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ip:27017"))
	if err != nil {
		t.Error(err)
	}
	// 获取collection对象，如果没有则会自动创建
	collection := client.Database("test").Collection("user")

	// 插入
	user := &User{
		Name:     "pibigstar",
		Password: "123456",
		Age:      20,
	}
	resp, err := collection.InsertOne(ctx, user)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.InsertedID.(primitive.ObjectID).Hex())

	users := []interface{}{
		&User{
			Name:     "小明",
			Password: "666",
			Age:      20,
		},
		&User{
			Name:     "小花",
			Password: "555",
			Age:      18,
		},
	}
	// 批量插入
	if manyResult, err := collection.InsertMany(ctx, users); err == nil {
		for _, id := range manyResult.InsertedIDs {
			t.Log(id.(primitive.ObjectID).Hex())
		}
	}

	// 查询
	var skip int64 = 0
	var limit int64 = 10
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cur, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		t.Error(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result User
		err := cur.Decode(&result)
		if err != nil {
			t.Error(err)
		}
		// do something with result....
		t.Logf("%+v", result)
	}
	if err := cur.Err(); err != nil {
		t.Error(err)
	}

}
