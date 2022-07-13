package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/Service"
	"shop/dto"
)

func CreateProduct(c *gin.Context) {
	var product dto.Product
	if err := c.BindJSON(&product); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	val, err := Service.ProdSer.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var response = map[string]interface{}{"id": val.ID,
			"product_name": val.Name,
			"price":        val.Price,
			"quantity":     val.Quantity,
			"message":      "Product Successfully Added",
		}
		c.JSON(http.StatusOK, response)
	}
}
