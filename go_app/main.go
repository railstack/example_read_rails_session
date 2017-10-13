package main

import (
	"flag"

	c "./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// The app will run on port 3000 by default, you can custom it with the flag -port
	servePort := flag.String("port", "3000", "Http Server Port")
	flag.Parse()

	// Here we are instantiating the router
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	// Then we bind some route to some handler(controller action)
	r.GET("/", c.ReadHandler)
	r.GET("/user", c.UserHandler)
	// Let's start the server
	r.Run(":" + *servePort)
}
