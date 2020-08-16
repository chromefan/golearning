package main

import (
	"fmt"
	"regexp"
)

var	pattern = map[string]string{
	"content_url":`<span><a href="(.+?)" target="_blank">.+?</a>.+?</span>`,
	"category_url":`<a href=\".[^\"]+?\">.+?</a>`,
	"article_html":`>(.+?)</h1>([\s\S]+?)</div>`,
	"article_title":">(.+?)</h1>",
	"article_content":`id="contson[\d\w]+">([\s\S]+?)$`,
	"article_tags_list":`<div class="tag">([\s\S]+?)</div>`,
	"article_tags":`>([^><]+?)</a>`,
	"article_info_html":`<p class="source">([\s\S.]+?)</a>\s</p>`,
	"article_info":`>([^<>：].+?)<`,
}
func main() {

	ages := map[string]string{}

	for i,_:= range pattern  {
		fmt.Println(i)
	}
	var info = "<p class=\"source\"><a href=\"https://www.gushiwen.org/shiwen/default.aspx?cstr=%e5%94%90%e4%bb%a3\">唐代</a><span>：</span><a href=\"/authorv_201a0677dee4.aspx\">元稹</a> </p>"
	reg := regexp.MustCompile(pattern["article_info"])
	result := reg.FindAllStringSubmatch(info,-1)
	for i, text := range result {
		if i == 0 {
			ages["dynasty"] = text[1]
		}
		if i == 1{
			ages["author"] = text[1]
		}
	}
	fmt.Println(ages)
}
