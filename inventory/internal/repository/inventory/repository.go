package inventory

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type repository struct {
	data *mongo.Collection
}

func NewRepository(db *mongo.Database) *repository {
	collection := db.Collection("parts")

	repo := repository{
		data: collection,
	}

	indexName := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},  // 1 означает индекс по возрастанию, -1 был бы по убыванию
		Options: options.Index().SetUnique(false), // Индекс не уникальный, могут быть записи с одинаковым title
	}

	indexUuid := mongo.IndexModel{
		Keys:    bson.D{{Key: "uuid", Value: 1}}, // 1 означает индекс по возрастанию, -1 был бы по убыванию
		Options: options.Index().SetUnique(true), // Индекс не уникальный, могут быть записи с одинаковым title
	}
	ctx := context.Background()
	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{indexName, indexUuid})

	if err != nil {
		log.Printf("Ошибка создания индекса: %v\n", err)
		return &repo
	}
	//
	//note1 := repomodel.Part{Uuid: "3", Price: 12.1, Name: "Product 4"}
	//note2 := repomodel.Part{Uuid: "4", Price: 1.3, Name: "Product 3"}

	// InsertOne вставляет один документ и возвращает его ID
	//res, err := collection.InsertMany(ctx, []any{note1, note2})
	//log.Println(res)
	//if err != nil {
	//	log.Printf("Ошибка вставки заметки: %v\n", err)
	//	return &repo
	//}

	return &repo
}
