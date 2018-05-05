package main


type rowsSorter []row

func (s rowsSorter) Less (i, j int) bool {
	n1, n2 := s[i], s[j]
	return n1.minX < n2.minX
}

func (s rowsSorter) Len () int {
	return len(s)
}

func (s rowsSorter) Swap (i, j int) {
	s[i], s[j] = s[j], s[i]
}