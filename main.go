package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Start with getting the config. If this is wrong, nothing will work.
	//cfg := config.Load()

	//Get a context which will be cancelled on interrupt
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	//Set a waitgroup in the context which can be uses by gorotines across the application
	//ctx = app.SetWaitGroup(ctx)

	//server := mux.NewServer(cfg.Server)
	//server.Start(ctx, cfg.Server)

	<-ctx.Done()
	log.Printf("Gracefully shutting down...")
	//server.GracefullyStop(ctx)

}
