package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      UpdateAttachment 
// @Description   This API performs the PUT operation on Attachment. It allows you to update Attachment records.
// @Tags          Attachment
// @Accept       json
// @Produce      json
// @Param        data body dto.Attachment false "string collection" 
// @Success      200  {array}   dto.Attachment "Status OK"
// @Success      202  {array}   dto.Attachment "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateAttachment [PUT]

    func UpdateAttachmentApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Attachment{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateAttachment(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

