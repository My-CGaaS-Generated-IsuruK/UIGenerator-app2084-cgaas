package dao

import (
    "context"
	"UIGenerator/dbConfig"
	"UIGenerator/dto"

)

func DB_CreateAttachment (object *dto.Attachment) error {

	_, err := dbConfig.DATABASE.Collection("Attachments").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}