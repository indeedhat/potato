package main

import (
	"flag"
	"log"

	"github.com/indeedhat/juniper"
	"github.com/indeedhat/potato/internal/command"
	"github.com/indeedhat/potato/internal/command/server"
	"github.com/indeedhat/potato/internal/env"
	"github.com/indeedhat/potato/internal/store"
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	if err := db.AutoMigrate(&store.Theory{}); err != nil {
		log.Fatalf("Failed to auto migrate db: %s", err)
	}

	repo := store.NewTheorySqlRepo(db)
	client := openai.NewClient(env.Get(env.OpenAiToken))

	// setup cil interface
	command.GenerateRegister(repo, client)
	flag.Usage = juniper.CliUsage(
		"Conspiracy Potato",
		"Web service providing ai generated conspiracy theories",
		"potato",
		command.CommandRegister,
	)
	flag.Parse()

	var (
		commandKey string
		args       []string
	)

	// chec for cli flags
	if len(flag.Args()) > 0 {
		commandKey = flag.Args()[0]
		args = flag.Args()[1:]
	}

	// run appropriate command
	switch commandKey {
	case "":
		commandKey = server.ServerKey
		fallthrough
	default:
		cmd := command.CommandRegister.Find(commandKey)
		if cmd == nil {
			log.Fatal("Command not found")
		}

		if err := cmd.Run(args); err != nil {
			log.Fatalf("command failed: %s", err)
		}
	}
}
