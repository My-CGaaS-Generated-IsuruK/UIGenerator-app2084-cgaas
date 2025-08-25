package dao

import (
	"UIGenerator/dbConfig"
	"UIGenerator/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindProjectbyProjectId (projectId string) (*dto.Project, error) {
	var object dto.Project

	err := dbConfig.DATABASE.Collection("Projects").FindOne(context.Background(), bson.M{"projectid": projectId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
