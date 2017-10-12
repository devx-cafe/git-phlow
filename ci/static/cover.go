package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("coverfiles/percentage")
	if err != nil {
		log.Panicln(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	percentage, _ := strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
	fmt.Printf("%g", percentage)

	if percentage >= 25.0 {
		fmt.Printf("%g above threshold \n", percentage)
		os.Exit(0)
	} else {
		fmt.Printf("%g below threshold \n", percentage)
		os.Exit(1)
	}

}
