package main

import (
	"context"
	"github.com/gin-gonic/gin"
	//"incomeCoach/repository"
	//"incomeCoach/repository/mdb"
	"incomeCoach/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()

	routes.SetUp(router)
	//// MongoDB setup
	//mongoRepo, err := mdb.NewMongoDBRepository()
	//if err != nil {
	//	// Handle error
	//	return
	//}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
		log.Println("starting server on port 8080")
	}()

	ctx, signalCancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer signalCancel()

	<-ctx.Done()

	log.Print("OS signal received - shutting down...\n")


	shutdownContext, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(shutdownContext); err != nil {
		log.Printf("Shutdown error: %v\n", err)
		os.Exit(1)
	}
	log.Print("Server gracefully stopped\n")
}
