package inventory

import (
	"context"
	"errors"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (repomodel.Part, error) {
	res := r.data.FindOne(ctx, bson.M{"uuid": uuid})
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return repomodel.Part{}, model.ErrPartNotFound
		}
		return repomodel.Part{}, res.Err()
	}

	var model repomodel.Part
	err := res.Decode(&model)

	if err != nil {
		return repomodel.Part{}, err
	}

	return model, nil
}
