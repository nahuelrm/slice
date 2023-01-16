package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	lines    []string
	red      string = "\033[31m"
	endColor string = "\033[0m\n"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	if len(os.Args) == 1 || len(os.Args) > 2 {
		helpPanel()
		os.Exit(0)
	}

	checkOptions(os.Args[1])

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	var options [][]string

	for _, option := range strings.Split(os.Args[1], ",") {
		options = append(options, strings.Split(option, ":"))
	}

	for _, option := range options {
		handleOption(option)
	}

}

func handleOption(option []string) {
	if len(option) == 1 {
		n, _ := strconv.Atoi(option[0])

		if n < 0 {
			n = len(lines) + n + 1
		}

		if n-1 >= 0 && n-1 < len(lines) {
			fmt.Println(lines[n-1])
		}

		return
	}

	if len(option[0]) == 0 && len(option[len(option)-1]) == 0 {
		return
	}

	start, _ := strconv.Atoi(option[0])
	end, _ := strconv.Atoi(option[len(option)-1])

	if len(option[0]) == 0 {
		start = 0
	} else if start < 0 {
		start = len(lines) + start
	} else {
		start--
	}

	if end < 0 {
		end = len(lines) + end
	}

	if start > end && len(option[len(option)-1]) != 0 {
		return
	}

	if start >= len(lines) || end > len(lines) || start < 0 || end < 0 {
		return
	}

	if len(option[len(option)-1]) == 0 {
		end = len(lines)
	}

	for _, element := range lines[start:end] {
		fmt.Println(element)
	}
	return
}

func checkOptions(options string) {
	if options == "--help" || options == "-h" {
		helpPanel()
		os.Exit(0)
	}

	validChars := regexp.MustCompile(`^[0-9,:-]*$`).MatchString
	if !validChars(options) {
		fmt.Fprintf(os.Stderr, red+"\n[!] Unsupported characters."+endColor)
		helpPanel()
		os.Exit(0)
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	if !hasNumber(options) {
		fmt.Fprintf(os.Stderr, red+"\n[!] You must enter at least one line number."+endColor)
		helpPanel()
		os.Exit(0)
	}
}

func helpPanel() {
	fmt.Println("\nUsage: <command> | slice '<options>'")
	fmt.Println("This tool works providing it STDIN, you can do it using pipes in bash or in any other way.")
	fmt.Println("\nNote: the ':' operator can be used to specify a range of lines to output.")
	fmt.Println("You can use multiple filter options separated by a comma, for example: '1,3:-7,-2'.")
	fmt.Println("\nExamples of possible options:")
	fmt.Println("  '5'      This option will print the 5th line.")
	fmt.Println("  '7:'     This option will print all the content from the 7th line to the end.")
	fmt.Println("  ':7'     This option will print all the content from initial line to the 7th line.")
	fmt.Println("  '3:7'    This option will print all the content from the 3rd line to the 7th line.")
	fmt.Println("  '4:-1'   This option will print all the content from the 4th line to the penultimate line.")
	fmt.Println("  '4,8,-1' This option combines 3 different options and will print the 4th, 8th and the penultimate line.")
}
