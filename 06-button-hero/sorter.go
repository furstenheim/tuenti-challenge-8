package main

type NotesSorter []Note

func (s NotesSorter) Less (i, j int) bool {
	n1, n2 := s[i], s[j]
	if n1.starT == n2.starT {
		return n1.endT < n2.endT
	}
	return n1.starT < n2.starT
}

func (s NotesSorter) Len () int {
	return len(s)
}

func (s NotesSorter) Swap (i, j int) {
	s[i], s[j] = s[j], s[i]
}