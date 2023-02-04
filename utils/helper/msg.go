package helper

import "fmt"

func LogDebug(msg ...interface{}) {
	if true {
		fmt.Print("\n\n", msg, "\n\n")
	}
}
