package matching

func matchParens(src string, matchs ...rune) bool {
	if len(matchs) > 0 {
		if len(matchs)%2 != 0 {
			return false
		}
	}
	open := make(map[rune]rune)
	closing := make(map[rune]rune)
	pos := 0
	var stack []rune
	for pos < len(matchs) {
		l, r := matchs[pos], matchs[pos+1]
		open[l] = r
		closing[r] = l
		pos += 2
	}
	for _, s := range src {
		if _, ok := open[s]; ok {
			stack = append(stack, s)
			continue
		}
		if r, ok := closing[s]; ok {
			if len(stack) > 0 {
				pop := stack[len(stack)-1]
				if r == pop {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}
	return len(stack) == 0
}
