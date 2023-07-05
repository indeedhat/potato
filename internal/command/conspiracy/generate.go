package conspiracy

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/indeedhat/juniper"
	"github.com/indeedhat/potato/internal/store"
	"github.com/sashabaranov/go-openai"
)

const (
	GenerateKey   = "conspiracy:generate"
	GenerateUsage = "Generate conspiracy theories and save them to the database"
)

// Generate conspiracy theories using chat GPT
func Generate(repo store.TheoryRepository, client *openai.Client) juniper.CliCommandFunc {
	return func(args []string) error {
		if len(args) < 1 {
			return errors.New("expected 1 arg (./potato --cmd conspiracy:generate [count: int] <interval: seconds>)")
		}

		var interval int
		count, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to parse count: %w", err)
		}

		if len(args) > 1 {
			interval, _ = strconv.Atoi(args[1])
		}

		ctx := context.Background()
		for i := 0; i < count; i++ {
			resp, err := client.CreateChatCompletion(
				ctx,
				openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: "Create me a single sentence conspiracy theory with a funny twist",
						},
					},
					MaxTokens: 2000,
				},
			)

			if err != nil {
				return fmt.Errorf("failed to generate conspiracy: %w", err)
			}

			if err = repo.Create(resp.Choices[0].Message.Content); err != nil {
				return fmt.Errorf("failed to save conspiracy: %w", err)
			}

			log.Printf("Generated: %d/%d - %s\n", i+1, count, resp.Choices[0].Message.Content)
			if interval != 0 {
				time.Sleep(time.Duration(interval) * time.Second)
			}
		}

		return nil
	}
}
