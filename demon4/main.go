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

func main() {
	schedule := make(map[string]*CronJob)
	now := time.Now()

	//register job
	for i := 1; i < 3; i++ {
		exprInfo := cronexpr.MustParse("*/5 * * * * * *")
		cj := &CronJob{
			expr: exprInfo,
			next: exprInfo.Next(now),
		}

		key := fmt.Sprintf("job%d", i)
		schedule[key] = cj
	}

	go func() {
		for {
			now = time.Now()

			for jk, jv := range schedule{
				if jv.next.Before(now) || jv.next.Equal(now) {
					go func(jbn string) {
						fmt.Println("执行：", jbn)
					}(jk)

					jv.next = jv.expr.Next(now)
					fmt.Println("下次执行时间：", jv.next)
				}
			}
		}

		select {
		case <- time.NewTimer(100 * time.Millisecond).C:
		}

	}()

	time.Sleep(100 * time.Second)
}
