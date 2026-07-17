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

// ======================================
// FORMAT RUPIAH
// ======================================

func FormatRupiah(amount float64) string {

	number := strconv.FormatFloat(amount, 'f', 0, 64)

	n := len(number)

	if n <= 3 {
		return "Rp " + number
	}

	result := ""

	count := 0

	for i := n - 1; i >= 0; i-- {

		result = string(number[i]) + result

		count++

		if count%3 == 0 && i != 0 {

			result = "." + result

		}

	}

	return "Rp " + result

}
