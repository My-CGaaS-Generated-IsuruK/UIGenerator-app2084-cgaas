package dao

import (
	"UIGenerator/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteProject (projectId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Projects").UpdateOne(context.Background(), bson.M{"projectid": projectId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

