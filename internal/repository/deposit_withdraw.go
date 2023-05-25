package repository

import (
	"context"
	"strings"
	"tcgasstation-backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) ListTcGasStationByStatus(statuses []entity.StatusTcGasStation) ([]*entity.TcGasStation, error) {
	resp := []*entity.TcGasStation{}
	filter := bson.M{
		"status": bson.M{"$in": statuses},
	}

	cursor, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) FindTcGasStationByStatus(status entity.StatusTcGasStation) ([]*entity.TcGasStation, error) {
	resp := []*entity.TcGasStation{}
	filter := bson.D{{"status", status}}
	cursor, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) InsertMultipleTcGasStation(listTcGasStation []*entity.TcGasStation) error {
	listDataIEntity := make([]entity.IEntity, 0, len(listTcGasStation))
	for _, TcGasStation := range listTcGasStation {
		_TcGasStation := TcGasStation
		listDataIEntity = append(listDataIEntity, _TcGasStation)
	}
	_, err := r.InsertMany(listDataIEntity)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertTcGasStation(data *entity.TcGasStation) error {
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateTcGasStation_Status_TxTcProcessDeposit_ByUuids(uuids []string, status entity.StatusTcGasStation, txTcProcessBuy string) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"uuid": bson.M{"$in": uuids},
	}

	update := bson.M{
		"$set": bson.M{
			"tx_tc_process_buy": txTcProcessBuy,
			"status":            status,
		},
	}

	result, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) UpdateTcGasStation_TxBtcProcessBuy_ByTxBuy(txTcProcessBuy, txBtcProcessBuy string, status entity.StatusTcGasStation) (*mongo.UpdateResult, error) {
	filter := bson.D{{"tx_tc_process_buy", txTcProcessBuy}}
	update := bson.M{"$set": bson.M{"status": status, "tx_btc_process_buy": txBtcProcessBuy}}
	result, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpdateTcGasStation_Status_ByTxDeposit(txTcProcessBuy string, status entity.StatusTcGasStation) (*mongo.UpdateResult, error) {
	filter := bson.D{{"tx_tc_process_buy", txTcProcessBuy}}
	update := bson.M{"$set": bson.M{"status": status}}
	result, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) FindByTcAddress(address string) ([]*entity.TcGasStation, error) {
	address = strings.ToLower(address)
	var resp []*entity.TcGasStation
	filter := bson.D{{"tc_address", address}}
	cursor, err := r.DB.Collection(entity.TcGasStation{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return resp, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// update full object
func (r *Repository) UpdateTcGasStation(TcGasStation *entity.TcGasStation) (*mongo.UpdateResult, error) {
	filter := bson.D{{"uuid", TcGasStation.UUID}}
	result, err := r.UpdateOneObject(entity.TcGasStation{}.CollectionName(), filter, TcGasStation)
	if err != nil {
		return nil, err
	}

	return result, nil
}
