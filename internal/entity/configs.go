package entity

import (
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
)

type Configs struct {
	BaseEntity `bson:",inline"`
	Key        string      `bson:"key"`
	Value      string      `bson:"value"`
	Data       interface{} `bson:"data"`
}

type FilterConfigs struct {
	BaseFilters
	Name       *string
	UploadedBy *string
}

func (u Configs) CollectionName() string {
	return "configs"
}

func (u Configs) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}
