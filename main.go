package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/PenguinKG/gator/internal/config"
	"github.com/PenguinKG/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	dbURL := cfg.Url
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	s := state{
		cfg: &cfg,
		db:  database.New(db),
	}

	cmds := commands{map[string]func(*state, command) error{}}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAggregator)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%v\n", fmt.Errorf("Error: Missing Arguments"))
		os.Exit(1)
	}

	name := args[1]

	currCmd := command{name, args[2:]}

	err = cmds.run(&s, currCmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
