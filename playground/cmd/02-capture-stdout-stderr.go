package main

import (
	"os/exec"
	"bytes"
	"log"
	"fmt"
)

func main(){
	cmd := exec.Command("ls","-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil{
		log.Fatalf("cmd.Run failed with %s\n",err)
	}

	outStr, errStr := string(stdout.Bytes()), stderr.Bytes()
	fmt.Printf("out:\n%s\nerr:\n%s\n",outStr,errStr)

}
