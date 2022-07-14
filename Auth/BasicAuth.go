package Auth

import (
	"github.com/gin-gonic/gin"
	"log"
)

func BasicAuthSeller(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "user" && password == "pass" {
		log.Printf("%s authenticated", user)
		//(log.Fields{"user": user}).Info("User authenticated")
	} else {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Abort()
		return
	}
}

func BasicAuthCustomer(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "cust" && password == "1234" {
		log.Printf("%s authenticated", user)
		//(log.Fields{"user": user}).Info("User authenticated")
	} else {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Abort()
		return
	}
}
