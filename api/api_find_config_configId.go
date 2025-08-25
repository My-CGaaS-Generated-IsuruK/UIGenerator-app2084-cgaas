package api

import (
  
"UIGenerator/utils"
"github.com/gofiber/fiber/v2"

  
    "UIGenerator/dao"
    
  

  
  
  
)

// @Summary      FindConfig 
// @Description   This API performs the GET operation on Config. It allows you to retrieve Config records.
// @Tags          Config
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Config "Status OK"
// @Success      202  {array}   dto.Config "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindConfig [GET]

    func FindConfigApi(c *fiber.Ctx) error {





    
        
            
        configId := c.Query("configId")
                    
                
                    
            



    
  returnValue, err := dao.DB_FindConfigbyConfigId(configId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

