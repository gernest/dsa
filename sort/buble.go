package sort

func Buble(src []int) {
	n := len(src) - 2
	t := true
	for n != 0 {
		if !t {
			break
		}
		for j := range make([]struct{}, n+1) {
			if src[j] <= src[j+1] {
				continue
			}
			src[j], src[j+1] = src[j+1], src[j]
			t = true
		}
		n--
	}
}
