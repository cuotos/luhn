package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cuotos/luhn/luhn"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)

	inputInt, err := strconv.Atoi(strings.TrimSpace(string(input)))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	isValid, err := luhn.IsValid(inputInt)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	if isValid {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
