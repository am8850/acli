package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Conditions  string  `json:"conditions"`
}

var (
	port string = ""
)

func init() {
	godotenv.Load()
	sport := os.Getenv("APP_PORT")
	if sport == "" {
		port = "8080"
	}
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	weatherList := []Weather{
		{
			City:        "New York",
			Temperature: 29.5,
			Conditions:  "Cloudy",
		},
		{
			City:        "London",
			Temperature: 10.0,
			Conditions:  "Rainy",
		},
	}
	router.GET("/api/weather", func(c *gin.Context) {
		c.JSON(http.StatusOK, weatherList)
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		msg := ""

		if _, err := os.Stat("/.dockerenv"); err == nil {
			msg = ":" + port
		} else {
			msg = "http://localhost:" + port
		}

		log.Println("Starting server at: " + msg)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
