package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	artls := getAllArticles()

	if len(artls) != len(articles) {
		t.Fail()
	}

	for i, v := range artls {
		if v.Content != articles[i].Content || v.ID != articles[i].ID ||
			v.Title != articles[i].Title {
			t.Fail()
			break
		}
	}
}
