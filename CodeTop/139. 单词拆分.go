package CodeTop

func wordBreak(s string, wordDict []string) bool {
	cache := make(map[string]bool)

	for i := 0; i < len(wordDict); i++ {
		cache[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && cache[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
