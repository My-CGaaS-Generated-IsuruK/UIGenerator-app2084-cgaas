package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      UpdateProject 
// @Description   This API performs the PUT operation on Project. It allows you to update Project records.
// @Tags          Project
// @Accept       json
// @Produce      json
// @Param        data body dto.Project false "string collection" 
// @Success      200  {array}   dto.Project "Status OK"
// @Success      202  {array}   dto.Project "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateProject [PUT]

    func UpdateProjectApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Project{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateProject(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

