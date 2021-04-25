package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("☭ ")
		inp, err := reader.ReadString('\n') // The ReadString function keeps reading input where \n is the delimiter, after
		//which it stops reading
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		inp = strings.TrimSuffix(inp, "\n")

		cmd := exec.Command(inp)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		errr := cmd.Run()
		if errr != nil {
			fmt.Fprintln(os.Stderr, errr)
		}

	}

}
