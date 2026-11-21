package main

import (
	"fmt"
	"os"

	"github.com/PenguinKG/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	s := state{&cfg}
	cmds := commands{map[string]func(*state, command) error{}}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%v\n", fmt.Errorf("Error: Missing Arguments"))
		os.Exit(1)
	}

	name := args[1]
	funcArgs := []string{}
	for _, arg := range args[2:] {
		funcArgs = append(funcArgs, arg)
	}

	currCmd := command{name, funcArgs}

	err = cmds.run(&s, currCmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
