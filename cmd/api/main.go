package main

import (
	"fmt"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.DBSource)
}
