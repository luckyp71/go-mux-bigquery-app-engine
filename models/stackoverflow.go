package models

type StackOverFlow struct {
	Title     string `bigquery:"title" json:"title"`
	URL       string `bigquery:"url" json:"url"`
	ViewCount int64  `bigquery:"view_count" json:"view_count"`
}
