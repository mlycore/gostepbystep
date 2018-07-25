package main

import (
	"bufio"
	"os"
	"fmt"
	strings "strings"
	"os/exec"
)

func execInput(input string) error{
	//remove the newline character
	input = strings.TrimSuffix(input,"\n")

	//Split the input to separate the command and the arguments.
	args := strings.Split(input," ")

	//pass the program and the arguments separate.
	cmd :=exec.Command(args[0],args[1:]...)


	/*//prepare the command to execute.
	cmd := exec.Command(input)*/

	//set a correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//execute command  and save it's output
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}



func main(){

	read := bufio.NewReader(os.Stdin)
	for {
		//Read input from keyboard
		input, err := read.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stdout,err)
		}
		//handle the execution of the input
		err = execInput(input)
		if err != nil{
			fmt.Println(err)
		}
	}



}


