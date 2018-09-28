package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		if(input.Text()==""){ break; }

		counts[input.Text()]++
	}

	for line,n := range counts {
		fmt.Println( line, n)
	}
}