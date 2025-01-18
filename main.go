package main

import (
	"fmt"
	"os"

	logger "github.com/can-zanat/gologger"
)

const serverPort = ":96"

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	loggerInfoLevel := logger.NewWithLogLevel("info")
	defer func() {
		err := loggerInfoLevel.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	store := NewStore()
	service := NewService(store)
	handler := NewHandler(service)

	New(serverPort, handler, loggerInfoLevel).Run()

	return nil
}
