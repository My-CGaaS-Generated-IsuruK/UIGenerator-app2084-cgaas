package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  "UIGenerator/functions"
    
  "UIGenerator/dto"
    "github.com/go-playground/validator/v10"
    
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      CreateAttachment 
// @Description   This API performs the POST operation on Attachment. It allows you to create Attachment records.
// @Tags          Attachment
// @Accept       json
// @Produce      json
// @Param        data body dto.Attachment false "string collection" 
// @Success      200  {array}   dto.Attachment "Status OK"
// @Success      202  {array}   dto.Attachment "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateAttachment [POST]

    func CreateAttachmentApi(c *fiber.Ctx) error {







    
  
    inputObj := dto.Attachment{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
AttachmentId, err := functions.IdGenerator("Attachments", "AttachmentId", "ATT")
	    if err != nil {
		    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	    }
        inputObj.AttachmentId = AttachmentId
        if err := functions.UniqueCheck(inputObj, "Attachments", []string{ "AttachmentId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err = dao.DB_CreateAttachment(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

