package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      UpdateConfig 
// @Description   This API performs the PUT operation on Config. It allows you to update Config records.
// @Tags          Config
// @Accept       json
// @Produce      json
// @Param        data body dto.Config false "string collection" 
// @Success      200  {array}   dto.Config "Status OK"
// @Success      202  {array}   dto.Config "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateConfig [PUT]

    func UpdateConfigApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Config{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateConfig(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

