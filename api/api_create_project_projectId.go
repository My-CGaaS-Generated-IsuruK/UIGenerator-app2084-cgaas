package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/functions"
    
  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      CreateProject 
// @Description   This API performs the POST operation on Project. It allows you to create Project records.
// @Tags          Project
// @Accept       json
// @Produce      json
// @Param        data body dto.Project false "string collection" 
// @Success      200  {array}   dto.Project "Status OK"
// @Success      202  {array}   dto.Project "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateProject [POST]

    func CreateProjectApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Project{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
ProjectId, err := functions.IdGenerator("Projects", "ProjectId", "PRO")
	    if err != nil {
		    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	    }
        inputObj.ProjectId = ProjectId
        if err := functions.UniqueCheck(inputObj, "Projects", []string{ "ProjectId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err = dao.DB_CreateProject(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

