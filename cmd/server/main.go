package main

import (
	"buy2play/config"
	"buy2play/routes"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)

	config.LoadEnvVariables()
	log.Info("Environment variables loaded")

	config.ConnectToDb()
	log.Info("Database connected")

	config.SyncDatabase()
	log.Info("Database synchronized")
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.ForwardedByClientIP = true
	if err := r.SetTrustedProxies(nil); err != nil {
		panic("SetTrustedProxies failed: " + err.Error())
	}

	// Auth routes
	routes.AuthRoutes(r)
	// Profile routes
	routes.UserRoutes(r)
	// Cart routes
	routes.CartRoutes(r)
	// Order routes
	routes.OrderRoutes(r)
	// Product routes
	routes.ProductRoutes(r)
	// Email routes
	routes.MailRoutes(r)
	// Chat routes
	routes.ChatRoutes(r)
	
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("Failed to start Service: %s", err))
	}

	log.Info("Backend server is running!")
}
