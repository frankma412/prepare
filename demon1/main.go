package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd    *exec.Cmd
	output []byte
	err    error
)

func main() {
	cmd = exec.Command("bash", "-c", "sleep 2;ls -l;echo hello")

	if output, err = cmd.CombinedOutput(); nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
