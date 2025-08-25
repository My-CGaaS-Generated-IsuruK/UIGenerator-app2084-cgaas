package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      UpdateChangeLog 
// @Description   This API performs the PUT operation on ChangeLog. It allows you to update ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        data body dto.ChangeLog false "string collection" 
// @Success      200  {array}   dto.ChangeLog "Status OK"
// @Success      202  {array}   dto.ChangeLog "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateChangeLog [PUT]

    func UpdateChangeLogApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.ChangeLog{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateChangeLog(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

