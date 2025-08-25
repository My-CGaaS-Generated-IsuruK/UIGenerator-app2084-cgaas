package dao

import (
	"UIGenerator/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteAttachment (attachmentId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Attachments").UpdateOne(context.Background(), bson.M{"attachmentid": attachmentId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

