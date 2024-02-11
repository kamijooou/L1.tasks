package main

import (
	"fmt"
	"strings"
)

// O(n) spd, O(n memory)
func isUnique(str string) bool {
	if len(str) == 0 {
		return false
	}

	str = strings.ToLower(str)
	fmt.Println(str)
	runes := map[rune]struct{}{}    // O(1) spd, O(n) memory
	for _, v := range []rune(str) { //O(n) spd, O(1) memory
		if _, ok := runes[v]; ok {
			return false
		}
		runes[v] = struct{}{}
	}
	return true
}

func main() {
	// здесь русская "о" и английская, будет true
	str := "😱gAbeloо, хернЯ5"
	fmt.Println(isUnique(str))
}
