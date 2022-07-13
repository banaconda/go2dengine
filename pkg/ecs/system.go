package ecs

type System interface {
	Priority() int
	Add(Identifier)
	Update()
	Remove()
}

type systems []System

// Len implements sort.Interface
func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	return s[i].Priority() > s[j].Priority()
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
