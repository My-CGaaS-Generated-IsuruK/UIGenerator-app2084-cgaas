package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      DeleteConfig 
// @Description   This API performs the DELETE operation on Config. It allows you to delete Config records.
// @Tags          Config
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Config "Status OK"
// @Success      202  {array}   dto.Config "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteConfig [DELETE]

    func DeleteConfigApi(c *fiber.Ctx) error {





    
        
            
        configId := c.Query("configId")
                    
                
                    
            



    
  err := dao.DB_DeleteConfig(configId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

