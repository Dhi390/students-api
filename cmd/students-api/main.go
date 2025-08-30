//package studentsapi --->Entry Point

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
	"github.com/Dhi390/students-api/internal/http/handlers/students"
	"github.com/Dhi390/students-api/internal/storage/sqlite"
)

func main() {

	//load config

	cfg := config.MustLoad()

	//database setup

	storage, err := sqlite.New(cfg) //yaha se postg ,other db use kr sakte h --> bss sqlite k jagah postg/other daal do nd sqlite k jaise ek package banana h postg or "implement interface" in postg way me
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	slog.Info("database connected", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", students.New(storage))
	router.HandleFunc("GET /api/students/{id}", students.GetById(storage))
	router.HandleFunc("GET /api/students", students.GetList(storage))
	router.HandleFunc("PUT /api/students/{id}", students.Update(storage))
	router.HandleFunc("DELETE /api/students/{id}", students.Delete(storage))

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

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("server shutdown failed", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successully")

}
