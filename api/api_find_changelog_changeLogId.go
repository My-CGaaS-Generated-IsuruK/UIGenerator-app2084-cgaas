package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindChangeLog 
// @Description   This API performs the GET operation on ChangeLog. It allows you to retrieve ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.ChangeLog "Status OK"
// @Success      202  {array}   dto.ChangeLog "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindChangeLog [GET]

    func FindChangeLogApi(c *fiber.Ctx) error {





    
        
            
        changeLogId := c.Query("changeLogId")
                    
                
                    
            



    
  returnValue, err := dao.DB_FindChangeLogbyChangeLogId(changeLogId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

