package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// In this example `messageId` is a value that
// needs to be stored in an environment variable to be
// used by the command that's going to be executed
func doIt(messageId string) error {
	//binary, err := exec.LookPath("print_env.sh")
	//if err != nil {
	//	return err
	//}

	cmd := exec.Command("printenv")
	env := os.Environ()
	env = append(env, fmt.Sprintf("MESSAGE_ID=%s", messageId))
	cmd.Env = env

	cmdOut, _ := cmd.StdoutPipe()
	//cmdErr, _ := cmd.StderrPipe()

	startErr := cmd.Start()
	if startErr != nil {
		return startErr
	}

	// read stdout and stderr
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	//errOutput, _ := ioutil.ReadAll(cmdErr)

	fmt.Printf("STDOUT: %s\n", stdOutput)
	//fmt.Printf("ERROUT: %s\n", errOutput)

	err := cmd.Wait()
	return err
}

func main() {
	messageId := "abc123"
	err := doIt(messageId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
