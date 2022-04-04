package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-basic/uuid"
)

const vuePressFormat = `---
title: %s
date: %s
tags:
%s
categories:
%s
---

%s
`

func toMarkdown(a Article, path string) {
	if path[:len(path)-1] != "/" {
		path += "/"
	}
	os.WriteFile(path+a.Title+".md", []byte(a.Content), 0o644)
}

func genTags(tags ...string) string {
	var tagStr string
	for _, tag := range tags {
		tagStr += fmt.Sprintf(" - %s\n", tag)
	}
	return tagStr
}

func toMarkdownForVuePress(a Article, path string) {
	if path[:len(path)-1] != "/" {
		path += "/"
	}
	content := fmt.Sprintf(vuePressFormat, a.Title, strings.ReplaceAll(a.FormatTime, ".", "-"), genTags(a.Tags...), genTags(a.Categories), a.Content)

	os.WriteFile(path+uuid.New()+".md", []byte(content), 0o644) // 写入文件, uuid作为文件名，因为中文名可能有问题
}

func getArticleByTitle(articles []Article, title string) Article {
	for _, article := range articles {
		if article.Title == title {
			return article
		}
	}
	return Article{}
}
