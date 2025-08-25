package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      DeleteProject 
// @Description   This API performs the DELETE operation on Project. It allows you to delete Project records.
// @Tags          Project
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Project "Status OK"
// @Success      202  {array}   dto.Project "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteProject [DELETE]

    func DeleteProjectApi(c *fiber.Ctx) error {





    
        
            
        projectId := c.Query("projectId")
                    
                
                    
            



    
  err := dao.DB_DeleteProject(projectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

