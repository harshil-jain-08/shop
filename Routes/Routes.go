package Routes

import (
	"github.com/gin-gonic/gin"
	"shop/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/shop-api")
	{
		grp1.POST("owner", Controllers.CreateProduct)

	}
	return r
}
