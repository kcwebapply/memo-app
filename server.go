package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/kcwebapply/examination/api"
)

func main() {
	r := gin.Default()
	r.Use(Logger())
	r = controller.SetRouter(r)
	r.Run(":8888")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
