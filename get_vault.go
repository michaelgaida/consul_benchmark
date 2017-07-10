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
	vaultServer := os.Args[1]
	benchmarkAmount, err := strconv.Atoi(os.Args[2])
	token := os.Args[3]

	if err != nil {
		fmt.Printf("benchmarkAmount[%s] is not convertable to a number", os.Args[2])
		os.Exit(1)
	}

	fmt.Printf("\nAgentAddress: %s\nbenchmarkAmount: %d\n", vaultServer, benchmarkAmount)

	keys := []string{"aaa","bbb","ccc","ddd","eee"}
	//averageElapsed := 0

	//for i := 0; i < len(keys); i++ {
		startTime := time.Now()
		realMain(keys, vaultServer, token, benchmarkAmount)
		timeTrack(startTime, "SECRET GET", benchmarkAmount)
		//averageElapsed = averageElapsed + timeTrack(startTime, "SECRET GET", benchmarkAmount)

  //}
	//duration := time.Duration(averageElapsed / len(keys))
	//log.Printf("%d TIMES AVERAGE RANDOM SECRET GET %s", len(keys), duration)

}

func timeTrack(start time.Time, name string, benchmark_amount int) int {
	elapsed := time.Since(start)
	log.Printf("%d %s took %s", benchmark_amount, name, elapsed)
	return int(elapsed)
}

func realMain(key []string, vault_server string, vault_token string, benchmark_amount int) {

	var cmd []*(exec.Cmd)

	for i := 0; i <= benchmark_amount; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Int() % len(key)

		//fmt.Printf("\nKey: %s\n", key[n])

		cmd = append(cmd, exec.Command("curl -H X-Vault-Token: "+vault_token+ "-X GET ", "http://"+vault_server+":8200/v1/secret/"+key[n]))
		cmdOutput := &bytes.Buffer{}
		cmd[i].Stdout = cmdOutput
		// err := cmd[i].Run()
		cmd[i].Run()
		// fmt.Printf("\nFinished: %d, Err: %v", i, err)
		// fmt.Printf("\nOutput: %s", string(cmdOutput.Bytes()))
	}
}
