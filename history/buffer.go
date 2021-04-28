// logic:
// create an array or slice of 100 characters
// this stores the history of commands
// after every loop add the command to array

// funcitonaliy:
// up , down arrow should give the aray values
// array is a stack DS probably

package history

import (
	"fmt"
)

var History []string = make([]string, 10)
var index int = 0

func Buffer(command string) []string {
	History[index] = command
	index++
	return History
}

func Showbuffer(buf []string) {
	fmt.Println("Last command: ")
	fmt.Println(buf[index])
	index++
}
