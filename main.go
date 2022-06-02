package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/chromedp"
)

func Handler(_ context.Context, _ json.RawMessage) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Headless,
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("single-process", true),
		chromedp.Flag("no-zygote", true),
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	var content string
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://example.com/"),
		chromedp.Text("body > div > p:nth-child(2)", &content),
	}); err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
	return nil
}

func main() {
	if _, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); exists {
		lambda.Start(Handler)
	} else {
		err := Handler(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
