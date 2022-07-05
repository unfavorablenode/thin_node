package handlers

import (
	"net/http"
	"time"
	"log"
	"os"
	"context"
	"os/signal"

	"github.com/gorilla/mux"
)

func GetRouterWithRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandlerGet).Methods("GET")
    return r;
}

func GetServer(Router *mux.Router) *http.Server	    {
    return &http.Server{
	Handler:	Router,
	Addr:		"127.0.0.1:8000",
	IdleTimeout:	time.Second * 60,
	WriteTimeout:	15 * time.Second,
	ReadTimeout:	15 * time.Second,
    }
}

func StartServerWithGracefullShutdown(Server *http.Server)	{
    var wait = time.Second * 15

    go func()	{
	if err := Server.ListenAndServe(); err != nil  {
	    log.Println(err)
	}
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+c)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive out signal.
    <-c

    // Create a deadline to wait for
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()

    Server.Shutdown(ctx)

    log.Println("Shutting down")
    os.Exit(0)
}
