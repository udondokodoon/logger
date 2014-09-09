package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/colorstring"
)

var (
	pattern = ""
	color   = "red"
)

func main() {
	for i, v := range os.Args {
		switch i {
		case 1:
			pattern = v
		case 2:
			color = v
		}
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("[error]: 正規表現が間違っています\n%v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		var response, oldString, newString string
		if re.MatchString(s) {
			oldString = re.FindString(s)
			newString = colorstring.Color(fmt.Sprintf("[%s]%s", color, oldString))
			response = strings.Replace(s, oldString, newString, -1)
		}
		fmt.Print(response)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}