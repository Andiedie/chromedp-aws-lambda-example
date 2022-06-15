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
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
		chromedp.Flag("no-zygote", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("single-process", true),

		// NetworkService,NetworkServiceInProcess from DefaultExecAllocatorOptions
		// SharedArrayBuffer from https://github.com/alixaxel/chrome-aws-lambda/blob/f9d5a9ff0282ef8e172a29d6d077efc468ca3c76/source/index.ts#L94
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),

		// site-per-process,Translate,BlinkGenPropertyTrees from DefaultExecAllocatorOptions
		// AudioServiceOutOfProcess from https://www.bannerbear.com/blog/ways-to-speed-up-puppeteer-screenshots/
		// IsolateOrigins from https://github.com/alixaxel/chrome-aws-lambda/blob/f9d5a9ff0282ef8e172a29d6d077efc468ca3c76/source/index.ts#L94
		chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),

		// DefaultExecAllocatorOptions
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),

		// https://www.bannerbear.com/blog/ways-to-speed-up-puppeteer-screenshots/
		chromedp.Flag("autoplay-policy", "user-gesture-required"),
		chromedp.Flag("disable-component-update", true),
		chromedp.Flag("disable-domain-reliability", true),
		chromedp.Flag("disable-notifications", true),
		chromedp.Flag("disable-offer-store-unmasked-wallet-cards", true),
		chromedp.Flag("disable-print-preview", true),
		chromedp.Flag("disable-speech-api", true),
		chromedp.Flag("no-pings", true),
		chromedp.Flag("use-gl", "swiftshader"),

		// https://github.com/chromedp/chromedp/issues/1074#issuecomment-1152887851
		//chromedp.Flag("use-gl", "angle"),
		//chromedp.Flag("use-angle", "swiftshader"),

		// https://github.com/alixaxel/chrome-aws-lambda/blob/f9d5a9ff0282ef8e172a29d6d077efc468ca3c76/source/index.ts#L94
		chromedp.Flag("allow-running-insecure-content", true),
		chromedp.Flag("disable-site-isolation-trials", true),
		chromedp.Flag("disable-web-security", true),

		// https://github.com/chromedp/chromedp/issues/298#issuecomment-878974456
		chromedp.Flag("homedir", "/tmp/chrome-homedir"),
		chromedp.Flag("data-path", "/tmp/chrome-data-path"),
		chromedp.Flag("disk-cache-dir", "/tmp/chrome-disk-cache-dir"),
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
