package prometheus

import (
	"context"
	"github.com/foarsitter/youless"
)

func QueryYouless() (youless.UploadedValues, error) {

	configuration := youless.NewConfiguration()

	apiClient := youless.NewAPIClient(configuration)

	resp, _, err := apiClient.DefaultApi.GetUploadedValues(context.Background()).Execute()

	return resp[0], err
}
