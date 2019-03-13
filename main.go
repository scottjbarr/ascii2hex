package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := ""

	if len(os.Args) > 1 {
		// get the supplied args
		for _, s := range os.Args[1:] {
			data += s + " "
		}
	}

	if len(data) == 0 {
		// no data so trying stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			data = s.Text()
		}
	}

	b := []byte(strings.Trim(data, " "))

	fmt.Printf("%x\n", b)
}
