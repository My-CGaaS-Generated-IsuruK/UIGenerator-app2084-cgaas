package dao

import (
	"UIGenerator/dbConfig"
	"UIGenerator/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindChangeLogbyChangeLogId (changeLogId string) (*dto.ChangeLog, error) {
	var object dto.ChangeLog

	err := dbConfig.DATABASE.Collection("ChangeLogs").FindOne(context.Background(), bson.M{"changelogid": changeLogId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
