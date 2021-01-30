package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var mem [100000]int
var pointer int = 0

func run(data *bufio.Scanner) {
	for data.Scan() {
		c := data.Text()
		if c == ">" {
			pointer++
		} else if c == "<" {
			pointer--
		} else if c == "+" {
			mem[pointer]++
		} else if c == "-" {
			mem[pointer]--
		} else if c == "." {
			fmt.Println(mem[pointer])
		}
	}
}

func main() {
	filename := flag.String("fn", "", "filename to parse(required)")
	flag.Parse()
	filebuffer, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	run(data)
}
