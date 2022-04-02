package main

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

// -------------------

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
	Content     string // 有markdown则存markdown，没有则存html
	Title       string
	CreateTime  int64
	FormatTime  string
	Categories  string
	Tags        []string
	ContentType string
}
