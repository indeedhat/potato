package command

import (
	"github.com/indeedhat/juniper"
	"github.com/indeedhat/potato/internal/command/conspiracy"
	"github.com/indeedhat/potato/internal/command/server"
	"github.com/indeedhat/potato/internal/store"
	"github.com/sashabaranov/go-openai"
)

var CommandRegister juniper.CliCommandEntries

// GenerateRegister creates the command register used by both the cli and cron interfaces
func GenerateRegister(repo store.TheoryRepository, client *openai.Client) {
	CommandRegister = juniper.CliCommandEntries{
		// web server
		{
			Key:   server.ServerKey,
			Usage: server.ServerUsage,
			Run:   server.Serve(repo),
		},

		// Cron things
		{
			Key:   CronTriggerKey,
			Usage: CronTriggerUsage,
			Run:   TriggerCronTasks(&CommandRegister),
		},

		// Conspiracy theories
		{
			Key:   conspiracy.GenerateKey,
			Usage: conspiracy.GenerateUsage,
			Run:   conspiracy.Generate(repo, client),
		},
	}
}
