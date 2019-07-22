package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
	models "we-blog/model"
)

func main() {
	log.Println("Starting...")

	// 根据本地时间创建一个新的 Cron job runner
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	// 向 Cron job runner 添加一个 func ，以按给定的时间表运行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	// 启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制
	c.Start()

	// 创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	t1 := time.NewTimer(time.Second * 10)
	for {
		// 阻塞 select 等待 channel
		select {
		case <-t1.C:
			// 会重置定时器，让它重新开始计时
			t1.Reset(time.Second * 10)
		}
	}
}
