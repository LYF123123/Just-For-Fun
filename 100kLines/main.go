package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Create("./main.c")
	if err != nil {
		panic(err)
	}

	f.WriteString(fmt.Sprintln(`#include <stdio.h>`))
	f.WriteString(fmt.Sprintln())
	f.WriteString(fmt.Sprintln("int main()"))
	f.WriteString(fmt.Sprintln("{"))
	f.WriteString(fmt.Sprintln("\tint i = 0;"))
	f.WriteString(fmt.Sprintln("\tscanf(\"%d\", &i);"))
	f.WriteString(fmt.Sprintln("\tswitch (i)"))
	f.WriteString(fmt.Sprintln("\t{"))
	for i := 0; i < 100000; i++ {
		num:=strconv.Itoa(i)
		f.WriteString(fmt.Sprintf("\tcase %d:\n", i))
		f.WriteString(fmt.Sprintf("\t\tprintf(\""))
		if i > 9999 {
			f.WriteString(fmt.Sprintf("是5位数"))
		} else if i > 999 {
			f.WriteString(fmt.Sprintf("是4位数"))
		} else if i > 99 {
			f.WriteString(fmt.Sprintf("是3位数"))
		} else if i > 9 {
			f.WriteString(fmt.Sprintf("是2位数"))
		} else {
			f.WriteString(fmt.Sprintf("是1位数"))
		}
		f.WriteString(fmt.Sprintf("\\n\");\n"))
		f.WriteString(fmt.Sprintf("\t\tprintf(\""))
		f.WriteString(fmt.Sprintf("倒过来是"+reserveNumber(num)))
		f.WriteString(fmt.Sprintf("\\n\");\n"))
		f.WriteString(fmt.Sprintf("\t\tbreak;\n"))
	}
	f.WriteString(fmt.Sprintln("\t}"))
	f.WriteString(fmt.Sprintln("\treturn 0;"))
	f.WriteString(fmt.Sprintln(`}`))
}

func reserveNumber(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input { 
			rune[n] = r
			n++
	} 
	rune = rune[0:n]
	// Reverse 
	for i := 0; i < n/2; i++ { 
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
	} 
	// Convert back to UTF-8. 
	output := string(rune)
	return output
}
