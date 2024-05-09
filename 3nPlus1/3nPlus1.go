package main

/* program inspiration:
 * https://youtu.be/094y1Z2wpJg?si=J5aw_yDn-MmA9nJY
 * https://en.wikipedia.org/wiki/Collatz_conjecture
 */
import (
	"fmt"
	"log"
	"strconv"
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

	for num != 1 {

		fmt.Println(num)

		even := num%2 == 0
		if even {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
	}
	fmt.Println(1)

}
