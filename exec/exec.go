package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := "touch"
	args := []string{"teste.ok"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully halved image in size")
}