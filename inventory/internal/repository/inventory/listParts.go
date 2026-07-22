package inventory

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func (r *repository) ListParts(ctx context.Context, filters *repomodel.PartsFilter) ([]*model.Part, error) {
	// преобразуем модель фильтра в фильтр для базы
	filter := bsonFilterFromPartsFilter(filters)
	// Достаём данные из базы
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(cursor.Err(), mongo.ErrNoDocuments) {
			return nil, model.ErrPartNotFound
		}
		return nil, err
	}
	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			log.Printf("failed to close cursor: %v\n", err)
		}
	}()
	// Извлекаем данный в массив
	var parts []*model.Part
	err = cursor.All(ctx, &parts)
	if err != nil {
		return nil, err
	}
	return parts, nil
}
