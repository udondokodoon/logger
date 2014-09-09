package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/mitchellh/colorstring"
)

func main() {
	pattern, color := parseArgs(os.Args)

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("[error]: 正規表現が間違っています\n%v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Print(re.ReplaceAllStringFunc(s, func(matched string) string {
			colored := colorstring.Color(fmt.Sprintf("[%s]%s", color, matched))
			return colored
		}))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parseArgs(osArgs []string) (string, string) {
	args := make([]string, 3)
	copy(args, os.Args)

	pattern := args[1]
	if pattern == "" {
		pattern = ""
	}
	color := args[2]
	if color == "" {
		color = "red"
	}
	return pattern, color
}