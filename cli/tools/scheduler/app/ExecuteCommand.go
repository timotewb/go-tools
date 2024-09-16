package app

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) < 1 {
		fmt.Println("Invalid command format")
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	var buffer []byte
	for {
		n, err := stdout.Read(buffer)
		if n == 0 && err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(buffer[:n]))
		buffer = buffer[n:]
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}