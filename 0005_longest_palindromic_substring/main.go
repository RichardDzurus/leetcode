package main

func isPalindrome(s string) bool {
	strLen := len(s)
	for i := 0; i < strLen/2; i += 1 {
		if s[i] != s[strLen-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen <= 1 {
		return s
	}

	longestPal := ""

	for i := 0; i < strLen-1; i += 1 {
		isPal2 := isPalindrome(s[i : i+2])
		isPal3 := false
		if (i + 2) < strLen {
			isPal3 = isPalindrome(s[i : i+3])
		}
		if isPal2 == false && isPal3 == false {
			continue
		}
		startOffset := 1
		endOffset := 1
		if isPal3 {
			startOffset = 2
			endOffset = 2
		}
		if isPal2 == true && isPal3 == true {
			startOffset = 1
			endOffset = 2
		}
		for offset := startOffset; offset <= endOffset; offset += 1 {
			j := i
			k := i + offset
			for true {
				if s[j] != s[k] {
					break
				}
				if len(s[j:k+1]) > len(longestPal) {
					longestPal = s[j : k+1]
				}
				j -= 1
				k += 1
				if j < 0 || k >= strLen {
					break
				}
			}
		}
	}

	if len(longestPal) == 0 && strLen > 0 {
		return string(s[0])
	}

	return longestPal
}

func main() {
	println(longestPalindrome("bac"))
	println(longestPalindrome("a"))
	println(longestPalindrome("aa"))
	println(longestPalindrome("cc"))
}
