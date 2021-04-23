package intSet

type IntSet struct {
	value []int
}

func Of(slice []int) *IntSet {
	mapSlice := make(map[int]bool)
	var uniqueSlice []int
	for _, element := range slice {
		if !mapSlice[element] {
			mapSlice[element] = true
			uniqueSlice = append(uniqueSlice, element)
		}
	}
	return &IntSet{value: uniqueSlice}
}

func (t IntSet) ToSlice() []int {
	return t.value
}
