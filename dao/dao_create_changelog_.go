package dao

import (
    "context"
	"UIGenerator/dbConfig"
	"UIGenerator/dto"

)

func DB_CreateChangeLog (object *dto.ChangeLog) error {

	_, err := dbConfig.DATABASE.Collection("ChangeLogs").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}