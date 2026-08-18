package inventory

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	databaseName   = "inventory"
	collectionName = "parts"
)

type repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(ctx context.Context, client *mongo.Client) *repository {
	repo := repository{
		client:     client,
		collection: client.Database(databaseName).Collection(collectionName),
	}

	indexName := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},  // 1 означает индекс по возрастанию, -1 был бы по убыванию
		Options: options.Index().SetUnique(false), // Индекс не уникальный, могут быть записи с одинаковым title
	}
	_, err := repo.collection.Indexes().CreateOne(ctx, indexName)
	if err != nil {
		log.Printf("Ошибка создания индекса: %v\n", err)
		return &repo
	}
	// partUUID := gofakeit.UUID()
	// now := time.Now()

	//	partDoc := bson.M{
	//	"_id":            partUUID,
	//	"name":           gofakeit.ProductName(),
	//	"description":    gofakeit.Sentence(15),
	//	"price":          gofakeit.Float64Range(1000, 500000),
	//	"stock_quantity": int64(gofakeit.Number(1, 100)),
	//	"category":       inventoryv1.Category_CATEGORY_WING.String(),
	//	"dimensions": bson.M{
	//		"length": gofakeit.Float64Range(10, 500),
	//		"width":  gofakeit.Float64Range(10, 200),
	//		"height": gofakeit.Float64Range(5, 100),
	//		"weight": gofakeit.Float64Range(1, 1000),
	//	},
	//	"manufacturer": bson.M{
	//		"name":    gofakeit.Company(),
	//		"country": gofakeit.Country(),
	//		"website": "https://" + gofakeit.DomainName(),
	//	},
	//	"tags":       []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
	//	"created_at": bson.NewDateTimeFromTime(now),
	//	"updated_at": bson.NewDateTimeFromTime(now),
	//}
	//
	//	_, err = repo.collection.InsertOne(ctx, partDoc)
	//	if err != nil {
	//	log.Printf("Ошибка вставки заметки: %v\n", err)
	//	return &repo
	//	}

	return &repo
}
