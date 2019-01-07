package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
		now  time.Time
		netx time.Time
	)

	if expr, err = cronexpr.Parse("* * * * *"); nil != err {
		fmt.Println(err)
		return
	}


	if expr, err = cronexpr.Parse("*/5 * * * * * *"); nil != err {
		fmt.Println(err)
		return
	}

	now = time.Now()
	netx = expr.Next(now)

	/*fmt.Println("now = ", now)
	fmt.Println("netx = ", netx)*/

	time.AfterFunc(netx.Sub(now), func() {
		fmt.Println("被调度了:", netx)
	})

	time.Sleep(6 * time.Second)
	expr = expr

}
