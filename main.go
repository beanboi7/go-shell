package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	for {
		print("@ ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = strings.TrimSuffix(text, "\r\n")

		stdout, err := exec.Command(text).Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(stdout))
	}

}
