package Controllers

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/Mutexes"
	"shop/Service"
	"shop/dto"
)

func CreateProduct(c *gin.Context) {
	var product dto.Product
	if err := c.BindJSON(&product); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	data, err := Service.ProdSer.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var response = map[string]interface{}{"id": data.ID,
			"product_name": data.Name,
			"price":        data.Price,
			"quantity":     data.Quantity,
			"message":      "Product Successfully Added",
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateProduct(c *gin.Context) {
	prodId := c.GetInt("id")
	data, err := Service.ProdSer.FindProduct(prodId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product ID not found"})
	}
	// Mutexes should also be in service layer??
	if ok := Mutexes.ProdMutex.Lock(prodId); !ok {
		C.JSON(http.StatusPreconditionFailed, gin.H{"error": "mutex Lock found"})
		return
	}
	data, err = Service.ProdSer.UpdateProduct(data)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Information Update Failed"})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func FindProduct(c *gin.Context) {
	prodId := c.GetInt("id")
	data, err := Service.ProdSer.FindProduct(prodId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product ID not found"})
	}
	c.JSON(http.StatusOK, data)
}

func GetAllProducts(c *gin.Context) {
	data, err := Service.ProdSer.GetAllProducts()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to get data"})
	}
	c.JSON(http.StatusOK, data)
}
