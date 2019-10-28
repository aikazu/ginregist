package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Artikel tidak ditemukan")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}
