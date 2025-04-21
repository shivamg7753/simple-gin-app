package main

import (
	"log"
	"net/http"
	"session20-gin-app/handlers"
	"session20-gin-app/middlewares"

	"github.com/gin-gonic/gin"
)

// go build -o myapp main.go
// nohup ./myapp > myapp.log 2>&1 &
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// custom types in golang
func main() {

	r := gin.Default() // default router for gin

	r.GET("/status", PingHandler)

	// We will create crud route
	users := make([]handlers.User, 0)
	crudHandler := handlers.NewHandler(users)

	crudRoutes := r.Group("/api/v1")
	crudRoutes.Use(middlewares.Middleware())
	crudRoutes.Use(gin.Logger())
	crudRoutes.POST("/create", crudHandler.Create)
	crudRoutes.GET("/get", crudHandler.Get)
	crudRoutes.GET("/get/:id", crudHandler.GetById)
	crudRoutes.PUT("/update/:id", crudHandler.Update)

	err := r.Run("0.0.0.0:8090")
	if err != nil {
		log.Println("Error starting the gin server")
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
