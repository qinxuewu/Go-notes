package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 电影列表地址
var url="http://zuikzy.cc/?m=vod-type-id-PAGE.html";
const YUMING  = "http://zuikzy.cc/"


//  爬虫电影网站资源 添加到数据库
func SaveVideo(page int){

	// 如果n为-1，则全部替换；如果 old 为空
	movieurl:=strings.Replace(url,"PAGE",strconv.Itoa(page),-1)
	log.Println(fmt.Sprintf("【开始爬取电影,页码:%d  请求地址为 %s 】",page,movieurl))
	res, err := http.Get(movieurl)
	checkErr(err,"【爬取网站地址请求异常 】")
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err,"爬取成功,解析html异常")

	doc.Find(".DianDian").Each(func(j int, tr *goquery.Selection) {
		title:=tr.Find("td").Eq(0).Text()  //标题
		detailsLink,_:=tr.Find("td").Eq(0).Find("a").Eq(0).Attr("href")  //详情页地址

		types:=tr.Find("td").Eq(1).Text()  // 电影类型

		// 爬取播放地址和详情链接
		res, err= http.Get(YUMING+detailsLink)
		checkErr(err,"【爬取电影: %s  异常】")
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc,_= goquery.NewDocumentFromReader(res.Body)

		//imgpath,_:=doc.Find(".videoPic img").Eq(0).Attr("src")
		//reamrk:=doc.Find(".movievod").Eq(1).Text();
		palyLink:=doc.Find(".contentURL li").Eq(1).Text()

		palyLink=strings.Replace(palyLink,"$","-",-1)

		log.Println(fmt.Sprintf("标题:%s  类型 %s 播放地址  %s ",title,types,palyLink))

		tracefile(fmt.Sprintf("标题:%s  类型 %s 播放地址  %s ",title,types,palyLink))


	})
	defer res.Body.Close()
}

func main()  {
	for i := 1; i < 10; i++ {
		SaveVideo(i)
	}
}

func tracefile(str_content string)  {
	// // 读写模式打开文件 如果不存在将创建一个新文件 写操作时将数据附加到文件尾部
	fd,_:=os.OpenFile("爬取全网VIP电影在线观看.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_content:=strings.Join([]string{str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}


// 错误检查
func checkErr(err error,msg string) {
	if err != nil {
		fmt.Println(msg,err)
	}
}