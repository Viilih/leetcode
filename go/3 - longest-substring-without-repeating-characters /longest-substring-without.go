package main

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	longestSubstring := ""
	for start < len(s) {
		substring := ""
		substring += string(s[start])
		end = start + 1
		for end < len(s) && !stringContainsChar(substring, s[end]) {
			substring += string(s[end])
			end++
		}
		if len(substring) > len(longestSubstring) {
			longestSubstring = substring
		}
		start++
	}
	return len(longestSubstring)
}

func stringContainsChar(s string, character byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == character {
			return true
		}
	}
	return false
}

//0 <= s.length <= 5 * 104
//s consists of English letters, digits, symbols and spaces.
//CORE IDEA OF THE PROBLEM:a
//Two pointers
//start and end
//longestSubstring = ""
//for loop (start = 0) -> first letter
//substring += s[start]
//for loop ( end = 0) -> first letter
//the second loop iterates and check if substring does not
//contain s[end] -> add the character
//end for loop (end)
//if (substring.length > longestSubstring.length ) {
//longestsubstring = substring
