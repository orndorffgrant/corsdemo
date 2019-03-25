package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Transfer struct {
	To string
	HowMuch string
}

func main() {
	app := gin.Default()
	app.LoadHTMLFiles("index.html")
	app.OPTIONS("/transfer", func(c *gin.Context) {
		c.Header("access-control-allow-origin", "www.myrealbank.com")
		c.Header("access-control-allow-methods", "OPTIONS,POST")
		c.Header("access-control-allow-headers", "Content-Type,Authorization")
		c.Header("access-control-allow-credentials", "false")
		c.Status(200)
	})
	app.POST("/transfer", func(c *gin.Context) {
		var transfer Transfer
		token := c.GetHeader("Authorization")
		if token != "supersecretauthenticationtoken" {
			log.Println("Unauthorized!")
			c.Status(401)
			return
		}
		if c.ShouldBind(&transfer) == nil {
			log.Printf("Transferring %s to %s", transfer.HowMuch, transfer.To)
			c.Status(200)
			return
		}
	})
	app.Run("0.0.0.0:80")
}
