package dao

import (
	"UIGenerator/dbConfig"
	"UIGenerator/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindAttachmentbyAttachmentId (attachmentId string) (*dto.Attachment, error) {
	var object dto.Attachment

	err := dbConfig.DATABASE.Collection("Attachments").FindOne(context.Background(), bson.M{"attachmentid": attachmentId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
