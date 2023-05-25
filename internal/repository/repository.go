package repository

import (
	"context"
	"errors"
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils"
	"tcgasstation-backend/utils/global"
	"tcgasstation-backend/utils/helpers"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Connection *mongo.Client
	DB         *mongo.Database
}

type PaginationKey struct {
	Colllection string
	Page        int64
	Limit       int64
	Filter      interface{}
	OrderBy     string
	Order       entity.SortType
}

type Sort struct {
	SortBy string
	Sort   entity.SortType
}

type Count struct {
	Id    string `bson:"_id" json:"id,omitempty"`
	Count int    `bson:"count" json:"count,omitempty"`
}

func NewRepository(g *global.Global) (*Repository, error) {

	clientOption := &options.ClientOptions{}
	opt := &options.DatabaseOptions{
		ReadConcern:    clientOption.ReadConcern,
		WriteConcern:   clientOption.WriteConcern,
		ReadPreference: clientOption.ReadPreference,
		Registry:       clientOption.Registry,
	}

	r := new(Repository)
	connection := g.DBConnection.GetType()
	r.Connection = connection.(*mongo.Client)
	r.DB = r.Connection.Database(g.Conf.Databases.Mongo.Name, opt)
	return r, nil
}

func (r *Repository) InsertOne(data entity.IEntity) (*mongo.InsertOneResult, error) {
	data.SetID()
	data.SetCreatedAt()
	insertedData, err := helpers.ToDoc(data)
	if err != nil {
		return nil, err
	}

	collectionName := data.CollectionName()
	inserted, err := r.DB.Collection(collectionName).InsertOne(context.TODO(), *insertedData)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}

func (r *Repository) InsertMany(data []entity.IEntity) (*mongo.InsertManyResult, error) {
	if len(data) <= 0 {
		return nil, errors.New("Insert data is empty")
	}
	insertedData := make([]interface{}, 0)
	for _, item := range data {
		item.SetID()
		item.SetCreatedAt()
		tmp, err := helpers.ToDoc(item)
		if err != nil {
			return nil, err
		}
		insertedData = append(insertedData, *tmp)
	}

	opts := options.InsertMany().SetOrdered(false)
	inserted, err := r.DB.Collection(data[0].CollectionName()).InsertMany(context.TODO(), insertedData, opts)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateOne(collectionName string, filter bson.D, updatedData bson.M) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateOneObject(collectionName string, filter bson.D, obj entity.IEntity) (*mongo.UpdateResult, error) {
	obj.SetUpdatedAt()
	bData, err := obj.ToBson()
	if err != nil {
		return nil, err
	}

	update := bson.D{{"$set", bData}}
	result, err := r.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	//Update cache
	// id := obj.GetID()
	// r.CreateCache(dbName, id, obj)

	return result, nil
}

func (r *Repository) UpdateMany(collectionName string, filter bson.D, updatedData bson.M) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateMany(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) ReplaceOne(filter bson.D, data entity.IEntity) (*mongo.UpdateResult, error) {
	bsonData, err := helpers.ToDoc(data)
	if err != nil {
		return nil, err
	}

	inserted, err := r.DB.Collection(data.CollectionName()).ReplaceOne(context.TODO(), filter, bsonData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) DeleteOne(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	deleted, err := r.DB.Collection(collectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func (r *Repository) DeleteMany(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	deleted, err := r.DB.Collection(collectionName).DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func (r *Repository) CountDocuments(collectionName string, filter bson.D) (*int64, *int64, error) {
	estCount, estCountErr := r.DB.Collection(collectionName).EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		return nil, nil, estCountErr
	}
	count, err := r.DB.Collection(collectionName).CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, nil, err
	}

	return &count, &estCount, nil
}

func (r *Repository) FindOne(collectionName string, filter bson.D) (*mongo.SingleResult, error) {

	sr := r.DB.Collection(collectionName).FindOne(context.TODO(), filter)
	if sr.Err() != nil {
		return nil, sr.Err()
	}

	return sr, nil
}
func (r *Repository) FilterOne(collectionName string, filter bson.D, opts ...*options.FindOneOptions) (*bson.M, error) {
	data := &bson.M{}

	err := r.DB.Collection(collectionName).FindOne(context.Background(), filter, opts...).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) Find(collectionName string, filter bson.D, limit int64, offset int64, result interface{}, sort bson.D) error {
	opts := &options.FindOptions{}
	opts.Limit = &limit
	opts.Skip = &offset
	opts.Sort = sort

	cursor, err := r.DB.Collection(collectionName).Find(context.TODO(), filter, opts)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Paginate(dbName string, page int64, limit int64, filter interface{}, selectFields interface{}, sorts []Sort, returnData interface{}) (*PaginatedData, error) {
	paginatedData := New(r.DB.Collection(dbName)).
		Context(context.TODO()).
		Limit(int64(limit)).
		Page(int64(page))

	if len(sorts) > 0 {
		for _, sort := range sorts {
			if sort.Sort == entity.SORT_ASC || sort.Sort == entity.SORT_DESC {
				//sortValue := bson.D{{"created_at", -1}}
				paginatedData.Sort(sort.SortBy, sort.Sort)
			}
		}
	} else {
		paginatedData.Sort("created_at", entity.SORT_DESC)
		paginatedData.Sort(utils.KEY_UUID, entity.SORT_ASC)
	}

	data, err := paginatedData.
		Select(selectFields).
		Filter(filter).
		Decode(returnData).
		Find()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) Aggregate(dbName string, page int64, limit int64, filter interface{}, selectFields interface{}, sorts []Sort, returnData interface{}) (*PaginatedData, error) {
	paginatedData := New(r.DB.Collection(dbName)).
		Context(context.TODO()).
		Limit(int64(limit)).
		Page(int64(page))

	if len(sorts) > 0 {
		for _, sort := range sorts {
			if sort.Sort == entity.SORT_ASC || sort.Sort == entity.SORT_DESC {
				//sortValue := bson.D{{"created_at", -1}}
				paginatedData.Sort(sort.SortBy, sort.Sort)
			}
		}
	} else {
		paginatedData.Sort("created_at", entity.SORT_DESC)
		paginatedData.Sort(utils.KEY_UUID, entity.SORT_ASC)
	}

	data, err := paginatedData.
		Select(selectFields).
		Decode(returnData).
		Aggregate(filter)

	if err != nil {
		return nil, err
	}

	return data, nil
}
