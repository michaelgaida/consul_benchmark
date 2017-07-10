package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
	"math/rand"
)

func main() {
	consulServer := os.Args[1]
	benchmarkAmount, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("benchmarkAmount[%s] is not convertable to a number", os.Args[2])
		os.Exit(1)
	}

	fmt.Printf("\nAgentAddress: %s\nbenchmarkAmount: %d\n", consulServer, benchmarkAmount)

	keys := []string{"consul_1","consul_2","consul_3","consul_4","consul_5"}
	//averageElapsed := 0

	//for i := 0; i < len(keys); i++ {
		startTime := time.Now()
		realMain(keys, consulServer, benchmarkAmount)
		timeTrack(startTime, "KV GET", benchmarkAmount)
		//averageElapsed = averageElapsed + timeTrack(startTime, "KV GET", benchmarkAmount)

  //}
	//duration := time.Duration(averageElapsed / len(keys))
	//log.Printf("%d TIMES AVERAGE RANDOM SECRET GET %s", len(keys), duration)

}

func timeTrack(start time.Time, name string, benchmark_amount int) int {
	elapsed := time.Since(start)
	log.Printf("%d %s took %s", benchmark_amount, name, elapsed)
	return int(elapsed)
}

func realMain(key []string, consul_server string, benchmark_amount int) {

	var cmd []*(exec.Cmd)

	for i := 0; i <= benchmark_amount; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Int() % len(key)

		//fmt.Printf("\nKey: %s\n", key[n])

		cmd = append(cmd, exec.Command("curl", "http://"+consul_server+":8500/v1/kv"+key[n]))
		cmdOutput := &bytes.Buffer{}
		cmd[i].Stdout = cmdOutput
		// err := cmd[i].Run()
		cmd[i].Run()
		// fmt.Printf("\nFinished: %d, Err: %v", i, err)
		// fmt.Printf("\nOutput: %s", string(cmdOutput.Bytes()))
	}
}
