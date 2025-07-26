package _8___find_index_first_occurrence_in_string

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

/*

Por que <= len(haystack)-len(needle)?

Porque o slice usa limite exclusivo. Se haystack tem tamanho 8 e needle tem tamanho 3,
o último começo válido é i = 5, já que haystack[5:8] tem tamanho 3. Em i = 6, 6+3 = 9 → estoura.


*/
