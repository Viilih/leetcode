package main

func main() {

}

func reverseWords(s string) string {
	lP, rP := 0, 0
	stringLength := len(s)
	bytes := []byte(s)
	for i := 0; i <= stringLength; i++ {
		if i == stringLength || bytes[i] == ' ' {
			rP = i - 1

			for lP < rP {
				bytes[lP], bytes[rP] = bytes[rP], bytes[lP]
				lP++
				rP--
			}
			lP = i + 1
		}
	}
	return string(bytes)
}

/*
Sliding Window Logic Explanation:
- The algorithm uses a sliding window approach to identify and reverse individual words
- lP (left pointer) marks the start of the current word window
- The loop variable i acts as a scanner that moves through the string
- When i finds a word boundary (space or end of string):
  - rP (right pointer) is set to i-1, marking the end of the current word
  - A two-pointer technique reverses characters within the window [lP, rP]
  - The window slides forward: lP moves to i+1 (start of next word)
- This creates a "sliding window" that captures each word and reverses it in-place
- Example: "hello world" → window1: "hello" → window2: "world" → result: "olleh dlrow"
*/
