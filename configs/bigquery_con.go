package configs

import (
	"context"

	"cloud.google.com/go/bigquery"
)

func BigqueryCon(projectID string, ctx context.Context) (*bigquery.Client, error) {
	client, err := bigquery.NewClient(ctx, projectID)
	return client, err
}
