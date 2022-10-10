package main

import (
	"context"
	"fmt"
	"github.com/foarsitter/youless"
	"os"
)

func QueryYouless() (youless.UploadedValues, error) {

	configuration := youless.NewConfiguration()
	apiClient := youless.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUploadedValues(context.Background()).Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AGet``: %v\n", err)
		_, _ = fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return youless.UploadedValues{}, err
	}

	return resp[0], nil
}
