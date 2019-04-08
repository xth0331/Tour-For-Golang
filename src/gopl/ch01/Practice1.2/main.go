// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%v %v\n", arg, index+1)
	}
}
