package prometheus

import (
	"context"
	"github.com/foarsitter/youless"
	jww "github.com/spf13/jwalterweatherman"
)

func QueryYouless(serverURL string) (youless.UploadedValues, error) {

	//configuration := youless.NewConfiguration()
	configuration := &youless.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/v0.1.7.1/go",
		Debug:         false,
		Servers: youless.ServerConfigurations{
			{
				URL:         serverURL,
				Description: "Local device",
			},
		},
		OperationServers: map[string]youless.ServerConfigurations{},
	}

	apiClient := youless.NewAPIClient(configuration)

	resp, r, err := apiClient.DefaultApi.GetUploadedValues(context.Background()).Execute()

	jww.INFO.Printf("Http status code: %s", r.Status)

	if err != nil {
		return youless.UploadedValues{}, err
	}
	return resp[0], err
}
