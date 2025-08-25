package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  "encoding/csv"
    "fmt"
    "io"
  

  "github.com/google/uuid"
      "UIGenerator/dto"
      "encoding/json"
      "strconv"
  
  
  
)

// @Summary      UploadChangeLog 
// @Description   This API performs the UPLOAD operation on ChangeLog. It allows you to  ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        data body dto. false "string collection" 
// @Success      200  {array}   dto. "Status OK"
// @Success      202  {array}   dto. "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UploadChangeLog [UPLOAD]

    func UploadChangeLogApi(c *fiber.Ctx) error {


    file, err := c.FormFile("file")
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Could not parse form file")
	}


	src, err := file.Open()
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Could not open uploaded file")
	}
	defer src.Close()


	reader := csv.NewReader(src)


	header, _ := reader.Read()


	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error reading CSV record")
		}


		object := make(map[string]interface{})


		for i, value := range record {
			key := header[i]

			if value == "false" || value == "true" || value == "FALSE" || value == "TRUE" {
				boolValue, err := strconv.ParseBool(value)
				if err == nil {
					object[key] = boolValue
				} else {
					object[key] = value
				}
			} else {
				intValue, err := strconv.Atoi(value)
				if err == nil {
					object[key] = intValue
				} else {
					object[key] = value
				}
			}
		}


		objectJSON, err := json.Marshal(object)
		if err != nil {
			fmt.Println("Error marshalling object map:", err)
			return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
		var ChangeLogStruct dto.ChangeLog
		if err := json.Unmarshal(objectJSON, &ChangeLogStruct); err != nil {
			fmt.Println("Error unmarshalling object JSON:", err)
			return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}

		ChangeLogStruct.ChangeLogId = uuid.New().String()


		err = dao.DB_CreateChangeLog(&ChangeLogStruct)
		if err != nil {
			return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
	}






        return utils.SendSuccessResponse(c)
        
    
}

