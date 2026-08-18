package inventory

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (repomodel.Part, error) {
	res := r.collection.FindOne(ctx, bson.M{"_id": uuid})
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return repomodel.Part{}, model.ErrPartNotFound
		}
		return repomodel.Part{}, res.Err()
	}

	var mod repomodel.Part
	err := res.Decode(&mod)
	if err != nil {
		return repomodel.Part{}, err
	}

	return mod, nil
}
