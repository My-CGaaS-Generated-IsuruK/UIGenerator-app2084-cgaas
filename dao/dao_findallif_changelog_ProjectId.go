package dao

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "UIGenerator/dbConfig"
    "UIGenerator/dto"
    "errors"
    "strconv"
    "time"
)

func DB_FindallifChangeLogbyProjectId (ProjectId string, page, size, searchTerm string, noPagination bool) (int64, *[]dto.ChangeLog, error) {
    var skip int64 = 0
    var limit int64 = 0
    
    if !noPagination {
        pageInt, err := strconv.Atoi(page)
        if err != nil || pageInt < 1 {
            return 0, nil, errors.New("invalid page number")
        }
        sizeInt, err := strconv.Atoi(size)
        if err != nil || sizeInt < 1 {
            return 0, nil, errors.New("invalid page size")
        }
        skip = int64((pageInt - 1) * sizeInt)
        limit = int64(sizeInt)
    }

   	
    
    
    filter := bson.M{"projectid" :ProjectId, "deleted" : false}

    if searchTerm != "" {
        searchConditions := []bson.M{
			{"changelogid": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"description": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"date": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"type": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"projectid": bson.M{"$regex": searchTerm, "$options": "i"}},
        }
        
        if len(filter) > 0 {
            filter = bson.M{"$and": []interface{}{filter, bson.M{"$or": searchConditions}}}
        } else {
            filter = bson.M{"$or": searchConditions}
        }
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    pipeline := []bson.M{
        {"$match": filter},
        {"$facet": bson.M{
            "metadata": []bson.M{{"$count": "total"}},
            "data": func() []bson.M {
                stages := []bson.M{}
                
                if !noPagination {
                    stages = append(stages, bson.M{"$skip": skip})
                    stages = append(stages, bson.M{"$limit": limit})
                }
                
                return stages
            }(),
        }},
    }

    cursor, err := dbConfig.DATABASE.Collection("ChangeLogs").Aggregate(ctx, pipeline)
    if err != nil {
        return 0, nil, err
    }
    defer cursor.Close(ctx)

    var results struct {
        Metadata []struct {
            Total int32 `bson:"total"`
        } `bson:"metadata"`
        Data []dto.ChangeLog `bson:"data"`
    }

    if cursor.Next(ctx) {
        if err := cursor.Decode(&results); err != nil {
            return 0, nil, err
        }
    }

    var objects []dto.ChangeLog
    if len(results.Data) > 0 {
        objects = results.Data
    } else {
        objects = []dto.ChangeLog{}
    }

    var count int64 = 0
    if len(results.Metadata) > 0 {
        count = int64(results.Metadata[0].Total)
    }

    return count, &objects, nil
}