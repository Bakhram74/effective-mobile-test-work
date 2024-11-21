package main

import (
	"fmt"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	"github.com/Bakhram74/effective-mobile-test-work.git/pkg/httpServer"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}
	server := httpServer.NewServer(&cfg, nil)
	if err := server.Run(); err != nil {
		panic(err)
	}

}
