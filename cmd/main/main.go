package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"template/config"
	"template/internal/app"
)

func main() {
	//получаем данные
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	// we can pass config here
	//создаем сервев
	server := app.NewServer()

	err = server.InitApp(cfg)
	if err != nil {
		panic(err)
	}

	log.Print("Server successfully init")
	// запускаем сервер
	err = server.Start()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server two closed\n")
	} else if err != nil {
		fmt.Printf("server error: %v\n", err)
		os.Exit(1)
	}
}
