package main

import (
	"add-category-rule/handler"

	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	runtime.Start(handler.HandleRequest)
}
