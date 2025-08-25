package dao

import (
    "context"
	"UIGenerator/dbConfig"
	"UIGenerator/dto"

)

func DB_CreateConfig (object *dto.Config) error {

	_, err := dbConfig.DATABASE.Collection("Configs").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}