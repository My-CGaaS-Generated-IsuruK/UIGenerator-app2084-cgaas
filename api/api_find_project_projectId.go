package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindProject 
// @Description   This API performs the GET operation on Project. It allows you to retrieve Project records.
// @Tags          Project
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Project "Status OK"
// @Success      202  {array}   dto.Project "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindProject [GET]

    func FindProjectApi(c *fiber.Ctx) error {





    
        
            
        projectId := c.Query("projectId")
                    
                
                    
            



    
  returnValue, err := dao.DB_FindProjectbyProjectId(projectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

