package dao

import (
    "context"
	"UIGenerator/dbConfig"
	"UIGenerator/dto"

)

func DB_CreateProject (object *dto.Project) error {

	_, err := dbConfig.DATABASE.Collection("Projects").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}