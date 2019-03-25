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
	app.GET("/", func(c *gin.Context) {
		c.SetCookie("auth", "supersecretauthenticationcookie", 300, "", "", false, false)
		c.HTML(200, "index.html", gin.H{})
	})
	app.POST("/transfer", func(c *gin.Context) {
		var transfer Transfer
		token, err := c.Cookie("auth")
		if err != nil || token != "supersecretauthenticationcookie" {
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
