package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	returnCode := realMain()
	timeTrack(startTime, "KV GET")
	os.Exit(returnCode)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func realMain() int {
	consulAgent := os.Args[1]
	key := os.Args[2]
	benchmarkAmount, err := strconv.Atoi(os.Args[3])

	if err != nil {
		fmt.Printf("benchmarkAmount[%s] is not convertable to a number", os.Args[3])
		return 1
	}

	fmt.Printf("\nAgentAddress: %s\nKey: %s\nbenchmarkAmount: %d\n", consulAgent, key, benchmarkAmount)
	var cmd []*(exec.Cmd)

	for i := 0; i <= benchmarkAmount; i++ {
		cmd = append(cmd, exec.Command("curl", "http://"+consulAgent+":8500/v1/kv/"+key))
		cmdOutput := &bytes.Buffer{}
		cmd[i].Stdout = cmdOutput
		// err := cmd[i].Run()
		cmd[i].Run()
		// fmt.Printf("\nFinished: %d, Err: %v", i, err)
		// fmt.Printf("\nOutput: %s", string(cmdOutput.Bytes()))
	}
	return 0
}
