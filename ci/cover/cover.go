package lol

import (
	"fmt"
	"os"
	"strconv"
)

func lol() {
	args := os.Args
	if len(args) <= 0 {
		fmt.Println("Not enough arguments")
		return
	}

	percentage, _ := strconv.Atoi(args[1])
	fmt.Println(percentage)

	if percentage >= 40 {
		fmt.Println("above threshold")
		os.Exit(0)
	} else {
		fmt.Println("below threshold")
		os.Exit(1)
	}

}
