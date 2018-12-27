package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	outpute []byte
	err error
}

var (
	cmd *exec.Cmd
	ctx context.Context
	cancel context.CancelFunc
	resChan chan *result
	res *result
)

//contexty chan byte
//cancelFunc close chan

func main() {
	ctx, cancel = context.WithCancel(context.TODO())
	resChan = make(chan *result, 1000)

	go func() {
		resItem := new(result)
		cmd = exec.CommandContext(ctx, "bash", "-c", "sleep 2;echo hello")
		resItem.outpute, resItem.err = cmd.CombinedOutput()
		resChan <- resItem
	}()

	time.Sleep(1 * time.Second)

	cancel()

	res = <- resChan

	fmt.Println(res.err, string(res.outpute))
}
