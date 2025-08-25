package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindallifChangeLogByProjectId 
// @Description   This API performs the GET operation on ChangeLog. It allows you to retrieve ChangeLog records.
// @Tags          ChangeLog
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.ChangeLog "Status OK"
// @Success      202  {array}   dto.ChangeLog "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallifChangeLogByProjectId [GET]

    func FindallifChangeLogByProjectIdApi(c *fiber.Ctx) error {





    
        
            
        ProjectId := c.Query("ProjectId")
                    
                
            page := c.Query("page", "1")
                size := c.Query("size", "10")
                searchTerm := c.Query("searchTerm", "")
                noPaginationStr := c.Query("noPagination", "")
                noPagination := noPaginationStr == "true"
                    
            



    
  Count, ChangeLog ,err := dao.DB_FindallifChangeLogbyProjectId(ProjectId,page,size,searchTerm,noPagination)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }    


returnValue := map[string]interface{}{
                "Count":    Count,
                "ChangeLog": ChangeLog,
            }
    
        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

