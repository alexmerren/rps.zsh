package src

type UniqueMap struct {
	entries map[string]struct{}
}

func NewUniqueMap() *UniqueMap {
	return &UniqueMap{
		entries: map[string]struct{}{},
	}
}

func (um *UniqueMap) Add(input string) bool {
	if _, present := um.entries[input]; !present {
		um.entries[input] = struct{}{}
		return true
	}

	return false
}
