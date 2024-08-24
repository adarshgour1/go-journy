package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	application "github.com/adarshgour1/go-journy/bank-application/internal/app"
	"github.com/adarshgour1/go-journy/bank-application/pkg/utils"
	"github.com/adarshgour1/go-journy/bank-application/routers"
)

func main() {

	// Setting up initial requirement
	log := utils.NewLogger("bank-management.log")
	db := utils.NewDbConnection(log)

	///////////////////////////////////////////////////////////
	/////                 START APPLICATION            ////////
	///////////////////////////////////////////////////////////
	log.Print("starting bank management application")

	app := application.NewApp(log, db)

	server := &http.Server{
		Addr: ":8080",
	}

	route := routers.NewRouter(log, server)
	route.Init()
	route.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithCancel(context.Background())

	cancel()
	route.Close(ctx)
	log.Println("Graceful shutdown complete.")
	app.Close()
}
