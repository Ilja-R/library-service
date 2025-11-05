package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Ilja-R/library-service/internal/bootstrap"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)


func main() {
	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	var cfg config.Config

	err := envconfig.ProcessWith(context.TODO(), &envconfig.Config{Target: &cfg, Lookuper: envconfig.OsLookuper()})
	if err != nil {
		panic(err)
	}
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	app := bootstrap.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quitSignal
		cancel()
	}()

	app.Run(ctx)

}