package no151

func reverseWords(s string) string {

	result := ""

	curWord := ""
	for i := len(s)-1; i >= 0; i-- {
		c := s[i]
		if c == ' ' {
			if curWord != "" {
				if result != "" {
					result += " "
				}
				result += curWord
				curWord = ""
			}
		} else {
			curWord = string(c) + curWord
		}
	}

	if curWord != "" {
		if result != "" {
			result += " "
		}
		result += curWord
	}
	
	return result
}