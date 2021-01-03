package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"log"
	"strings"
	"time"
)

func Crawl(url string)  {
	// 定义正则表达式规则，匹配 http 地址
	pattern := `/topics/\d+`

	// 初始化一个 Collector
	c := colly.NewCollector()
	redisAddr := ":6379"
	// 连接Redis
	conn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatalf("get redis conn error: %v", err)
	}
	defer conn.Close()

	// OnRequest 在请求新闻页之前触发，先验证请求url 是否满足格式
	c.OnRequest(func(request *colly.Request) {
		topic, ok := regexMatch(request.URL.Path, pattern)
		if ok {
			request.Visit(fmt.Sprintf("https://gocn.vip%s", topic))
		}
	})

	// 请求收到Response 响应时调用
	c.OnResponse(func(response *colly.Response) {
		// 字符串匹配
		topic := strings.Replace(response.Request.URL.Path, "/topics/", "", -1)
		isExists, err := exists(conn, topic)
		if isExists == 1 || err != nil {
			return
		}

		// 解析响应内容
		title, content, ok := parseContent(string(response.Body))
		titleAndContent := fmt.Sprintf("<h3>%s</h3>%s<hr>", title, content)
		fmt.Println("title and content: ", titleAndContent)

		// 从标题中获取日期
		date := getDate(title)
		// 与当前日期进行比较
		if curDate := time.Now().Format("2006-01-02"); curDate != date {
			return
		}

		if ok && content != "" && title != "" {
			// 推送到Github
			fmt.Println("-----------------")
		}

		save(conn, topic, date)
	})

	// 如果收到的内容是html，就在 onResponse 执行后调用
	c.OnHTML("a[title]", func(element *colly.HTMLElement) {
		// 通过a 标签的 href 属性获得新闻id
		path := element.Attr("href")
		topic, ok := regexMatch(path, pattern)
		if ok {
			// 通过新闻id 拼接url，并发起请求
			element.Request.Visit(fmt.Sprintf("https://gocn.vip%s", topic))
		}
	})

	c.Visit(url)
}

// 解析响应内容
func parseContent(body string) (string, string, bool) {
	// 预定义匹配内容：<p>GoCN 每日新闻 (2020-04-02)</p>
	pattern := `<p>GoCN(.|\n|\t)*每日新闻(.*?)</p>`
	title, _:= regexMatch(body, pattern)
	if title == "" {
		pattern = `<h[0-9]>GoCN(.|\n|\t)*每日新闻(.|\n|\t)*</h[0-9]>?`
		title, _ = regexMatch(body, pattern)
		if title == ""{
			return "", "", false
		}

		pattern = `>(.|\n|\t)*每日新闻(.|\n|\t)*<`
		title, _ = regexMatch(title, pattern)
		title = strings.Replace(title, "<", "", 1)
		title = strings.Replace(title, ">", "", 1)
	}

	pattern = `<ol>(.|\n|\t)*</ol>`
	content, _:= regexMatch(body, pattern)

	return title, content, true
}

// 通过正则表达式从标题中获取日期
func getDate(title string) string {
	pattern := `[0-9]{4}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}`
	date, _ := regexMatch(title, pattern)

	return date
}