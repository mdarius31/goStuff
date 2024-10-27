package main

import (
	"fmt"
	"log"
	"strconv"
	t "threeNPlusOne"
)

func main() {
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalln(err)
	}

	pInput, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		log.Fatalln(err)
	}

	num := pInput

	threeNPlusOne := t.GenThreeNPlusOne(num)

	for _, n := range threeNPlusOne {

		fmt.Println(n)

	}

}
