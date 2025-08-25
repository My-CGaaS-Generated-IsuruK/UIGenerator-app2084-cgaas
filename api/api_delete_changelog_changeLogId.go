package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      DeleteChangeLog 
// @Description   This API performs the DELETE operation on ChangeLog. It allows you to delete ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.ChangeLog "Status OK"
// @Success      202  {array}   dto.ChangeLog "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteChangeLog [DELETE]

    func DeleteChangeLogApi(c *fiber.Ctx) error {





    
        
            
        changeLogId := c.Query("changeLogId")
                    
                
                    
            



    
  err := dao.DB_DeleteChangeLog(changeLogId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

