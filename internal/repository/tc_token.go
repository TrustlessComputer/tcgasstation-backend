package repository

import (
	"context"
	"fmt"
	"strings"
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) ListTcToken() ([]*entity.TcToken, error) {
	resp := []*entity.TcToken{}
	filter := bson.M{"status": 1}

	cursor, err := r.DB.Collection(entity.TcToken{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) InsertTcToken(data *entity.TcToken) error {
	data.TcTokenID = strings.ToLower(data.TcTokenID)
	data.OutChainTokenID = strings.ToLower(data.OutChainTokenID)
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindTcTokenByTCTokenID(tcToken string) (*entity.TcToken, error) {

	fmt.Println("tcToken", tcToken)

	tcToken = strings.ToLower(tcToken)

	resp := &entity.TcToken{}

	filter := bson.D{{"tc_token_id", tcToken}}

	usr, err := r.FilterOne(entity.TcToken{}.CollectionName(), filter)

	fmt.Println("xxx111", usr)

	if err != nil {
		return nil, err
	}

	err = helpers.Transform(usr, resp)
	fmt.Println("xxx2222")
	if err != nil {
		return nil, err
	}
	fmt.Println("xxx333")
	return resp, nil
}

func (r *Repository) FindTcTokenByOutTokenID(outToken string) (*entity.TcToken, error) {

	fmt.Println("outToken", outToken)

	outToken = strings.ToLower(outToken)

	resp := &entity.TcToken{}

	filter := bson.D{{"out_chain_token_id", outToken}}

	usr, err := r.FilterOne(entity.TcToken{}.CollectionName(), filter)

	fmt.Println("xxx111", usr)

	if err != nil {
		return nil, err
	}

	err = helpers.Transform(usr, resp)
	fmt.Println("xxx2222")
	if err != nil {
		return nil, err
	}
	fmt.Println("xxx333")
	return resp, nil
}
