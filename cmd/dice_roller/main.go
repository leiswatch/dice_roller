package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	Quits    = []string{"q", "quit"}
	Regex, _ = regexp.Compile("[0-9]+[d][0-9]+")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter dice roll: ")
		scanner.Scan()
		text := scanner.Text()

		if slices.Contains(Quits, text) {
			os.Exit(0)
		}

		if matched := Regex.MatchString(text); !matched {
			log.Fatal("Wrong value!")
		}

		texts := strings.Split(text, "d")

		d, err := strconv.Atoi(texts[0])
		if err != nil {
			log.Fatalf("Wrong value: %s", texts[0])
		}

		v, err := strconv.Atoi(texts[1])
		if err != nil {
			log.Fatalf("Wrong value: %s", texts[1])
		}

		sum := 0

		for i := range d {
			roll := rand.N(v) + 1

			if d > 1 {
				sum += roll
			}

			fmt.Printf("Dice #%d roll: %d\n", i+1, roll)
		}

		if sum > 1 {
			fmt.Printf("\nSum of rolls: %d\n", sum)
		}
	}
}
