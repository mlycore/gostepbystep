package main


import (
	"bufio"
	"os"
	"fmt"
	strings "strings"
	"os/exec"
	"errors"
)

//ErrNoPath is return when "cd" was called without the second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error{
	//remove the newline character
	input = strings.TrimSuffix(input,"\n")

	//Split the input to separate the command and the arguments.
	args := strings.Split(input," ")

	//check for build-in commands.
	switch args[0] {
	case "cd":
		// cd to home dir with empty path not yet supported.
		if len(args) < 2{
			return ErrNoPath
		}
		err := os.Chdir(args[1])
		if err != nil{
			return err
		}
		//stop further processing.
		return nil
	case "exit":
		os.Exit(0)
	}

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
		fmt.Print(">")
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


