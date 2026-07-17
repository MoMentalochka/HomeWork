package inventory

import (
	"context"

	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (repomodel.Part, error) {
	res := r.data.FindOne(ctx, bson.M{"uuid": uuid})
	if res.Err() != nil {
		return repomodel.Part{}, res.Err()
	}
	var model repomodel.Part
	err := res.Decode(&model)
	if err != nil {
		return repomodel.Part{}, err
	}
	return model, nil
}
