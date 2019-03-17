package configs

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func BqCon(projName string) (*bigquery.Client, error) {
	ctx := context.Background()
	jsonPath := "D:/go/src/go-bigquery/go-mux-ae-e3f93675f2bc.json"
	client, err := bigquery.NewClient(ctx, projName, option.WithCredentialsFile(jsonPath))
	if err != nil {
		return nil, err
	}
	return client, err
}
