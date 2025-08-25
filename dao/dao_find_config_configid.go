package dao

import (
	"UIGenerator/dbConfig"
	"UIGenerator/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindConfigbyConfigId (configId string) (*dto.Config, error) {
	var object dto.Config

	err := dbConfig.DATABASE.Collection("Configs").FindOne(context.Background(), bson.M{"configid": configId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
