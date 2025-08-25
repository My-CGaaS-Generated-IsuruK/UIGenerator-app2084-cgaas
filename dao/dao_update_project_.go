package dao

import (
	"UIGenerator/dbConfig"
    "UIGenerator/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateProject (object *dto.Project)  error {

	result, err := dbConfig.DATABASE.Collection("Projects").UpdateOne(context.Background(), bson.M{"projectid": object.ProjectId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}