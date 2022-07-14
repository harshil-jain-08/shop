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

func CreateOrder(c *gin.Context) {
	var order dto.Order
	if err := c.BindJSON(&order); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	productData, err := Service.ProdSer.FindProduct(order.ProductId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid Product ID"})
	}
	//I think the below code should be in service layer
	if productData.Quantity < order.Quantity {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Insufficient Stock"})
	}
	data, err := Service.OrderSer.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Failed to create Order"})
	} else {
		productData.Quantity -= order.Quantity
		if ok := Mutexes.ProdMutex.Lock(productData.ID); !ok {
			C.JSON(http.StatusPreconditionFailed, gin.H{"error": "mutex Lock found"})
			return
		}
		defer Mutexes.ProdMutex.UnLock(productData.ID)

		_, err := Service.ProdSer.UpdateProduct(productData)

		if err != nil {
			//Should Call Delete order here
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Information Update Failed"})
		} else {
			c.JSON(http.StatusOK, data)
		}
		var response = map[string]interface{}{
			"id":         data.ID,
			"product_id": data.ProductId,
			"quantity":   data.Quantity,
			"status":     "Order Placed",
		}
		c.JSON(http.StatusOK, response)
	}
}

func FindOrder(c *gin.Context) {
	orderId := c.GetInt("id")
	data, err := Service.OrderSer.FindOrder(orderId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Order ID not found"})
	}
	c.JSON(http.StatusOK, data)
}
