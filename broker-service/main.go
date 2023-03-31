package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct{}

const WebPort string = "8080"

var app Config

func main() {

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	go func() {
		log.Printf("Server Listening on port %s\n", WebPort)
		err := s.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Got Signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
