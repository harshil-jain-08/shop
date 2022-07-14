package Routes

import (
	"github.com/gin-gonic/gin"
	"shop/Auth"
	"shop/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/shop-api")
	{
		grp1.POST("product", Auth.BasicAuthSeller, Controllers.CreateProduct)
		grp1.PATCH("product/:id", Auth.BasicAuthSeller, Controllers.UpdateProduct)
		grp1.GET("product/:id", Controllers.FindProduct)
		grp1.GET("product", Controllers.GetAllProducts)
		grp1.POST("order", Auth.BasicAuthCustomer, Controllers.CreateOrder)
		grp1.GET("order/:id", Auth.BasicAuthCustomer, Controllers.FindOrder)
	}
	return r
}
