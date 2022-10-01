package main

import (
	"context"
	"fmt"
	"github.com/foarsitter/youless"
	"os"
)

func QueryYouless() *youless.BasicStatusInfo {
	f := youless.OutputFormat("j") // OutputFormat |  (optional)

	configuration := youless.NewConfiguration()
	apiClient := youless.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.AGet(context.Background()).F(f).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	return resp
}
