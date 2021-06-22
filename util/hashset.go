package util

type HashSet []string

func (h *HashSet) Add(s string) {
	for _, it := range *h {
		if it == s {
			return
		}
	}
	*h = append(*h, s)
}
