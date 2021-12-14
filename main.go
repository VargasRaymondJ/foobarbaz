package main

import (
	"context"
	modzy "github.com/modzy/sdk-go"
	"log"
	"os"
)

func main() {

	client := modzy.NewClient(os.Getenv("MODZY_URL")).WithAPIKey(os.Getenv("MODZY_API_KEY"))

	listModelsInput := (&modzy.ListModelsInput{}).WithPaging(1000, 0)
	out, err := client.Models().ListModels(context.Background(), listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("all models: %d\n", len(out.Models))
	
}
