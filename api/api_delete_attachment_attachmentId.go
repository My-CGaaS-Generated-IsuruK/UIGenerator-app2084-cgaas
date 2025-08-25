package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      DeleteAttachment 
// @Description   This API performs the DELETE operation on Attachment. It allows you to delete Attachment records.
// @Tags          Attachment
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Attachment "Status OK"
// @Success      202  {array}   dto.Attachment "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteAttachment [DELETE]

    func DeleteAttachmentApi(c *fiber.Ctx) error {





    
        
            
        attachmentId := c.Query("attachmentId")
                    
                
                    
            



    
  err := dao.DB_DeleteAttachment(attachmentId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

