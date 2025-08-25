package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindallifAttachmentByChangeLogId 
// @Description   This API performs the GET operation on Attachment. It allows you to retrieve Attachment records.
// @Tags          Attachment
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Attachment "Status OK"
// @Success      202  {array}   dto.Attachment "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallifAttachmentByChangeLogId [GET]

    func FindallifAttachmentByChangeLogIdApi(c *fiber.Ctx) error {





    
        
            
        ChangeLogId := c.Query("ChangeLogId")
                    
                
            page := c.Query("page", "1")
                size := c.Query("size", "10")
                searchTerm := c.Query("searchTerm", "")
                noPaginationStr := c.Query("noPagination", "")
                noPagination := noPaginationStr == "true"
                    
            



    
  Count, Attachment ,err := dao.DB_FindallifAttachmentbyChangeLogId(ChangeLogId,page,size,searchTerm,noPagination)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }    


returnValue := map[string]interface{}{
                "Count":    Count,
                "Attachment": Attachment,
            }
    
        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

