package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	username string
	cookie   string
	blogNum  int // 博客数量, 爬列表的时候用
)

func Init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	username = viper.GetString("username")
	cookie = viper.GetString("cookie")
	if username == "" || cookie == "" {
		panic("username or cookie is empty")
	}

	blogNum = viper.GetInt("blog_num")
	if blogNum == 0 {
		panic("blogNum is 0")
	}
	fmt.Println("读取配置成功, username:", username)
	// fmt.Println("blog_num:", blogNum)
}

func main() {
	Init()
	list := getArticleList(username)
	// getArticleContent(list[0])
	articles := make([]Article, 0, len(list))
	for i, meta := range list {
		article := getArticle(meta)
		articles = append(articles, article)
		time.Sleep(time.Millisecond * 100) // 爬太快怕被封
		fmt.Printf("第%d篇文章爬取成功! (%d/%d)\n", i+1, i+1, len(list))
	}
	b, _ := json.Marshal(articles)
	ioutil.WriteFile("articles.json", b, 0o644)
	os.Mkdir("articles", 0o755)
	for _, article := range articles {
		toMarkdown(article, "articles")
	}
}
