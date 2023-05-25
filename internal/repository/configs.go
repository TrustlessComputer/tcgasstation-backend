package repository

import (
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) FindConfig(key string) (*entity.Configs, error) {
	resp := &entity.Configs{}
	usr, err := r.FilterOne(entity.Configs{}.CollectionName(), bson.D{{"key", key}})
	if err != nil {
		return nil, err
	}

	err = helpers.Transform(usr, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (r *Repository) FindConfigCustom(key string, result interface{}) error {
	usr, err := r.FilterOne(entity.Configs{}.CollectionName(), bson.D{{"key", key}})
	if err != nil {
		return err
	}

	err = helpers.Transform(usr, result)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteConfig(uuid string) (*mongo.DeleteResult, error) {
	filter := bson.D{{"uuid", uuid}}
	result, err := r.DeleteOne(entity.Configs{}.CollectionName(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) InsertConfig(data *entity.Configs) error {
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SaveConfig(data *entity.Configs) error {

	dataExist, _ := r.FindConfig(data.Key)
	if dataExist == nil {
		return r.InsertConfig(data)
	} else {
		dataExist.Data = data.Data
		_, err := r.UpdateConfig(dataExist)
		return err
	}
}

func (r *Repository) ListConfigs(filter entity.FilterConfigs) (*entity.Pagination, error) {
	confs := []entity.Configs{}
	resp := &entity.Pagination{}
	f := bson.M{}

	p, err := r.Paginate(entity.Configs{}.CollectionName(), filter.Page, filter.Limit, f, bson.D{}, []Sort{}, &confs)
	if err != nil {
		return nil, err
	}
	resp.Result = confs
	resp.Page = p.Pagination.Page
	resp.Total = p.Pagination.Total
	resp.PageSize = filter.Limit
	return resp, nil
}

func (r *Repository) UpdateConfig(config *entity.Configs) (*mongo.UpdateResult, error) {
	filter := bson.D{{"uuid", config.UUID}}
	result, err := r.UpdateOneObject(entity.Configs{}.CollectionName(), filter, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}
