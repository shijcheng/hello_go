// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		counts[input.Text()]++

		if input.Text() == "end" {
			break
		}
	}

	fmt.Println(counts)

	// 注意：忽略input.Err（）的潜在错误
	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n,line)
		}
	}
}