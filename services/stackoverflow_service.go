package services

import (
	"context"
	cfg "go-mux-bigquery-app-engine/configs"
	m "go-mux-bigquery-app-engine/models"
	"log"

	"google.golang.org/api/iterator"
)

func GetPostQuestionsService(projID string, topic string, ctx context.Context) []m.StackOverFlow {

	var questions []m.StackOverFlow
	client, err := cfg.BigqueryCon(projID, ctx)
	if err != nil {
		log.Println(err.Error())
	}

	query := client.Query(`SELECT title, CONCAT('https://stackoverflow.com/questions/', CAST(id as STRING)) as url, view_count
                FROM ` + "`bigquery-public-data.stackoverflow.posts_questions`" + `
                WHERE tags like '%` + topic + `%' ORDER BY view_count DESC LIMIT 10;`)

	rows, err := query.Read(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var row m.StackOverFlow
		err := rows.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf(err.Error())
		}

		questions = append(questions, m.StackOverFlow{
			Title:     row.Title,
			URL:       row.URL,
			ViewCount: row.ViewCount,
		})
	}
	return questions
}
