//package studentsapi

package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dhi390/students-api/internal/config"
)

func main() {

	//load config

	cfg := config.MustLoad()
	//fmt.Printf("Loaded config: %+v\n", cfg)

	//database setup

	//setup router

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Welcome to Students API"))
	})

	//setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	// 	fmt.Println("server started at", cfg.HTTPServer.Addr)

	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatalf("failed to start server: %v", err)
	// 	}
	// }

	// Channel to listen for OS signals

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//Run server in a goroutine-->GRACEFULL SHUTDOWN***

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server ")
		}
	}()

	//Wait for interrupt signal
	<-done
	slog.Info("shutting down the server")

	// Context with timeout for graceful shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("server shutdown failed", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successully")

}
