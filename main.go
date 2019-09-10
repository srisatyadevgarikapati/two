package main

import (
	"context"
	"encoding/json"
	//"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	//flag.DurationVar(&wait, "graceful-timeout", 60 * time.Second, "the duration for which the server gracefully waits")
	//flag.Parse()

	fmt.Println("Starting the web endpoint...")

	router := mux.NewRouter()
	router.HandleFunc("/",homePage)

	server:= &http.Server{
		Handler: router,
		Addr:"0.0.0.0:7778",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	fmt.Println("Shutting down the web endpoint...")
	os.Exit(0)
}

type message struct{
	Message string
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	response := message{Message:"From Two"}
	data,err := json.Marshal(response)

	if err !=nil{
		panic("Erorr at One JSON MARSHALL")
	}

	fmt.Fprint(writer,string(data))

}
