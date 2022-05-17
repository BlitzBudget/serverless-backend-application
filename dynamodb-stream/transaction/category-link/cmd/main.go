package main

import (
	"category-link/handler"

	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	runtime.Start(handler.HandleRequest)
}
