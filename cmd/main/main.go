package main

import (
	"github.com/google/logger"
	"io/ioutil"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

func handleRequest(ctx context.Context) {
	logger.Infof("Hello World.")
}

func main() {
	defer logger.Init("emopti-ecm-license-ddbcache-populator", true, false, ioutil.Discard).Close()
	lambda.Start(handleRequest)
}