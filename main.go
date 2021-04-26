package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	for {
		print("☭ ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if err = execute(text); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		break
	}
}

func execute(input string) error {
	input = strings.TrimSuffix(input, "\r\n")

	arg := strings.Split(input, " ") //this gives an array of the input passed in

	switch arg[0] {
	case "cd":
		if len(arg) < 2 {
			return errors.New("Path required")
		}

		f := os.Chdir(arg[1])
		if f != nil {
			log.Fatal(f)
		}

	case "mkdir":
		if len(arg) < 2 {
			return errors.New("Name of directory required")
		}
		return os.Mkdir(arg[1], 0755)

	case "rm":
		if len(arg) < 2 {
			return errors.New("Enter name of directory to be removed")
		}

		err := os.Remove(arg[1])
		if err != nil {
			log.Fatal(err)
		}

	}

	cmd := exec.Command(arg[0], arg[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
