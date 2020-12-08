package main

func longestPalindrome(s string) string {
	length := len(s)
	mxLen := 0
	substring := ""
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			dpp[i][j] = -1
		}
	}
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if palindrome(i, j, s) == 1 {
				if j - i + 1 > mxLen {
					mxLen = j - i + 1
					substring = s[i:j + 1]
				}
			}
		}
	}
	return substring
}

var dpp[1005][1005] int

func palindrome(i int, j int, s string) int {
	if i >= j {
		return 1
	}
	ret := 0
	if dpp[i][j] != -1 {
		return dpp[i][j]
	}
	if s[i] == s[j] {
		ret = ret | palindrome(i + 1, j - 1, s)
	}
	dpp[i][j] = ret
	return dpp[i][j]
}

// Given a string s, return the longest palindromic substring in s.