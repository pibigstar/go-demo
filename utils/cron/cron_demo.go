package main

import (
	"log"
	"github.com/robfig/cron"
	"time"
)

/**
定时任务
 */
func main() {
	c := cron.New()

	// 当时间 为 1s时才会执行
	c.AddFunc("1 * * * * *", func() {
		log.Println("when 1s run ....")
	})

	// 每秒执行一次
	c.AddFunc("@every 1s", func() {
		log.Println("every 1s run ....")
	})

	// 新建job任务
	myJob := new(myJob)
	c.AddJob("@every 1s",myJob)

	c.Start()

	time.Sleep(5*time.Second)

	c.Stop()
}

type myJob struct {}

func (myJob) Run ()  {
	log.Println("job run ....")
}

func init() {
	log.Println("initial......")
}
