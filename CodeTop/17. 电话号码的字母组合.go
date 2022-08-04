package CodeTop

var phoneMap map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var combinations []string
	var backtrack func(digits string, index int, combination []byte)
	backtrack = func(digits string, index int, combination []byte) {
		if index == len(digits) {
			combinations = append(combinations, string(combination))
		} else {
			letters := phoneMap[digits[index]]
			lettersCount := len(letters)
			for i := 0; i < lettersCount; i++ {
				backtrack(digits, index+1, append(combination, letters[i]))
			}
		}
	}
	backtrack(digits, 0, []byte{})
	return combinations
}
