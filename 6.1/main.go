package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.TrimSpace(s)
	r := regexp.MustCompile("\\s+")
	s = r.ReplaceAllString(s, " ")
	fmt.Println(s)
}
