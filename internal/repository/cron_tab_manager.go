package repository

import (
	"context"
	"time"

	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) InsertCronJobManager(data *entity.CronJobManager) error {
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindCronJobManager(groupName, jobName string) (*entity.CronJobManager, error) {
	resp := &entity.CronJobManager{}
	f := bson.D{
		{Key: "group", Value: groupName},
		{Key: "job_name", Value: jobName},
	}

	usr, err := r.FilterOne(entity.CronJobManager{}.CollectionName(), f)
	if err != nil {
		return nil, err
	}

	err = helpers.Transform(usr, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repository) FindCronJobManagerByUUID(key string) (*entity.CronJobManager, error) {
	resp := &entity.CronJobManager{}
	usr, err := r.FilterOne(entity.CronJobManager{}.CollectionName(), bson.D{{"uuid", key}})
	if err != nil {
		return nil, err
	}

	err = helpers.Transform(usr, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repository) FindCronJobManagerByJobKey(jobKey string) ([]entity.CronJobManager, error) {
	resp := []entity.CronJobManager{}
	filter := bson.M{
		"job_key": jobKey,
	}

	cursor, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) UpdateCronJobManager(model *entity.CronJobManager) (*mongo.UpdateResult, error) {

	filter := bson.D{{Key: "uuid", Value: model.UUID}}
	result, err := r.UpdateOneObject(model.CollectionName(), filter, model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpdateCronJobManagerStatus(uuid string, status bool) (*mongo.UpdateResult, error) {
	f := bson.D{
		{Key: "uuid", Value: uuid},
	}

	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}
	result, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).UpdateOne(context.TODO(), f, update)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) UpdateCronJobManagerLastSatus(uuid, lastStatus string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{Key: "uuid", Value: uuid},
	}

	update := bson.M{
		"$set": bson.M{
			"last_status": lastStatus,
			"updated_at":  time.Now(),
		},
	}
	result, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).UpdateOne(context.TODO(), f, update)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) InsertCronJobManagerLogs(data *entity.CronJobManagerLogs) error {
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllCronJobManagerByJobKey() ([]entity.CronJobManager, error) {
	resp := []entity.CronJobManager{}
	filter := bson.M{}

	cursor, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) UpdateCronJobManagerStatusByJobKey(jobKey string, status bool) (*mongo.UpdateResult, error) {
	f := bson.M{"job_key": jobKey}

	update := bson.M{
		"$set": bson.M{
			"enabled":    status,
			"updated_at": time.Now(),
		},
	}
	result, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).UpdateMany(context.TODO(), f, update)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) UpdateCronJobManagerStatusByJobName(funcName string, status bool) (*mongo.UpdateResult, error) {
	f := bson.M{"function_name": funcName}

	update := bson.M{
		"$set": bson.M{
			"enabled":    status,
			"updated_at": time.Now(),
		},
	}
	result, err := r.DB.Collection(entity.CronJobManager{}.CollectionName()).UpdateOne(context.TODO(), f, update)
	if err != nil {
		return nil, err
	}

	return result, err
}
