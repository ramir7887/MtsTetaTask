package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s:= CountLetters("sSaAUuuuaaaAAAA")
	fmt.Println(s)
}

func CountLetters(s string) string {
	var (
		out string
		currSimb rune
		count int
	)
	strs := strings.Split(s, "")
	strs = strs[1:len(strs)-1]
	sort.Strings(strs)
	s = strings.Join(strs, "")
	for idx, v := range s {
		fmt.Println(string(v))
		if idx == 0 {
			currSimb = v
			count = 1
			continue
		}
		if v == currSimb {
			count++
		} else {
			out += fmt.Sprintf("%c%d", currSimb, count)
			count = 1
			currSimb = v
		}
	}
	out += fmt.Sprintf("%c%d", currSimb, count)
	return out
}
