---
title: Triggering renders from Go
source: Remotion Documentation
last_updated: 2024-12-22
---

# Triggering renders from Go

- [Home page](/)
- [Lambda](/docs/lambda)
- Rendering from Go

On this page

# Triggering renders from Go

EXPERIMENTAL

This feature is new. Please report any issues you encounter.

To trigger a Lambda render using Go, you can use the Remotion Lambda Go client. Note the following:

- You first need to [complete the Lambda setup](/docs/lambda/setup).
- Sending large input props (>200KB) is not supported with Go at the moment.
- Always match the version of the Go client with the version of the Lambda function you deployed. Otherwise, calls will fail due to version mismatch!

```

main.go
go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/remotion-dev/lambda_go_sdk"
)

type ValidationError struct {
	Field   string
	Message string
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"

	}
	return fe.Error() // default error
}

func main() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Specify the URL to your Webpack bundle
	serveUrl := os.Getenv("REMOTION_APP_SERVE_URL")
	// Specify the function you would like to call
	functionName := os.Getenv("REMOTION_APP_FUNCTION_NAME")
	// Specify the region you deployed to, for example "us-east-1"
	region := os.Getenv("REMOTION_APP_REGION")

	// Set parameters for render
	renderInputRequest := lambda_go_sdk.RemotionOptions{
		ServeUrl:     serveUrl,
		FunctionName: functionName,
		Region:       region,
		// The data that composition will use
		InputProps: map[string]interface{}{
			"data": "Let's play",
		},
		Composition: "main", // The composition to use

	}

	// Execute the render process
	renderResponse, renderError := lambda_go_sdk.RenderMediaOnLambda(renderInputRequest)

	// Check if there are validation errors
	if renderError != nil {

		validationOut := make([]ValidationError, len(renderError.(validator.ValidationErrors)))

		for i, fieldError := range renderError.(validator.ValidationErrors) {

			validationOut[i] = ValidationError{fieldError.Field(), msgForTag(fieldError)}
		}

		for _, apiError := range validationOut {
			fmt.Printf("%s: %s\n", apiError.Field, apiError.Message)
		}
		return

	}

	fmt.Print(renderResponse.RenderId)
	/// Get bucket information
	fmt.Printf("bucketName: %s\nRenderId: %s\n", renderResponse.RenderId, renderResponse.RenderId)

	// Render Progress request
	renderProgressInputRequest := lambda_go_sdk.RenderConfig{
		FunctionName: functionName,
		Region:       region,
		RenderId:     renderResponse.RenderId,
		BucketName:   renderResponse.BucketName,
		LogLevel:		  "info",
	}
	// Execute getting the render progress
	renderProgressResponse, renderProgressError := lambda_go_sdk.GetRenderProgress(renderProgressInputRequest)

	// Check if we have error
	if renderProgressError != nil {
		log.Fatalf("%s %s", "Invalid render progress response", renderProgressError)
	}

	// Get the overall render progress
	fmt.Printf("overallprogress: %f ", renderProgressResponse.OverallProgress)
}
```

```

main.go
go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/remotion-dev/lambda_go_sdk"
)

type ValidationError struct {
	Field   string
	Message string
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"

	}
	return fe.Error() // default error
}

func main() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Specify the URL to your Webpack bundle
	serveUrl := os.Getenv("REMOTION_APP_SERVE_URL")
	// Specify the function you would like to call
	functionName := os.Getenv("REMOTION_APP_FUNCTION_NAME")
	// Specify the region you deployed to, for example "us-east-1"
	region := os.Getenv("REMOTION_APP_REGION")

	// Set parameters for render
	renderInputRequest := lambda_go_sdk.RemotionOptions{
		ServeUrl:     serveUrl,
		FunctionName: functionName,
		Region:       region,
		// The data that composition will use
		InputProps: map[string]interface{}{
			"data": "Let's play",
		},
		Composition: "main", // The composition to use

	}

	// Execute the render process
	renderResponse, renderError := lambda_go_sdk.RenderMediaOnLambda(renderInputRequest)

	// Check if there are validation errors
	if renderError != nil {

		validationOut := make([]ValidationError, len(renderError.(validator.ValidationErrors)))

		for i, fieldError := range renderError.(validator.ValidationErrors) {

			validationOut[i] = ValidationError{fieldError.Field(), msgForTag(fieldError)}
		}

		for _, apiError := range validationOut {
			fmt.Printf("%s: %s\n", apiError.Field, apiError.Message)
		}
		return

	}

	fmt.Print(renderResponse.RenderId)
	/// Get bucket information
	fmt.Printf("bucketName: %s\nRenderId: %s\n", renderResponse.RenderId, renderResponse.RenderId)

	// Render Progress request
	renderProgressInputRequest := lambda_go_sdk.RenderConfig{
		FunctionName: functionName,
		Region:       region,
		RenderId:     renderResponse.RenderId,
		BucketName:   renderResponse.BucketName,
		LogLevel:		  "info",
	}
	// Execute getting the render progress
	renderProgressResponse, renderProgressError := lambda_go_sdk.GetRenderProgress(renderProgressInputRequest)

	// Check if we have error
	if renderProgressError != nil {
		log.Fatalf("%s %s", "Invalid render progress response", renderProgressError)
	}

	// Get the overall render progress
	fmt.Printf("overallprogress: %f ", renderProgressResponse.OverallProgress)
}
```

## Changelog [​](\#changelog "Direct link to Changelog")

`v4.0.6`: The response payload structure has changed. See the history of this page to see the previous structure.

## See also [​](\#see-also "Direct link to See also")

- [Using Lambda without IAM user](/docs/lambda/without-iam)
- [Permissions](/docs/lambda/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/go.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Rendering from PHP](/docs/lambda/php) [Next\
\
Rendering from Python](/docs/lambda/python)

- [Changelog](#changelog)
- [See also](#see-also)
