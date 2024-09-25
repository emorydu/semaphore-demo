package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	return articles
}

func getArticleByID(id int) (*article, error) {
	for _, art := range articles {
		if art.ID == id {
			return &art, nil
		}
	}

	return nil, errors.New("Article not found")
}
