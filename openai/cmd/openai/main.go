package main

import (
	"context"
	"flag"
	"fmt"
	"openai/pkg/screenshot"
	"openai/pkg/vision"
	"os"

	"github.com/openai/openai-go/v3"
)

const screenshotImage = "/tmp/screenshot.png"

var (
	content string
	prompt  string
)

func init() {
	flag.StringVar(&content, "content", "code", "Define the type of content to respond.")

	flag.Usage = func() {
		fmt.Println("Usage: openai --content <code|quiz|rewrite>")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	// Create context
	ctx := context.Background()

	// Read OpenAI key from environment
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: OPENAI_API_KEY environment variable not set")
		os.Exit(1)
	}

	if content == "" {
		flag.Usage()
		os.Exit(1)
	} else if content == "code" {
		prompt = "Complete the code in order to solve the exercise."

		if err := sendImage(ctx, openaiKey, prompt); err != nil {
			panic(err)
		}
	} else if content == "quiz" {
		prompt = "Answer the question as concise as possible. If there is a multiple choice just give the correct answer."

		if err := sendImage(ctx, openaiKey, prompt); err != nil {
			panic(err)
		}
	} else if content == "rewrite" {
		prompt = "Rewrite the text to improve it."

		client := openai.NewClient()

		chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Say this is a test"),
			},
			Model: openai.ChatModelGPT4o,
		})
		if err != nil {
			panic(err.Error())
		}

		println(chatCompletion.Choices[0].Message.Content)

		os.Exit(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}

}

// sendImage takes a screenshot and sends it to OpenAI Vision with the given prompt.
func sendImage(ctx context.Context, openaiKey, prompt string) error {
	// Take screenshot from display 0
	if err := screenshot.Take(screenshotImage); err != nil {
		panic(err)
	}

	// Send image to OpenAI Vision
	res, err := vision.SendImage(ctx, openaiKey, screenshotImage, prompt)
	if err != nil {
		panic(err)
	}

	// Print response
	println(res.Choices[0].Message.Content)

	return nil
}
