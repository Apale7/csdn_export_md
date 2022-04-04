package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func Test_toMarkdown(t *testing.T) {
	b, _ := ioutil.ReadFile("release/articles.json")
	articles := make([]Article, 0)
	json.Unmarshal(b, &articles)
	os.Mkdir("blogs", 0o755)
	for _, article := range articles {
		toMarkdownForVuePress(article, "blogs")
	}
}
