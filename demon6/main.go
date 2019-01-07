package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	next time.Time
}

var (
	schejob  map[string]*CronJob
	now time.Time
)

func main() {
	schejob = make(map[string]*CronJob)
	now = time.Now()

	//调度协程 一直检查是否有任务到期 需要执行
	for i := 1; i < 3; i++ {
		expr := cronexpr.MustParse("*/5 * * * * * *")
		cj := &CronJob{
			expr: expr,
			next: time.Now(),
		}

		key := fmt.Sprintf("%s%d", "job", i)
		schejob[key] = cj
	}

	go func() {
		for {
			now = time.Now()

			for ck, cv := range schejob  {
				if cv.next.Before(now) || cv.next.Equal(now) {
					go func(jn string) {
						fmt.Printf("%s 被调用了\n", jn)
					}(ck)

					cv.next = cv.expr.Next(now)
					fmt.Printf("%s下次执行时间：%s\n", ck, cv.next)
				}
			}

			select {
			case <- time.NewTimer(100 * time.Millisecond).C:
				
			}

			//等同于  time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(100 * time.Second)

}
