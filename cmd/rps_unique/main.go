package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexmerren/rps/src"
)

func main() {
	uniqueMap := src.NewUniqueMap()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if uniqueMap.Add(line) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
