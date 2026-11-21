package main

import (
	"fmt"

	"github.com/PenguinKG/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	cfg.Username = "PenguinK"
	cfg.SetUser()

	cfg, err = config.Read()
	if err != nil {
		return
	}

	fmt.Printf("%v\n", cfg)
}
