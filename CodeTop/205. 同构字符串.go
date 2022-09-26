package CodeTop

// 双哈希表
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t, t2s := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		bs, bt := s[i], t[i]
		if s2t[bs] == 0 && t2s[bt] == 0 {
			s2t[bs], t2s[bt] = bt, bs
			continue
		}
		if s2t[bs] != bt || t2s[bt] != bs {
			return false
		}
	}

	return true
}
