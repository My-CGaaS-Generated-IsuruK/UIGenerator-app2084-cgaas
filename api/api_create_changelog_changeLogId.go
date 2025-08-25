package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/functions"
    
  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      CreateChangeLog 
// @Description   This API performs the POST operation on ChangeLog. It allows you to create ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        data body dto.ChangeLog false "string collection" 
// @Success      200  {array}   dto.ChangeLog "Status OK"
// @Success      202  {array}   dto.ChangeLog "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateChangeLog [POST]

    func CreateChangeLogApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.ChangeLog{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
ChangeLogId, err := functions.IdGenerator("ChangeLogs", "ChangeLogId", "CHA")
	    if err != nil {
		    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	    }
        inputObj.ChangeLogId = ChangeLogId
        if err := functions.UniqueCheck(inputObj, "ChangeLogs", []string{ "ChangeLogId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err = dao.DB_CreateChangeLog(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

