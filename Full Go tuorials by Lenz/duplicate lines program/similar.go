package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
lets_count:
	make(map[string]int)
input:
	bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lets_count[input.Text()]++

		for line, n := range lets_count {
			if n > 1 {
				fmt.Println("%d\t%s\n", n, line)
			}

		}
	}
}
