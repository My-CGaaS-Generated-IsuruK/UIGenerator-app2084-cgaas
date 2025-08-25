package dao

import (
	"UIGenerator/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteChangeLog (changeLogId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("ChangeLogs").UpdateOne(context.Background(), bson.M{"changelogid": changeLogId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

