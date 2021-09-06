package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Core packages!")
	// strings
	// true/false: contains as substring or not.
	text := "sakib"
	pattern := "ki"
	fmt.Println(strings.Contains(text, pattern))

	// int: counts the number of times st occurs as substring in name
	name := "aabbaaccdd"
	st := "aa"
	fmt.Println(strings.Count(name, st))

	// true/false: checks if a string has occurred as a prefix of another string
	prefixTest := "appscode"
	prefix := "apps"
	fmt.Println(strings.HasPrefix(prefixTest, prefix))

	// true/false: checks if a string has occurred as a prefix of another string
	suffixTest := "appscode"
	suffix := "apps"
	fmt.Println(strings.HasPrefix(suffixTest, suffix))

	// first occurrenece index if present, -1 otherwise
	indexTest := "sakibalamin"
	fmt.Println(strings.Index(indexTest, "x"))

	fmt.Println(strings.Join([]string{"abc", "def"}, "xyz")) // output: abcxyzdef
	fmt.Println(strings.Repeat("sakib", 3))                  // repeat this string 3 times sakibsakibsakib

	// output: bbbkkbba, replaces first 5 occurrences of a by b
	// replaces the substring with the value given, need not to be of same length
	fmt.Println(strings.Replace("aaakkaaa", "a", "b", 5))
	fmt.Println(strings.Replace("aabaab", "aab", "x", 2))
	// if n < 0, there is no limit on the number of replacements.
	info := "sakib.alamin.appscode.dhaka-1230,bangladesh"
	fmt.Println(strings.Replace(name, ".", "-", -1))

	// split by the second param string: output: xyz, returns a list of string
	charList := strings.Split("xabyabz", "ab")
	fmt.Println(charList)
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// to change case of string
	fmt.Println(strings.ToLower("SAKIB"), strings.ToUpper("alamin"))

	// converting a string to a slice of binary data, [116 101 115 116] & "test"
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
	// see the fileIO part for file handling code & frequently used part for sorting array or slices

}
