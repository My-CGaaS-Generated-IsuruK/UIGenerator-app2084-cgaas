package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/functions"
    
  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      CreateConfig 
// @Description   This API performs the POST operation on Config. It allows you to create Config records.
// @Tags          Config
// @Accept       json
// @Produce      json
// @Param        data body dto.Config false "string collection" 
// @Success      200  {array}   dto.Config "Status OK"
// @Success      202  {array}   dto.Config "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateConfig [POST]

    func CreateConfigApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Config{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
ConfigId, err := functions.IdGenerator("Configs", "ConfigId", "CON")
	    if err != nil {
		    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	    }
        inputObj.ConfigId = ConfigId
        if err := functions.UniqueCheck(inputObj, "Configs", []string{ "ConfigId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err = dao.DB_CreateConfig(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

