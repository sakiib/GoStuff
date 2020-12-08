package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	cnt := make(map[string] int)
	for i := 0; i < len(str); i++ {
		cur := str[i : i + 1]
		cnt[cur] ++
	}

	if len(cnt) % 2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}