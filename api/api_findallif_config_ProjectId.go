package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindallifConfigByProjectId 
// @Description   This API performs the GET operation on Config. It allows you to retrieve Config records.
// @Tags          Config
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Config "Status OK"
// @Success      202  {array}   dto.Config "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallifConfigByProjectId [GET]

    func FindallifConfigByProjectIdApi(c *fiber.Ctx) error {





    
        
            
        ProjectId := c.Query("ProjectId")
                    
                
            page := c.Query("page", "1")
                size := c.Query("size", "10")
                searchTerm := c.Query("searchTerm", "")
                noPaginationStr := c.Query("noPagination", "")
                noPagination := noPaginationStr == "true"
                    
            



    
  Count, Config ,err := dao.DB_FindallifConfigbyProjectId(ProjectId,page,size,searchTerm,noPagination)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }    


returnValue := map[string]interface{}{
                "Count":    Count,
                "Config": Config,
            }
    
        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

