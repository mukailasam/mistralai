package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mukailasam/mistralai"
)

func main() {
	mistralai.LoadConfig("../.env")

	for {
		fmt.Println("Prompt:")
		readInput := bufio.NewReader(os.Stdin)
		out, err := readInput.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		switch runtime.GOOS {
		case "windows":
			out = strings.TrimSuffix(out, "\r\n")
		case "linux":
			out = strings.TrimSuffix(out, "\n")
		}

		mistralai.Request(out)
		fmt.Println()
	}

}
