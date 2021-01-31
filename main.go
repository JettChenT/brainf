package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const MAXN = 100000

var mem [MAXN]int
var jmp [MAXN]int
var stack [MAXN]int
var cmd [MAXN]string
var pointer int = 0

func run(data *bufio.Scanner, printstr bool) {
	ind := 0
	stackp := 0
	for data.Scan() {
		c := data.Text()
		cmd[ind] = c
		if c == "[" {
			stack[stackp] = ind
			stackp++
		}
		if c == "]" {
			stackp--
			jmp[stack[stackp]] = ind
			jmp[ind] = stack[stackp]
		}
		ind++
	}
	for i := 0; i < ind; i++ {
		cur := cmd[i]
		if cur == ">" {
			pointer++
		}
		if cur == "<" {
			pointer--
		}
		if cur == "+" {
			mem[pointer]++
		}
		if cur == "-" {
			mem[pointer]--
		}
		if cur == "." {
			if printstr {
				fmt.Print(string(mem[pointer]))
			} else {
				fmt.Print(mem[pointer])
			}
		}
		if cur == "," {
			var n int
			fmt.Print("\ninput:")
			_, err := fmt.Scanf("%d", &n)
			if err != nil {
				fmt.Println(err)
			}
			mem[pointer] = n
		}
		if cur == "[" && mem[pointer] == 0 {
			i = jmp[i]
		}
		if cur == "]" && mem[pointer] != 0 {
			i = jmp[i]
		}
	}
}

func main() {
	filename := flag.String("fn", "", "filename to parse(required)")
	pntstr := flag.Bool("ps", false, "whether or not the program outputs the string")
	flag.Parse()
	filebuffer, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	run(data, *pntstr)
}
