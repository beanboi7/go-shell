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

	"github.com/beanboi7/go-shell/history"
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
	stringArg := strings.Join(arg, " ")

	switch arg[0] {

	//changes directory
	case "cd":
		if len(arg) < 2 {
			return errors.New("Path required")
		}

		f := os.Chdir(arg[1])
		if f != nil {
			log.Fatal(f)
		}

	//makes directory
	case "mkdir":
		if len(arg) < 2 {
			return errors.New("Name of directory required")
		}
		return os.Mkdir(arg[1], 0755)

	//removes directory
	case "rm":
		if len(arg) < 2 {
			return errors.New("Enter name of directory to be removed")
		}

		err := os.RemoveAll(arg[1])
		if err != nil {
			print("inside rm body: ", err)
		}

	//gets the host machine's name
	case "ok":
		val, _ := os.Hostname()
		fmt.Println("Host machine: " + val)

	//reads the contents of a file
	case "read":
		if len(arg) < 2 {
			return errors.New("File required to read")
		}
		data, _ := os.ReadFile(arg[1])
		os.Stdout.Write(data)

	//similar to "ok" command
	case "say":
		if arg[1] == "my" && arg[2] == "name" {
			name, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}
			print("Heisenberg's ", name.Username)
		}
	case "tt":
		getHistory(stringArg)

	}

	cmd := exec.Command(arg[0], arg[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func getHistory(commands string) {
	h := history.Buffer(commands)
	history.Showbuffer(h)

	fmt.Println("the buffer is: ", h)
}
