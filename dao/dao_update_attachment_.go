package dao

import (
	"UIGenerator/dbConfig"
    "UIGenerator/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateAttachment (object *dto.Attachment)  error {

	result, err := dbConfig.DATABASE.Collection("Attachments").UpdateOne(context.Background(), bson.M{"attachmentid": object.AttachmentId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}