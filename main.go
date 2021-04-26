package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
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
			//fmt.Fprintln(os.Stderr, err)
			print("\nerror from Inside the main body")
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

		err := os.RemoveAll(arg[1])
		if err != nil {
			print("inside rm body: ", err)
		}
	case "ok":
		val, _ := os.Hostname()
		fmt.Println("Host machine: " + val)

	case "read":
		if len(arg) < 2 {
			return errors.New("File required to read")
		}
		data, _ := os.ReadFile(arg[1])
		os.Stdout.Write(data)
	case "say":
		if arg[1] == "my" && arg[2] == "name" {
			name, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}
			print("Heisenberg's ", name.Username)
		}

	}

	cmd := exec.Command(arg[0], arg[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
