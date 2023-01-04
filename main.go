package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s string

	fmt.Print("\nType q to quit \n\n")

	fmt.Print("Roll: ")
	fmt.Scan(&s)

	s = strings.TrimSpace(s)

	if s == "q" {
		os.Exit(0)
	}

	match, _ := regexp.MatchString("^([1-9]*\\d*k[1-9]+\\d*)$", s)

	if match {
		numbers := strings.Split(s, "k")

		dices_num, _ := strconv.Atoi(numbers[0])
		dice_sides, _ := strconv.Atoi(numbers[1])

		if dices_num == 0 {
			dices_num = 1
		}

		sum := 0

		for i := 0; i < dices_num; i++ {
			rand.Seed(time.Now().UnixNano())

			num := rand.Intn(dice_sides) + 1
			sum += num

			fmt.Printf("Dice(%d): %d\n", i+1, num)
		}

		fmt.Print("\nSum: ", sum, "\n")
	}
}
