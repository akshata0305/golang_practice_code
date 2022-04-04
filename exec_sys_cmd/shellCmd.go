package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type CommandRequest struct {
	Command string
	Args    string
}

func main() {
	cr := CommandRequest{"ls", "-la"}

	fmt.Println(cr.ExecuteCommand())

}

func (cr *CommandRequest) ExecuteCommand() error {

	var op_err error

	cmd := exec.Command(cr.Command, cr.Args)

	stdout, err := cmd.Output()

	if err != nil {
		op_err = errors.New("ExecuteCommand() failed!")
	}

	//return string(stdout)
	fmt.Println("Command Output:", string(stdout))

	file, _ := json.MarshalIndent(string(stdout), "\n", "\n")
	_ = ioutil.WriteFile("output.json", file, 0644)

	return op_err

}
