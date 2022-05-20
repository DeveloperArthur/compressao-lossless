package main

import (
	"fmt"
	"os"
	c "tdd/stocks"
)

func main() {
	input := os.Args[len(os.Args)-1]
	fmt.Println(c.ExecuteCompressionLossLess(input))
}
