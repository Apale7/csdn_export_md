# 导出csdn博客markdown源码
## 背景
不想继续在粪坑网站写博客了，于是自己[用Vuepress搭了一个静态博客](https://blog.apale7.cn/)，配合github pages和服务器使用效果还挺不错。但之前写的博客基本全在CSDN，想导出来，人工一篇一篇复制就太累了，于是用go写了个爬虫自动把自己csdn上的所有文章的markdown源码导出并迁移到新站点上。这也应该是我CSDN上的最后一篇博客
## 思路
- 个人主页https://blog.csdn.net/Apale_8可以获取博客列表
- 编辑页可以查看markdown源码
- 导出为统一的格式(标题、写作时间、分类、标签和内容等信息)
- 转换为VuePress的格式

博客的元信息基本都有导出，可以很方便地通过二次开发直接生成hexo、vuepress静态页面
## 功能
- 导出个人csdn账号下，每篇博客的标题、内容、创建时间、分类、标签等信息
- 所有博客都输出到一个json文件中
- 无markdown源码的博客，content字段为html代码

## 编译
见build.sh

## 使用

- 编译
  - 直接下载对应的release也行
- 填写conf.yml配置文件
- 运行

*conf.yml必须与可执行文件在同一目录下*


## 博客列表
可以观察到博客列表是动态加载的，因此一定有接口
![在这里插入图片描述](https://img-blog.csdnimg.cn/c585e3d2e66c451fb1fbb9eaa0b5a391.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQXBhbGVfNw==,size_20,color_FFFFFF,t_70,g_se,x_16)
f12查看网络，先清空之前的所有请求，然后向下滚动，可以看到一个这样的接口
![在这里插入图片描述](https://img-blog.csdnimg.cn/686fe5b4aa6b49048381f2dbc76509eb.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQXBhbGVfNw==,size_20,color_FFFFFF,t_70,g_se,x_16)
![在这里插入图片描述](https://img-blog.csdnimg.cn/c45ad20afe534f029dcb8a889002223a.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQXBhbGVfNw==,size_20,color_FFFFFF,t_70,g_se,x_16)

参数很简单，page和size，所以直接以page=1，size取一个>=自己博客数量的数字即可获取到所有博客的链接

```go
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

type ListRespData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Meta struct {
	Type         string `json:"type"`
	FormatTime   string `json:"formatTime"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	HasOriginal  bool   `json:"hasOriginal"`
	DiggCount    int    `json:"diggCount"`
	CommentCount int    `json:"commentCount"`
	PostTime     int64  `json:"postTime"`
	CreateTime   int64  `json:"createTime"`
	URL          string `json:"url"`
	ArticleType  int    `json:"articleType"`
	ViewCount    int    `json:"viewCount"`
	Rtype        string `json:"rtype"`
}
type Data struct {
	List  []Meta      `json:"list"`
	Total interface{} `json:"total"`
}

```
(代码里resp的几个结构体直接拷贝网页请求用工具生成即可)
## 博客内容
博客内容要从编辑页面获取，因此需要登录

先f12从请求里面复制一下cookie备用
![在这里插入图片描述](https://img-blog.csdnimg.cn/4fd3fbf011c140d6bea883225134e0fc.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQXBhbGVfNw==,size_20,color_FFFFFF,t_70,g_se,x_16)

然后在编辑页面按f12找找接口，是可以找到getArticle这个接口的
https://bizapi.csdn.net/blog-console-api/v3/editor/getArticle?id=119971640&model_type=

但遗憾的是这个接口做了签名认证，读懂前端js的签名逻辑对我还是有点困难，遂放弃

### 柳暗花明
github搜了一圈，发现有人用另一个不需要签名的接口爬过csdn markdown，应该是一个旧版的接口，但还能用
https://blog-console-api.csdn.net/v1/editor/getArticle?id=%s

于是成功地爬到了博客内容

需要注意的是，富文本编辑器和markdown编辑器的内容，在不同的字段，需要判断下

```go
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

type ArticleRespData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ArticleData `json:"data"`
}
type ArticleData struct {
	ArticleID        string        `json:"article_id"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	Content          string        `json:"content"`
	Markdowncontent  string        `json:"markdowncontent"`
	Tags             string        `json:"tags"`
	Categories       string        `json:"categories"`
	Type             string        `json:"type"`
	Status           int           `json:"status"`
	ReadType         string        `json:"read_type"`
	Reason           string        `json:"reason"`
	ResourceURL      string        `json:"resource_url"`
	OriginalLink     string        `json:"original_link"`
	AuthorizedStatus bool          `json:"authorized_status"`
	CheckOriginal    bool          `json:"check_original"`
	EditorType       int           `json:"editor_type"`
	Plan             []interface{} `json:"plan"`
	VoteID           int           `json:"vote_id"`
	ScheduledTime    int           `json:"scheduled_time"`
	Level            string        `json:"level"`
	CoverType        string        `json:"cover_type"`
	CoverImages      []interface{} `json:"cover_images"`
}

type Article struct {
	Title       string
	CreateTime  int64
	FormatTime  string
	Categories  string
	Tags        []string
	ContentType string
	Content     string // 有markdown则存markdown，没有则存html
}

```
## 转换格式
VuePress里面格式是这样的
![在这里插入图片描述](https://img-blog.csdnimg.cn/ca389c88b05a47e2bfd3df11b389bc0a.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQXBhbGVfNw==,size_20,color_FFFFFF,t_70,g_se,x_16)

简单转一下即可

输出的文件名用了随机串，因为中文名称在编译的时候有点问题
```go
func toMarkdownForVuePress(a Article, path string) {
	if path[:len(path)-1] != "/" {
		path += "/"
	}
	content := fmt.Sprintf(vuePressFormat, a.Title, strings.ReplaceAll(a.FormatTime, ".", "-"), genTags(a.Tags...), genTags(a.Categories), a.Content)

	os.WriteFile(path+uuid.New()+".md", []byte(content), 0o644) // 写入文件, uuid作为文件名，因为中文名可能有问题
}
```
