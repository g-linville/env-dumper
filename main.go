package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/grant/env.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = file.Close()
	}()

	for _, e := range os.Environ() {
		_, _ = file.WriteString(e + "\n")
	}
}
