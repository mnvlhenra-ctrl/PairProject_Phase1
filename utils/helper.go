package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ======================================
// UI
// ======================================

func PrintTitle(title string) {

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println(title)
	fmt.Println("======================================================")

}

func PrintLine() {

	fmt.Println("------------------------------------------------------")

}

func Pause() {

	fmt.Println()
	fmt.Print("Press ENTER to continue...")

	reader.ReadString('\n')

}

// ======================================
// INPUT
// ======================================

func InputString(prompt string) string {

	for {

		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input != "" {
			return input
		}

		fmt.Println("Input cannot be empty.")

	}

}

func InputInt(prompt string) int {

	for {

		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")

	}

}

func InputFloat(prompt string) float64 {

	for {

		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")

	}

}
