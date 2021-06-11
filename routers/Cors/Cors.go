package Cors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Options struct {
	Origin string
}

func CORS(options Options) gin.HandlerFunc {
	fmt.Println("loading...")
	return func(c *gin.Context) {
		fmt.Println("llll")
		fmt.Println("origin:", c.Request.Header.Get("Origin"))
		if c.Request.Header.Get("Origin") != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//	//c.JSON(http.StatusOK, "ok")
		//} else {
		//	c.Next()
		//}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			//c.JSON(http.StatusOK, "ok")
		}

		//c.Next()
	}
}
