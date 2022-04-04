package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	articleApi = `https://blog-console-api.csdn.net/v1/editor/getArticle?id=%s`
)

func getArticle(meta Meta) Article {
	articleData := getArticleContent(meta)
	content := articleData.Markdowncontent
	contentType := "markdown"
	if content == "" { // 如果没有markdown内容, 就是用富文本编辑器编写的
		content = articleData.Content
		contentType = "HTML"
	}

	return Article{
		Content:     content,
		Title:       articleData.Title,
		CreateTime:  meta.CreateTime,
		FormatTime:  meta.FormatTime,
		Categories:  articleData.Categories,
		Tags:        strings.Split(articleData.Tags, ","),
		ContentType: contentType,
	}
}

func getArticleContent(meta Meta) ArticleData {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(articleApi, meta.URL[strings.LastIndex(meta.URL, "/")+1:]), nil)
	req.Header.Set("cookie", cookie)
	resp, err := client.Do(req)
	if err := err; err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var respData ArticleRespData
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		panic(err)
	}
	return respData.Data
}
