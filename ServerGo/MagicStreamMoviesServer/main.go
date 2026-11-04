package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nayarsahab/MagicStreamMovies/ServerGo/MagicStreamMoviesServer/database"
	"github.com/nayarsahab/MagicStreamMovies/ServerGo/MagicStreamMoviesServer/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	// This is the main function

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed Origin: http://localhost:5173")
	}

	// CORS (Cross-Origin Resource Sharing.)
	// It is a browser security mechanism that allows or restricts a web page
	//  from requesting resources hosted on a different domain, protocol, or port than its own.
	config := cors.Config{}
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	//config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	var client *mongo.Client = database.Connect()

	// sends a quick test signal to the database using an empty,
	// non-expiring background context to verify that the connection is active and responsive.
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}

	}()

	// /movies,  /register, /login, /logout, /genres, /refresh
	routes.SetupUnProtectedRoutes(router, client)

	// /movies/:imdb_id, /addmovie, /recommendedmovies, /updatereview/:imdb_id
	routes.SetupProtectedRoutes(router, client)

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router) Note: this method will block the calling goroutine indefinitely
	// unless an error happens
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}

// package main

// import (
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	router.GET("/hello", func(c *gin.Context){
// 		c.String(200, "Hello from MagicMovies...")
// 	})

// 	if err:=router.Run(":8080"); err != nil {
// 		fmt.Println("Failed to start Server ", err)
// 	}

// }
