/**
 *  基础jquery方式的  爬虫
	https://godoc.org/github.com/PuerkitoBio/goquery
   https://studygolang.com/articles/18654
 * @author qinxuewu
 * @create 19/6/16下午11:57
 * @since 1.0.0
 */
package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

const BASE="http://zuikzy.cc/?m=vod-type-id-TYPE-pg-PAGE.html"

func main()  {

	res, err := http.Get("http://www.yongjiuzy.cc/?m=vod-type-id-1-pg-1.html")
	if err != nil {

	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("【爬取成功,解析html异常：】",err)
	}


	// id选择器查询
	//doc.Find(".DianDian").Each(func(i int, selection *goquery.Selection) {
	//
	//
	//
	//})

	doc.Find(".DianDian").Each(func(j int, tr *goquery.Selection) {
		fmt.Println(tr.Html())
		title:=tr.Find("td").Eq(0).Text()
		detailsurl,_:=tr.Find("td").Eq(1).Find("a").Eq(0).Attr("href")
		types:=tr.Find("td").Eq(1).Text()
		country:=tr.Find("td").Eq(2).Text()
		state:=tr.Find("td").Eq(3).Text()
		date:=tr.Find("td").Eq(4).Text()
		fmt.Println(title,types,country,state,date,detailsurl)
		fmt.Println("-------------------------------------------")
	})

	//fmt.Println(doc.Text())
}