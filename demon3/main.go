package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		next time.Time
	)

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); nil != err {
		fmt.Println(err)
		return
	}

	now = time.Now()

	next = expr.Next(now)

	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("被调度了：", next)
	})


	time.Sleep(5 * time.Second)


}
