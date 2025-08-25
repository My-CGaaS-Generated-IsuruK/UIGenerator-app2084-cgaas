package dao

import (
	"UIGenerator/dbConfig"
    "UIGenerator/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateChangeLog (object *dto.ChangeLog)  error {

	result, err := dbConfig.DATABASE.Collection("ChangeLogs").UpdateOne(context.Background(), bson.M{"changelogid": object.ChangeLogId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}