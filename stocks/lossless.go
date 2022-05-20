package stocks

import (
	"strconv"
)

func ExecuteCompressionLossLess(input string) string {
	var result string = ""
	var count int = 1
	for i := 0; i < len(input); i++ {
		char := string(input[i])

		if i == (len(input) - 1) {
			result = result + (char + strconv.Itoa(count))
			break
		}

		nextChar := string(input[i+1])

		if char == nextChar {
			count = count + 1
		} else {
			result = result + (char + strconv.Itoa(count))
			count = 1
		}
	}

	return result
}
