package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	Link        string `json:"link"`
	Description string `json:"description"`
	ImageLink   string `json:"imageLink"`
}

func handleLambdaEvent(event event) (string, error) {
	mandaPraJesus(event)
	return fmt.Sprintf("%s - %s - %s", event.Link, event.Description, event.ImageLink), nil
}

func mandaPraJesus(event event) {
	resp, err := http.Get("http://jesus.com/")
	if err != nil {
		log.Fatalf("Error http get: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(errors.New("jesus retornou erro"))
	}
}
func main() {
	lambda.Start(handleLambdaEvent)
}
