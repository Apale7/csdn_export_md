package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const listApi = `https://blog.csdn.net/community/home-api/v1/get-business-list?page=%d&size=%d&businessType=lately&noMore=false&username=%s`

func getArticleList(username string) []Meta {
	resp, err := http.Get(fmt.Sprintf(listApi, 1, blogNum, username))
	if err := err; err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var respData ListRespData
	b, err := io.ReadAll(resp.Body)
	if err := err; err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &respData); err != nil {
		panic(err)
	}
	fmt.Printf("读取博客列表成功, 总共: %d 篇\n", len(respData.Data.List))
	return respData.Data.List
}
