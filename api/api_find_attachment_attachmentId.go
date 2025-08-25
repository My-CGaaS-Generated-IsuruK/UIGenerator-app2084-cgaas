package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindAttachment 
// @Description   This API performs the GET operation on Attachment. It allows you to retrieve Attachment records.
// @Tags          Attachment
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Attachment "Status OK"
// @Success      202  {array}   dto.Attachment "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindAttachment [GET]

    func FindAttachmentApi(c *fiber.Ctx) error {





    
        
            
        attachmentId := c.Query("attachmentId")
                    
                
                    
            



    
  returnValue, err := dao.DB_FindAttachmentbyAttachmentId(attachmentId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

