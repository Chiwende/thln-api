package main

import (
	"github.com/gin-gonic/gin"
	"qinsights.com/thln/initializers"
	"qinsights.com/thln/routes"
)

func main() {
	initializers.LoadEnv()

	// Connect to Redis
	// redisClient, err := db.ConnectRedis(nil)
	// if err != nil {
	// 	log.Fatal("Failed to connect to Redis:", err)
	// }
	// defer db.CloseRedis(redisClient)

	// db, err := db.Connect(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	router := gin.Default()
	routes.RegisterPaymentRoutes(router)
	routes.RegisterKYCRoutes(router)
	routes.RegisterCallbackRoutes(router)
	router.Run(":9080")
}
